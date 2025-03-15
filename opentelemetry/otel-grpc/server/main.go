// go install google.golang.org/protobuf/cmd/protoc-gen-go
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
// go get go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc
// go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc

package main

import (
	config2 "OpenTelemetry/otel-grpc/config"
	"OpenTelemetry/otel-grpc/pb/pb"
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

var tracer = otel.Tracer("grpc-server-demo")

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	_, span := tracer.Start(ctx, "SayHello",
		trace.WithAttributes(
			attribute.String("name", in.GetName()),
			attribute.StringSlice("client-id", md.Get("client-id")),
			attribute.StringSlice("user-id", md.Get("user-id")),
		),
	)
	defer span.End()
	return &pb.HelloResponse{Reply: "Hello " + in.GetName()}, nil
}

func main() {
	ctx := context.Background()
	tp, err := config2.InitTracer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
