package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"
	"time"
)

const (
	serviceName     = "http-client"
	peerServiceName = "http-client-peer"
	jaegerEndpoint  = "127.0.0.1:4318"
	blogUrl         = "https://baidu.com"
)

func newJaegerTraceProvider(ctx context.Context) (*traceSDK.TracerProvider, error) {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
		otlptracehttp.WithInsecure())
	if err != nil {
		return nil, err
	}

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)))
	if err != nil {
		return nil, err
	}
	tracerProvider := traceSDK.NewTracerProvider(
		traceSDK.WithResource(res),
		traceSDK.WithSampler(traceSDK.AlwaysSample()),
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)))

	return tracerProvider, nil
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

	tr := otel.Tracer("http-client")

	ctx, span := tr.Start(ctx, "GET BLOG", trace.WithAttributes(semconv.PeerService(peerServiceName)))
	defer span.End()

	body, err := CallHttp(ctx, http.MethodGet, blogUrl, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
