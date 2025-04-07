package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"time"
)

type HttpHelper struct {
	client *http.Client
	tracer trace.Tracer
}

func NewHttpHelper(tracer trace.Tracer) *HttpHelper {
	def := http.DefaultTransport
	defPot, ok := def.(*http.Transport)
	if !ok {
		panic("init transport 出错")
	}
	defPot.MaxIdleConns = 100
	defPot.MaxIdleConnsPerHost = 100
	defPot.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	return &HttpHelper{
		client: &http.Client{
			Timeout:   time.Minute,
			Transport: defPot,
		},
		tracer: tracer,
	}
}

func (h *HttpHelper) Get(ctx context.Context, url string, header map[string]string, params map[string]interface{}) ([]byte, error) {
	ctx, span := h.tracer.Start(ctx, "http-request-get")
	defer span.End()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			v := ToString(val)
			q.Add(key, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	r, err := h.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		span.RecordError(err)
		return nil, err
	}

	bb, err := io.ReadAll(r.Body)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return bb, err
}

func (h *HttpHelper) Post(ctx context.Context, url string, header map[string]string, params map[string]interface{}) (rr []byte, e error) {
	ctx, span := h.tracer.Start(ctx, "http-request-get")
	defer func() {
		if rr == nil && e == nil {
			fmt.Println("")
		}
		span.End()
	}()

	dd, _ := json.Marshal(params)
	re := bytes.NewReader(dd)
	req, err := http.NewRequest("POST", url, re)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	if header != nil {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	r, err := h.client.Do(req)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	defer r.Body.Close()
	if r.StatusCode != 200 {
		span.RecordError(errors.New(r.Status))
		return nil, errors.New(r.Status)
	}

	bb, err := io.ReadAll(r.Body)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}
	return bb, err
}
