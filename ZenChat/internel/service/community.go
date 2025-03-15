package service

import (
	"fmt"
	"github.com/mazezen/zenchat/internel/dao"
	"github.com/mazezen/zenchat/internel/handler/in"
	"github.com/mazezen/zenchat/internel/models"
)

type CommunityService struct{}

func NewCommunityService() *CommunityService {
	return &CommunityService{}
}

func (c *CommunityService) CreateCommunity(payload *in.CreateCommunityPayload) error {
	communityDao := dao.NewCommunityDao
	communityBean := &models.CCommunity{
		Name:    payload.Name,
		OwnerId: payload.OwnerId,
		Type:    payload.Type,
		Avatar:  payload.Avatar,
		Desc:    payload.Desc,
	}

	if err := communityDao.CreateCommunity(communityBean); err != nil {
		return err
	}

	return nil
}

func (c *CommunityService) UpdateCommunity(payload *in.UpdateCommunityPayload) error {
	return dao.NewCommunityDao.UpdateCommunity(payload)
}

func (c *CommunityService) JoinCommunity(payload *in.JoinCommunityPayload) error {
	user, err := dao.NewUserDao().FindUserById(payload.OwnerId)
	if err != nil {
		return err
	}
	communityDao := dao.NewCommunityDao
	communityData, err := communityDao.ExistsCommunityById(payload.CommunityId)
	if err != nil {
		return err
	}

	// 判断是否已经在群里中
	if isInCommunity := dao.NewCommunityRelationDao.IsInCommunity(payload.OwnerId, payload.CommunityId); isInCommunity {
		return fmt.Errorf("[%s]已经在群聊中", user.Username)
	}

	return communityDao.JoinCommunity(payload, user.Username, communityData.Name)
}

func (c *CommunityService) ListCommunity(payload *in.ListByOwnerPayload) (int64, []*models.CCommunity, error) {
	_, err := dao.NewUserDao().FindUserById(payload.OwnerId)
	if err != nil {
		return 0, nil, nil
	}
	return dao.NewCommunityDao.ListByOwner(payload)
}
