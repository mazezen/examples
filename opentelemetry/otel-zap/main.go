// go get github.com/uptrace/opentelemetry-go-extra/otelzap

package main

import (
	"context"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

func main() {

	logger := otelzap.New(
		zap.NewExample(),
		otelzap.WithMinLevel(zap.InfoLevel),
		otelzap.WithStackTrace(true))
	defer logger.Sync()

	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	otelzap.L().Info("test zap's global loggers")
	otelzap.L().Info("测试日志", zap.String("abc", "123"))
	otelzap.Ctx(context.TODO()).Info("... and with context")
}
