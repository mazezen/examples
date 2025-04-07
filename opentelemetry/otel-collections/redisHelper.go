package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
)

// RedisHelper struct holds the Redis client and tracer for Redis operations.
type RedisHelper struct {
	client *redis.Client
	tracer trace.Tracer
}

// NewRedisHelper creates a new RedisHelper with the provided Redis client and tracer.
func NewRedisHelper(client *redis.Client, tracer trace.Tracer) *RedisHelper {
	// Ensure OpenTelemetry tracing is enabled for Redis.
	if err := redisotel.InstrumentTracing(client); err != nil {
		fmt.Printf("Failed to enable redis otel: %v\n", err)
	}
	return &RedisHelper{
		client: client,
		tracer: tracer,
	}
}

// Set sets a key-value pair in Redis with OpenTelemetry tracing.
func (rh *RedisHelper) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	// Create a new span for the Redis Set operation.
	ctx, span := rh.tracer.Start(ctx, "redis-set")
	defer span.End()

	// Perform the Redis SET operation.
	err := rh.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		// Record any error that occurs during the Redis operation.
		span.RecordError(err)
	}
	return err
}

// Get retrieves a value from Redis for the given key with OpenTelemetry tracing.
func (rh *RedisHelper) Get(ctx context.Context, key string) (string, error) {
	// Create a new span for the Redis Get operation.
	ctx, span := rh.tracer.Start(ctx, "redis-get")
	defer span.End()

	// Perform the Redis GET operation.
	val, err := rh.client.Get(ctx, key).Result()
	if err != nil {
		// Record any error that occurs during the Redis operation.
		span.RecordError(err)
	}
	return val, err
}
