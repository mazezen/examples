package main

import (
	"bytes"
	"context"
	"encoding/json"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io"
	"net/http"
)

type ClientOption struct {
	Transport *http.Transport
}

type ClientOptionFunc func(*ClientOption)

func WithClientTransport(transport *http.Transport) ClientOptionFunc {
	return func(option *ClientOption) {
		option.Transport = transport
	}
}

func CallHttp(ctx context.Context, method string, url string, rb io.Reader, headers http.Header, options ...ClientOptionFunc) ([]byte, error) {
	clientOption := &ClientOption{}
	for _, o := range options {
		o(clientOption)
	}

	client := http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
	if clientOption.Transport != nil {
		client.Transport = otelhttp.NewTransport(clientOption.Transport)
	}

	var reqBody io.Reader
	if rb != nil {
		payload, err := json.Marshal(rb)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(payload)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}
	if headers != nil {
		req.Header = headers
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
