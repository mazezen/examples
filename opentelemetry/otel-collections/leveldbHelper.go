package main

import (
	"context"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrEmptyKey = errors.New("key could not be empty")
)

// LevelDBHelper struct holds the LevelDB client and tracer for operations.
type LevelDBHelper struct {
	db     *leveldb.DB
	tracer trace.Tracer
}

// NewLevelDBHelper creates a new LevelDBHelper with the provided LevelDB client and tracer.
func NewLevelDBHelper(db *leveldb.DB, tracer trace.Tracer) *LevelDBHelper {
	return &LevelDBHelper{
		db:     db,
		tracer: tracer,
	}
}

// Put sets a key-value pair in LevelDB with OpenTelemetry tracing.
func (lh *LevelDBHelper) Put(ctx context.Context, key, value []byte) error {
	// Create a new span for the LevelDB Put operation.
	ctx, span := lh.tracer.Start(ctx, "leveldb-put")
	defer span.End()

	if len(key) < 1 {
		span.RecordError(ErrEmptyKey)
		return ErrEmptyKey
	}

	// Perform the LevelDB Put operation.
	err := lh.db.Put(key, value, nil)
	if err != nil {
		span.RecordError(err)
		return err
	}
	return err
}

// Get retrieves a value from LevelDB for the given key with OpenTelemetry tracing.
func (lh *LevelDBHelper) Get(ctx context.Context, key string) ([]byte, error) {
	// Create a new span for the LevelDB Get operation.
	ctx, span := lh.tracer.Start(ctx, "leveldb-get")
	defer span.End()

	// Perform the LevelDB Get operation.
	value, err := lh.db.Get([]byte(key), nil)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return value, nil
}
