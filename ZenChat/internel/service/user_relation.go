package service

import (
	"github.com/mazezen/zenchat/internel/dao"
	"github.com/mazezen/zenchat/internel/handler/in"
	"github.com/mazezen/zenchat/internel/models"
)

type UserRelationService struct{}

func NewRelationService() *UserRelationService {
	return new(UserRelationService)
}

func (r *UserRelationService) RelationList(payload *in.RelationListPayload) (int64, []*models.CUser, error) {
	return dao.NewUserRelationDao().FriendList(payload.UserId)
}

func (r *UserRelationService) FriendAddByName(payload *in.FriendAddByNamePayload) error {
	return dao.NewUserRelationDao().FriendAddByName(payload.UserId, payload.TargetName)
}

func (r *UserRelationService) FriendRemoveByName(payload *in.RemoveRelationByName) error {
	return dao.NewUserRelationDao().FriendRemove(payload.UserId, payload.TargetName)
}
