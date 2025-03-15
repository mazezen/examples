package middle

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/em"
	"go.uber.org/zap"
	"strings"
)

var BodyDumpConfig = middleware.BodyDumpConfig{
	Skipper: bodyDumpDefaultSkipper,
	Handler: bodyDumpHandler,
}

func bodyDumpDefaultSkipper(c echo.Context) bool {
	if !strings.HasPrefix(c.Path(), "/chat/") {
		return true
	}
	return false
}

func bodyDumpHandler(context echo.Context, requestBody, responseBody []byte) {
	if !strings.HasPrefix(context.Path(), "/chat/") {
		return
	}
	uid := context.Get(em.SysTraceId).(string)
	itools.AppLog.Info("end of request", zap.String("request traceId", uid), zap.String("request ip", context.RealIP()))
}
