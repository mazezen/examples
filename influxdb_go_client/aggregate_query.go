package main

import (
	"context"
	"fmt"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	TOKEN = "LVbJ5lreeauQy8QpILuBoTrqArti8hTRtrQDSG4_xoQefYIPl9yqJFYBt03a2kO5GNgcgKU0diYIMnEnJCVsYg=="
	url = "http://127.0.0.1:8086"
)

func main() {
	client := influxdb2.NewClient(url, TOKEN)
	defer client.Close()

	org := "influxdb"
	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "influxdb")
              |> range(start: -10m)
              |> filter(fn: (r) => r._measurement == "measurement1")
              |> mean()`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}

	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}