package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	const (
		influxURL   = "http://localhost:8086"
		influxToken = "admin:admin123"
		org         = "my-org"
		bucket      = "my-bucket"
	)

	client := influxdb2.NewClient(influxURL, influxToken)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(org, bucket)

	for i := 0; i < 10; i++ {
		deviceID := fmt.Sprintf("sensor-%d", rand.Intn(5)+1)
		temperature := 20.0 + rand.Float64()*10.0
		humidity := 40.0 + rand.Float64()*20.0

		p := influxdb2.NewPointWithMeasurement("device_data").
			AddTag("device_id", deviceID).
			AddField("temperature", temperature).
			AddField("humidity", humidity).
			SetTime(time.Now())

		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			fmt.Println("Erro ao escrever no InfluxDB:", err)
		} else {
			fmt.Printf("Enviado: %s temp=%.2f hum=%.2f\n", deviceID, temperature, humidity)
		}

		time.Sleep(1 * time.Second)
	}
}