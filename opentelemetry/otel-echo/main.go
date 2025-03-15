// https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho
// go get go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho
// go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp
// go get go.opentelemetry.io/otel/sdk/trace

package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"
)

const (
	serviceName    = "Echo-Jaeger-Demo2"
	jaegerEndPoint = "127.0.0.1:4318"
)

var tracer = otel.Tracer("echo-server2")

func main() {
	ctx := context.Background()

	tp, err := initTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()

	e.Use(otelecho.Middleware(serviceName))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := trace.SpanFromContext(c.Request().Context()).SpanContext().TraceID().String()
			fmt.Println(s)
			c.Request().Header.Set("Trace-Id", s)
			fmt.Println(c.Request().Header.Get("Trace-Id"))
			return next(c)
		}
	})

	e.GET("/echo/:id", func(c echo.Context) error {
		id := c.Param("id")
		name := getUser(c, id)

		var m = make(map[string]string)
		m["name"] = name
		return c.JSON(http.StatusOK, name)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func newJaegerTraceProvider(ctx context.Context) (*traceSDK.TracerProvider, error) {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndPoint),
		otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceName(serviceName)))
	if err != nil {
		return nil, err
	}

	traceProvider := traceSDK.NewTracerProvider(
		traceSDK.WithResource(res),
		traceSDK.WithSampler(traceSDK.AlwaysSample()),
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

func initTracer(ctx context.Context) (*traceSDK.TracerProvider, error) {
	tp, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)

	return tp, nil
}

func getUser(c echo.Context, id string) string {
	_, span := tracer.Start(
		c.Request().Context(), "getUser", trace.WithAttributes(attribute.String("id", id)),
	)
	defer span.End()

	if id == "7" {
		return "qimi"
	}
	return "unknown"
}
