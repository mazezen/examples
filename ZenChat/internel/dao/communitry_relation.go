package dao

import (
	"errors"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/models"
	"xorm.io/xorm"
)

type CCommunityRelationDao struct{}

var NewCommunityRelationDao = new(CCommunityRelationDao)

func (cr *CCommunityRelationDao) InsertCommunityRelationWithSession(s *xorm.Session, relation *models.CCommunityRelation) error {
	affected, err := s.Insert(relation)
	if err != nil {
		return err
	}
	if affected == 0 {
		return errors.New("创建群聊关系失败")
	}
	return nil
}

func (cr *CCommunityRelationDao) IsInCommunity(ownerId, targetId uint64) bool {
	isIn, err := itools.Db.Where("owner_id = ? AND target_id = ?", ownerId, targetId).Exist(new(models.CCommunityRelation))
	if err != nil {
		return false
	}
	return isIn
}

func (cr *CCommunityRelationDao) FindUsers(communityId uint64) (*[]uint64, error) {
	relations := make([]models.CCommunityRelation, 0)
	if err := itools.Db.Where("target_id = ?", communityId).Find(&relations); err != nil {
		return nil, err
	}
	userIds := make([]uint64, 0)
	for _, relation := range relations {
		userIds = append(userIds, relation.OwnerID)
	}
	return &userIds, nil
}
