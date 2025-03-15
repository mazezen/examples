// go get github.com/redis/go-redis/extra/redisotel/v9

package main

import (
	"context"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
	"sync"
	"time"
)

const (
	serviceName    = "redis-Jaeger-Demo"
	jaegerEndpoint = "127.0.0.1:4318"
)

var tracer = otel.Tracer("redis-demo")

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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:7379",
		Password: "asdasdzxc",
		DB:       4,
	})

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	if err = redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	ctx, span := tracer.Start(ctx, "doSomething")
	defer span.End()

	if err = doSomething(ctx, rdb); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

func doSomething(ctx context.Context, rdb *redis.Client) error {
	if err := rdb.Set(ctx, "name", "otel-redis", time.Minute).Err(); err != nil {
		return err
	}
	if err := rdb.Set(ctx, "tag", "Otel", time.Minute).Err(); err != nil {
		return err
	}
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			val := rdb.Get(ctx, "tag").Val()
			if val != "Otel" {
				log.Printf("%q != %q", val, "Otel")
			}
		}()
	}
	wg.Wait()

	if err := rdb.Del(ctx, "name").Err(); err != nil {
		return err
	}
	if err := rdb.Del(ctx, "tag").Err(); err != nil {
		return err
	}
	log.Println("done!")
	return nil
}

func newJaegerTraceProvider(ctx context.Context) (*traceSDK.TracerProvider, error) {
	exp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(jaegerEndpoint),
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
		traceSDK.WithBatcher(exp, traceSDK.WithBatchTimeout(time.Second)))

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
