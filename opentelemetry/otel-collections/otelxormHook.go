package main

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"xorm.io/xorm/contexts"
)

type otelXormHook struct {
	tracer trace.Tracer
}

func NewOtelXormHook(tracer trace.Tracer) *otelXormHook {
	return &otelXormHook{tracer: tracer}
}

// BeforeProcess: 在执行 SQL 前开启一个 span
func (h *otelXormHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	// 获取操作类型（从 SQL 推断）
	op := "xorm-query"
	if len(c.SQL) >= 6 {
		prefix := c.SQL[:6]
		switch {
		case prefix == "INSERT":
			op = "xorm-insert"
		case prefix == "SELECT":
			op = "xorm-select"
		case prefix == "UPDATE":
			op = "xorm-update"
		case prefix == "DELETE":
			op = "xorm-delete"
		}
	}

	ctx, span := h.tracer.Start(c.Ctx, op)
	span.SetAttributes(
		attribute.String("db.system", "mysql"),
		attribute.String("db.statement", c.SQL),
	)
	c.Ctx = context.WithValue(ctx, "otelXormSpan", span)
	return c.Ctx, nil
}

// AfterProcess: 执行 SQL 后结束 span
func (h *otelXormHook) AfterProcess(c *contexts.ContextHook) error {
	if val := c.Ctx.Value("otelXormSpan"); val != nil {
		if span, ok := val.(trace.Span); ok {
			span.End()
		}
	}
	return nil
}
