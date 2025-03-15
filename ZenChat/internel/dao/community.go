package dao

import (
	"errors"
	"fmt"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/handler/in"
	"github.com/mazezen/zenchat/internel/models"
	"xorm.io/xorm"
)

type CommunityDao struct{}

var NewCommunityDao = new(CommunityDao)

// CreateCommunity 创建群聊
func (c *CommunityDao) CreateCommunity(community *models.CCommunity) error {
	has, err := c.existsCommunity(community.Name)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("群 [%s] 已经存在", community.Name)
	}

	err = itools.Transaction(func(s *xorm.Session) error {
		// 创建群聊
		if err = c.insertCommunityWithSession(s, community); err != nil {
			return err
		}

		user, err := NewUserDao().FindUserById(community.OwnerId)
		if err != nil {
			return err
		}

		communityRelationBean := &models.CCommunityRelation{
			OwnerID:  community.OwnerId,
			TargetID: community.ID,
			Desc:     fmt.Sprintf("[%s] 创建了群 [%s]", user.Username, community.Name),
		}

		// 将群主加入群关系
		if err = NewCommunityRelationDao.InsertCommunityRelationWithSession(s, communityRelationBean); err != nil {
			return err
		}

		return nil
	})

	return err
}

func (c *CommunityDao) UpdateCommunity(payload *in.UpdateCommunityPayload) error {
	exists, err := c.existsCommunityById(payload.CommunityId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("群聊已解散")
	}
	communityBean := &models.CCommunity{
		Name:   payload.Name,
		Type:   payload.Type,
		Avatar: payload.Avatar,
		Desc:   payload.Desc,
	}
	affected, err := itools.Db.Where("id = ?", payload.CommunityId).Update(communityBean)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("更新群聊信息失败")
	}
	return nil
}

func (c *CommunityDao) JoinCommunity(payload *in.JoinCommunityPayload, username, communityName string) error {
	err := itools.Transaction(func(s *xorm.Session) error {
		bean := &models.CCommunityRelation{
			OwnerID:  payload.OwnerId,
			TargetID: payload.CommunityId,
			Desc:     fmt.Sprintf("[%s]加入了群聊[%s]", username, communityName),
		}
		err := NewCommunityRelationDao.InsertCommunityRelationWithSession(s, bean)
		return err
	})
	return err
}

func (c *CommunityDao) ExistsCommunityById(id uint64) (*models.CCommunity, error) {
	return c.getCommunityById(id)
}

func (c *CommunityDao) ListByOwner(payload *in.ListByOwnerPayload) (int64, []*models.CCommunity, error) {
	var list []*models.CCommunity

	var orm = itools.Db.Where("1=1")
	if payload.Page != 0 && payload.PageSize != 0 {
		orm = orm.Limit(int(payload.PageSize), (int(payload.Page)-1)*int(payload.PageSize))
	}
	count, err := orm.OrderBy("id desc").FindAndCount(&list)
	if err != nil {
		return 0, nil, err
	}
	return count, list, nil
}

func (c *CommunityDao) insertCommunityWithSession(s *xorm.Session, community *models.CCommunity) error {
	affected, err := s.Insert(community)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("创建群聊失败")
	}
	return nil
}

func (c *CommunityDao) existsCommunity(name string) (bool, error) {
	return itools.Db.Where("name = ?", name).Exist(new(models.CCommunity))
}

func (c *CommunityDao) existsCommunityById(id uint64) (bool, error) {
	return itools.Db.Where("id = ?", id).Exist(new(models.CCommunity))
}

func (c *CommunityDao) getCommunityById(id uint64) (*models.CCommunity, error) {
	var bean = new(models.CCommunity)
	exist, err := itools.Db.Where("id = ?", id).Exist(bean)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("群聊已解散")
	}
	return bean, nil
}
