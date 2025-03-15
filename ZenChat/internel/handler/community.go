package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/em"
	"github.com/mazezen/zenchat/internel/handler/in"
	"github.com/mazezen/zenchat/internel/models"
	"github.com/mazezen/zenchat/internel/service"
	"net/http"
)

type CommunityHandler struct{}

func NewCommunityHandler() *CommunityHandler {
	return &CommunityHandler{}
}

// CreateCommunity 创建群聊
func (community *CommunityHandler) CreateCommunity(c echo.Context) error {
	var payload = new(in.CreateCommunityPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(payload); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}
	if err := service.NewCommunityService().CreateCommunity(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "create community success"))
}

// UpdateCommunity 更新群聊信息
func (community *CommunityHandler) UpdateCommunity(c echo.Context) error {
	var payload = new(in.UpdateCommunityPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(payload); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}
	if err := service.NewCommunityService().UpdateCommunity(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "update community success"))
}

// JoinCommunity 加入群聊
func (community *CommunityHandler) JoinCommunity(c echo.Context) error {
	var payload = new(in.JoinCommunityPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(payload); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}

	if err := service.NewCommunityService().JoinCommunity(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "join community success"))
}

// ListByOwner 获取群列表
func (community *CommunityHandler) ListByOwner(c echo.Context) error {
	var payload = new(in.ListByOwnerPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(payload); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}
	lgnUser := c.Get(em.LgnPwdUser).(*models.CUser)
	if lgnUser.ID != payload.OwnerId {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, "用户ID不正确"))
	}

	count, communities, err := service.NewCommunityService().ListCommunity(payload)
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success", itools.ResponsePage.Pagination(int(count), communities)))
}
