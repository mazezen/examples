package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"math/rand"
	"time"
)

const (
	serviceName    = "Go-Jaeger-Demo"
	jaegerEndPoint = "127.0.0.1:4318"
)

func setupTracer(ctx context.Context) (func(context.Context) error, error) {
	traceProvider, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(traceProvider)
	return traceProvider.Shutdown, nil
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
	tracerProvider := traceSDK.NewTracerProvider(
		traceSDK.WithResource(res),
		traceSDK.WithSampler(traceSDK.AlwaysSample()),
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)))
	return tracerProvider, nil
}

func main() {
	ctx := context.Background()
	shutdown, err := setupTracer(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = shutdown(ctx)
	}()

	testTracer(ctx)
}

func testTracer(ctx context.Context) {
	tracer := otel.Tracer("test-tracer")
	baseAttrs := []attribute.KeyValue{
		attribute.String("domain", "opendemo.com"),
		attribute.Bool("plagiarize", false),
		attribute.Int("code", 7),
	}
	ctx, span := tracer.Start(ctx, "parent-span", trace.WithAttributes(baseAttrs...))
	defer span.End()

	for i := range 10 {
		_, iSpan := tracer.Start(ctx, fmt.Sprintf("span-%d", i))
		time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
		iSpan.End()
	}
	fmt.Println("done!")
}
