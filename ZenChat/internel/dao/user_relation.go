package dao

import (
	"errors"
	"fmt"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/em"
	"github.com/mazezen/zenchat/internel/models"
	"xorm.io/xorm"
)

type UserUserRelationDao struct{}

func NewUserRelationDao() *UserUserRelationDao {
	return new(UserUserRelationDao)
}

// FriendList 好友列表
func (r *UserUserRelationDao) FriendList(userId uint64) (int64, []*models.CUser, error) {
	var relationsIds []uint64
	if err := itools.Db.Table(&models.CUserRelation{}).
		Where("owner_id = ?", userId).
		And("`type` = ?", em.FriendRelationShip).
		Select("owner_id").Find(&relationsIds); err != nil {
		return 0, nil, err
	}

	var users []*models.CUser
	count, err := itools.Db.In("id", relationsIds).FindAndCount(&users)
	if err != nil {
		return 0, nil, err
	}
	return count, users, nil
}

// FriendAdd 通过ID 添加好友
func (r *UserUserRelationDao) FriendAdd(userId uint64, targetUserData *models.CUser) error {
	if userId == targetUserData.ID {
		return errors.New("不能添加自己为好友")
	}
	userDao := NewUserDao()
	userData, err := userDao.FindUserById(userId)
	if err != nil {
		return err
	}

	exists, err := r.FindFriendIsExists(userId, targetUserData.ID)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("好友关系已存在")
	}

	if err = itools.Transaction(func(s *xorm.Session) error {
		userRelation := &models.CUserRelation{
			OwnerID:  userId,
			TargetID: targetUserData.ID,
			Type:     em.FriendRelationShip,
			Desc:     fmt.Sprintf("%s向%s添加好友", userData.Username, targetUserData.Username),
		}
		targetRelation := &models.CUserRelation{
			OwnerID:  targetUserData.ID,
			TargetID: userId,
			Type:     em.FriendRelationShip,
			Desc:     fmt.Sprintf("%s向%s添加好友", userData.Username, targetUserData.Username),
		}
		affected, err := s.Insert(userRelation, targetRelation)
		if err != nil {
			return err
		}
		if affected != 2 {
			return errors.New("添加好友失败")
		}

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *UserUserRelationDao) FriendAddByName(userId uint64, targetName string) error {
	userDao := NewUserDao()
	targetUserData, err := userDao.FindUserByName(targetName)
	if err != nil {
		return err
	}
	if err = r.FriendAdd(userId, targetUserData); err != nil {
		return err
	}
	return nil
}

func (r *UserUserRelationDao) FriendRemove(userId uint64, targetName string) error {
	userDao := NewUserDao()
	// 查询好友是否存在
	targetUser, err := userDao.FindUserByName(targetName)
	if err != nil {
		return err
	}
	// 查询好友关系是否存在
	exists, err := r.FindFriendIsExists(userId, targetUser.ID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("好友关系不存在")
	}
	if err = itools.Transaction(func(s *xorm.Session) error {
		err = r.FriendDelete(s, userId, targetUser.ID)
		if err != nil {
			return err
		}
		if err = r.FriendDelete(s, targetUser.ID, userId); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// FindFriendIsExists 查询好友关系是否存在
func (r *UserUserRelationDao) FindFriendIsExists(userId, targetId uint64) (bool, error) {
	var relation = new(models.CUserRelation)
	has, err := itools.Db.Where("owner_id = ? and target_id = ?", userId, targetId).Get(relation)
	if err != nil {
		return false, err
	}
	return has, nil
}

func (r *UserUserRelationDao) FriendDelete(s *xorm.Session, userId uint64, targetId uint64) error {
	affected, err := s.Where("owner_id = ? and target_id = ?", userId, targetId).Delete(new(models.CUserRelation))
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("delete relation failed")
	}
	return nil
}
