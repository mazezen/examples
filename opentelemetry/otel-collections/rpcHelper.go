package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"sync/atomic"
)

type RPC struct {
	client *http.Client
	host   string
	tracer trace.Tracer
}

type ethRequest struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      string        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
}

var co = int32(0)

func init() {
	atomic.StoreInt32(&co, 0)
}

func (r *RPC) GetBlockNumber(ctx context.Context) (int64, error) {
	ctx, span := r.tracer.Start(ctx, "rpc-call-GetBlockNumber")
	defer span.End()

	body, err := r.call(ctx, "eth_blockNumber", nil)
	if err != nil {
		return 0, nil
	}
	var resp struct {
		Result string `json:"result"`
		Error  string `json:"error,omitempty"`
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		span.RecordError(err)
		return 0, err
	}
	count, err := HexToEthereumNumber(resp.Result)
	if err != nil {
		span.RecordError(err)
		return 0, err
	}
	return count.Int64(), nil
}

func (r *RPC) call(ctx context.Context, method string, params []interface{}) ([]byte, error) {
	ctx, span := r.tracer.Start(ctx, "rpc-call")
	defer span.End()

	data := ethRequest{
		Method:  method,
		Params:  params,
		Id:      "64",
		JSONRPC: "2.0",
	}
	body, err := json.Marshal(data)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", r.host, bytes.NewReader(body))
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	resp, err := r.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func NewRpc(url string, tracer trace.Tracer) *RPC {
	return &RPC{
		client: new(http.Client),
		host:   url,
		tracer: tracer,
	}
}
