package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/internel/service"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return new(UserHandler)
}

func (u *UserHandler) Register(c echo.Context) error {
	type Payload struct {
		Username       string `json:"username" validate:"required"`
		Password       string `json:"password" validate:"required"`
		RepeatPassword string `json:"repeat_password" validate:"required"`
		Phone          string `json:"phone" validate:"required"`
		Email          string `json:"email" validate:"required,email"`
		Identity       string `json:"identity" validate:"required"`
	}
	var p Payload
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(p); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}

	if err := service.NewUserService().Register(p.Username, p.Password, p.RepeatPassword, p.Email, p.Phone, p.Identity); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}

	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "register success"))
}

func (u *UserHandler) LgnPwd(c echo.Context) error {
	type Payload struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var p Payload
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(p); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}
	token, err := service.NewUserService().LgnPwd(p.Username, p.Password)
	if err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}

	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "lgn password success", token))
}

func (u *UserHandler) Update(c echo.Context) error {
	type Payload struct {
		Username       string `json:"username" validate:"required"`
		Password       string `json:"password" validate:"required"`
		RepeatPassword string `json:"repeat_password" validate:"required"`
		Phone          string `json:"phone" validate:"required"`
		Email          string `json:"email" validate:"required,email"`
		Identity       string `json:"identity" validate:"required"`
		Avatar         string `json:"avatar" validate:"required"`
		Gender         string `json:"gender"`
	}
	var p Payload
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(p); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}
	if err := service.NewUserService().Update(
		p.Username,
		p.Password,
		p.RepeatPassword,
		p.Email,
		p.Phone,
		p.Identity,
		p.Avatar,
		p.Gender); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "update success"))
}

func (u *UserHandler) Delete(c echo.Context) error {
	var payload struct {
		Id uint64 `json:"id" validate:"required,numeric"`
	}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	if msg := itools.ZhValidateParam(payload.Id); len(msg) > 0 {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, msg))
	}
	if err := service.NewUserService().Delete(payload.Id); err != nil {
		return c.JSON(http.StatusOK, itools.Response.ResponseError("", 5000, err.Error()))
	}
	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "delete success"))
}

func (u *UserHandler) List(c echo.Context) error {
	count, list, err := service.NewUserService().List()
	if err != nil {
		itools.AppLog.Error("user list error: ", zap.Error(err))
	}

	return c.JSON(http.StatusOK, itools.Response.ResponseSuccess("", "success", itools.ResponsePage.Pagination(int(count), list)))
}
