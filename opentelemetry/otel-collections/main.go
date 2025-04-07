package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/syndtr/goleveldb/leveldb"
	leveldbErrors "github.com/syndtr/goleveldb/leveldb/errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	otelEcho "go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	db              *xorm.Engine
	rdb             *redis.Client
	rdbHelper       *RedisHelper
	levelDBHelper   *LevelDBHelper
	httpUtilsHelper *HttpHelper
	httpTracer      trace.Tracer
	dbTracer        trace.Tracer
	rdbTracer       trace.Tracer
	levelTracer     trace.Tracer
	httpUtilsTracer trace.Tracer
	rpcTracer       trace.Tracer
	rpcHelper       *RPC
	grpcTracer      trace.Tracer
	grpcHelper      *GrpcHelper
)

// 创建 TracerProvider 并返回 shutdown 函数
func newTracerProvider(serviceName string) (*sdktrace.TracerProvider, trace.Tracer) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint("http://localhost:14268/api/traces"),
	))
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	return tp, tp.Tracer(serviceName)
}

// 初始化两个 TracerProvider：一个用于 HTTP，一个用于 DB
func initTracers() (func(context.Context) error, func(context.Context) error, func(context.Context) error,
	func(context.Context) error, func(context.Context) error, func(context.Context) error,
	func(context.Context) error) {
	// HTTP
	httpTP, httpT := newTracerProvider("go-http-service")
	otel.SetTracerProvider(httpTP)
	httpTracer = httpT

	// DB
	dbTP, dbT := newTracerProvider("go-db-service")
	dbTracer = dbT

	// redis
	redisTP, redisT := newTracerProvider("go-redis-service")
	rdbTracer = redisT

	// leveldb
	levelDBTP, levelDBT := newTracerProvider("go-leveldb-service")
	levelTracer = levelDBT

	// http utils request
	httpUtilsTP, httpUtilsT := newTracerProvider("go-http-utils-service")
	httpUtilsTracer = httpUtilsT

	// rpc
	rpcTP, rpcT := newTracerProvider("go-rpc-service")
	rpcTracer = rpcT

	// grpc
	grpcTP, grpcT := newTracerProvider("go-grpc-service")
	grpcTracer = grpcT

	return httpTP.Shutdown, dbTP.Shutdown, redisTP.Shutdown, levelDBTP.Shutdown, httpUtilsTP.Shutdown,
		rpcTP.Shutdown, grpcTP.Shutdown
}

func initDB() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/otel_demo?charset=utf8mb4"
	db, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	// 添加自动追踪 Hook
	db.AddHook(NewOtelXormHook(dbTracer))

	type User struct {
		ID   int64  `xorm:"pk autoincr"`
		Name string `xorm:"varchar(50)"`
	}
	_ = db.Sync2(new(User))
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:7379",
		Password: "asdasdzxc",
		DB:       0,
	})

	rdbHelper = NewRedisHelper(rdb, rdbTracer)
}

func initHttpRequest() {
	httpUtilsHelper = NewHttpHelper(httpUtilsTracer)
}

func initLevelDB() {
	ldb, err := leveldb.OpenFile("./leveldb_data", nil)
	if _, corrupted := err.(*leveldbErrors.ErrCorrupted); corrupted {
		ldb, err = leveldb.RecoverFile("./leveldb_data", nil)
	}
	if err != nil {
		panic(err)
	}
	levelDBHelper = NewLevelDBHelper(ldb, levelTracer)
}

func initRpc() {
	rpcHelper = NewRpc("https://api.shasta.trongrid.io/jsonrpc", rpcTracer)
}

func initGrpcHelper() {
	var err error
	grpcHelper, err = NewGrpcHelper("", "", grpcTracer)
	if err != nil {
		panic(err)
	}
}

// 接口处理函数
func handler(c echo.Context) error {
	ctx := c.Request().Context()
	ctx, span := httpTracer.Start(ctx, "handler-span")
	defer span.End()

	// Redis 设置一个 key
	err := rdbHelper.Set(ctx, "some-key", "hello-redis", time.Minute)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Redis Error")
	}

	// Redis 获取
	val, err := rdbHelper.Get(ctx, "some-key")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Redis Get Error")
	}

	// 模拟数据库操作
	type User struct {
		ID   int64
		Name string
	}
	newUser := &User{Name: fmt.Sprintf("user-%d", time.Now().Unix())}

	_, err = db.Context(ctx).Insert(newUser)
	if err != nil {
		span.RecordError(err)
		return c.String(http.StatusInternalServerError, "DB Error")
	}

	// 模拟leveldb操作
	err = levelDBHelper.Put(ctx, []byte("abc"), []byte("123"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB Put Error")
	}

	value, err := levelDBHelper.Get(ctx, "abc")
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB Get Error")
	}

	// 模拟HTTP发送请求
	res, err := httpUtilsHelper.Get(ctx, "https://www.baidu.com", nil, nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, "http request Error")
	}

	// 模拟rpc发送请求
	blockNumber, err := rpcHelper.GetBlockNumber(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "http request Error")
	}

	// 模拟grpc发送请求
	blockHeight, err := grpcHelper.GetBlockHeight(ctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, "grpc GetBlockHeight Error")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "OK",
		"user":         newUser,
		"redis_val":    val,
		"leveldb_val":  string(value),
		"http_val":     string(res),
		"block_number": blockNumber,
		"block_height": blockHeight,
	})
}

func main() {
	ctx := context.Background()
	shutdownHTTP, shutdownDB, shutdownRedis, shutdownLevelDB, shutdownHttpRequest, shutdownRpc, shutdownGrpc := initTracers()
	defer shutdownHTTP(ctx)
	defer shutdownDB(ctx)
	defer shutdownRedis(ctx)
	defer shutdownLevelDB(ctx)
	defer shutdownHttpRequest(ctx)
	defer shutdownRpc(ctx)
	defer shutdownGrpc(ctx)

	initDB()
	initRedis()
	initLevelDB()
	initHttpRequest()
	initRpc()
	initGrpcHelper()

	e := echo.New()
	e.Use(otelEcho.Middleware("go-http-service"))

	e.GET("/create", handler)

	log.Println("Listening on :8082")
	e.Logger.Fatal(e.Start(":8082"))
}
