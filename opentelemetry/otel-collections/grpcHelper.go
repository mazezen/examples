package main

import (
	"context"
	"errors"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)

type GrpcHelper struct {
	GrpcUrl string
	Proxy   string
	Client  *client.GrpcClient
	tracer  trace.Tracer
}

func NewGrpcHelper(proxy, url string, tracer trace.Tracer) (*GrpcHelper, error) {
	var grpcUrl string
	if len(grpcUrl) == 0 {
		grpcUrl = url
	}

	this := &GrpcHelper{
		GrpcUrl: grpcUrl,
		Proxy:   proxy,
	}

	dailFunc := func(addr string, d time.Duration) (conn net.Conn, err error) {
		if this.Proxy == "" {
			return net.DialTimeout("tcp", addr, d)
		}
		return nil, nil
		// 使用代理拨号器
	}

	c := client.NewGrpcClient(grpcUrl)

	dailOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(50*1024*1024),
			grpc.MaxCallSendMsgSize(50*1024*1024),
		),
		grpc.WithDialer(dailFunc),
	}
	err := c.Start(dailOptions...)
	if err != nil {
		return nil, err
	}

	this.Client = c
	this.tracer = tracer

	return this, nil
}

func (t *GrpcHelper) GetBlockHeight(ctx context.Context) (int64, error) {
	ctx, span := t.tracer.Start(ctx, "grpc-GetBlockHeight")
	defer span.End()

	if t.Client != nil {
		nowBlock, err := t.Client.GetNowBlock()
		if err != nil {
			span.RecordError(err)
			return 0, err
		}

		return nowBlock.BlockHeader.RawData.Number, nil

	}
	span.RecordError(errors.New("no start"))
	return 0, errors.New("no start")
}
