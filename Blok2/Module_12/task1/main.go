package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SensorData struct {
	Type      string
	SensorID  int
	Latitude  float64
	Longitude float64
	Value     float64
	Time      time.Time
}

func sensor(ctx context.Context, wg *sync.WaitGroup, sensorType string, id int, ch chan<- SensorData) {
	defer wg.Done()

	lat := rand.Float64()*180 - 90
	lon := rand.Float64()*360 - 180

	for {
		select {

		case <-ctx.Done():
			fmt.Printf("%s sensor %d stopped\n", sensorType, id)
			return

		default:

			data := SensorData{
				Type:      sensorType,
				SensorID:  id,
				Latitude:  lat,
				Longitude: lon,
				Value:     rand.Float64() * 100,
				Time:      time.Now(),
			}

			ch <- data

			time.Sleep(time.Duration(rand.Intn(700)+300) * time.Millisecond)
		}
	}
}

func startSensors(ctx context.Context, wg *sync.WaitGroup, sensorType string, count int, ch chan<- SensorData) {

	for i := 1; i <= count; i++ {
		wg.Add(1)
		go sensor(ctx, wg, sensorType, i, ch)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	dataCh := make(chan SensorData)

	var wg sync.WaitGroup

	parentCtx := context.Background()

	pressureCtx, cancelPressure := context.WithCancel(parentCtx)
	humidityCtx, cancelHumidity := context.WithCancel(parentCtx)
	seismicCtx, cancelSeismic := context.WithCancel(parentCtx)

	startSensors(pressureCtx, &wg, "pressure", rand.Intn(3)+2, dataCh)
	startSensors(humidityCtx, &wg, "humidity", rand.Intn(3)+2, dataCh)
	startSensors(seismicCtx, &wg, "seismic", rand.Intn(3)+2, dataCh)

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	go func() {

		time.Sleep(5 * time.Second)
		fmt.Println("\nSTOP PRESSURE SENSORS\n")
		cancelPressure()

		time.Sleep(5 * time.Second)
		fmt.Println("\nSTOP HUMIDITY SENSORS\n")
		cancelHumidity()

		time.Sleep(5 * time.Second)
		fmt.Println("\nSTOP SEISMIC SENSORS\n")
		cancelSeismic()

	}()

	for data := range dataCh {

		fmt.Printf(
			"[%s] sensor:%d (%.2f, %.2f) value=%.2f\n",
			data.Type,
			data.SensorID,
			data.Latitude,
			data.Longitude,
			data.Value,
		)
	}

	fmt.Println("\nAll sensors stopped. Meteocenter shutdown.")
}