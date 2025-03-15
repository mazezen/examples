package middle

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/em"
	"github.com/mazezen/zenchat/common/sdk"
	"github.com/mazezen/zenchat/internel/dao"
	"net/http"
	"strings"
	"time"
)

var lgnPwdInfo struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func JwtAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get(echo.HeaderAuthorization)
			split := strings.SplitN(authorization, " ", 2)
			if len(split) != 2 || split[0] != "Bearer" {
				return c.JSON(http.StatusOK, itools.Response.ResponseError("", echo.ErrUnauthorized.Code, echo.ErrUnauthorized.Error()))
			}
			tokenString := split[1]
			ex := time.Duration(sdk.GetConf().Jwt.Expire) * time.Second
			claims, err := itools.NewJwt(ex, sdk.GetConf().Jwt.Secret).ParseToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusOK, itools.Response.ResponseError("", echo.ErrUnauthorized.Code, err.Error()))
			}
			bytes, _ := json.Marshal(claims.LoginInfo)
			var l = lgnPwdInfo
			_ = json.Unmarshal(bytes, &l)
			user, err := dao.NewUserDao().FindUserById(l.Id)
			if err != nil {
				return c.JSON(http.StatusOK, itools.Response.ResponseError("", echo.ErrUnauthorized.Code, echo.ErrUnauthorized.Error()))
			}
			if l.Password != user.Password {
				return c.JSON(http.StatusOK, itools.Response.ResponseError("", echo.ErrUnauthorized.Code, "password error, please try login"))
			}
			c.Set(em.LgnPwdUser, user)
			return next(c)
		}
	}
}
