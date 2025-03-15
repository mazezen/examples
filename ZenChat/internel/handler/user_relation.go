package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/handler/in"
	"github.com/mazezen/zenchat/internel/service"
	"net/http"
)

type UserRelationHandler struct{}

func NewUserRelationHandler() *UserRelationHandler {
	return new(UserRelationHandler)
}

// RelationList 好友关系列表
func (r *UserRelationHandler) RelationList(c echo.Context) error {
	payload := new(in.RelationListPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, "failed"))
	}
	count, users, err := service.NewRelationService().RelationList(payload)
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success list", itools.ResponsePage.Pagination(int(count), users)))
}

// FriendAddByName 根据用户名添加好友
func (r *UserRelationHandler) FriendAddByName(c echo.Context) error {
	payload := new(in.RemoveRelationByName)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, "friend add by name failed"))
	}
	if err := service.NewRelationService().FriendAddByName(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success add by name"))
}

// RemoveRelationByName 移除指定好友关系
func (r *UserRelationHandler) RemoveRelationByName(c echo.Context) error {
	payload := new(in.RemoveRelationByName)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if err := service.NewRelationService().FriendRemoveByName(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success remove by name"))
}
