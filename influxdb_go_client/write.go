package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/influxdata/influxdb-client-go/api/write"
// 	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
// )

// const (
// 	TOKEN = "LVbJ5lreeauQy8QpILuBoTrqArti8hTRtrQDSG4_xoQefYIPl9yqJFYBt03a2kO5GNgcgKU0diYIMnEnJCVsYg=="
// 	url = "http://127.0.0.1:8086"
// )
// func main() {
// 	client := influxdb2.NewClient(url, TOKEN)
// 	defer client.Close()

// 	org := "influxdb"
// 	bucket := "influxdb"
// 	writeAPI := client.WriteAPIBlocking(org, bucket)
// 	for value := 0; value < 5; value++ {
// 		tags := map[string]string{
// 			"tagname1": "tagvalue1",
// 		}
// 		fields := map[string]interface{}{
// 			"field1": value,
// 		}
// 		point := write.NewPoint("measurement1", tags, fields, time.Now())
// 		time.Sleep(1 * time.Second) // separate points by 1 second

// 		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }
