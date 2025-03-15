package middle

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mazezen/itools"
	"github.com/mazezen/zenchat/common/em"
	"go.uber.org/zap"
	"runtime"
	"strings"
)

func RequestLog() echo.MiddlewareFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := make([]byte, 4<<10)
					length := runtime.Stack(stack, false)
					itools.AppLog.Error("program crash", zap.String("crash log", string(stack[:length])))
				}
			}()
			if !strings.HasPrefix(context.Path(), "/chat/") {
				itools.AppLog.Warn("request to start, url lack `chat` keyword", zap.Any(context.Request().RequestURI, "web request"))
				return handlerFunc(context)
			}
			uid := uuid.New().String()
			context.Set(em.SysTraceId, uid)
			itools.AppLog.Info("request to start", zap.Any(context.Request().RequestURI, uid))
			context.Set("TraceId", uid)
			err := handlerFunc(context)
			return err
		}
	}
}
