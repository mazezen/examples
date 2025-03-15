package main

import (
	"OpenTelemetry/otel-grpc/config"
	"OpenTelemetry/otel-grpc/pb/pb"
	"context"
	"flag"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

var tracer = otel.Tracer("grpc-client-demo")

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:7777", "the address to connect to")
	name = flag.String("name", defaultName, "name to say hello")
)

func main() {
	flag.Parse()

	tp, err := config.InitTracer(context.Background())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("error shutting down tracing provider: %v", err)
		}
	}()

	var conn *grpc.ClientConn
	conn, err = grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		log.Fatalf("did not conn: %s", err)
	}
	defer func() { _ = conn.Close() }()

	c := pb.NewGreeterClient(conn)
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "hello-client-echo",
		"user-id", "88")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx, span := tracer.Start(ctx, "c.SayHello", trace.WithAttributes(attribute.String("name", "gRPC-client")))
	defer span.End()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Hellp: %s", r.GetReply())
}

func callSayHello(c api.HelloServiceClient) error {
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
		"client-id", "web-api-client-us-east-1",
		"user-id", "some-test-user-id")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	response, err := c.SayHello(ctx, &api.HelloRequest{Greeting: "World"})
	if err != nil {
		return fmt.Errorf("calling SayHello: %w", err)
	}
	log.Printf("response from server: %s", response.Reply)
	return nil
}
