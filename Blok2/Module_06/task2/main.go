package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(ch chan int) {

		sleep := (rand.Intn(10) + 1) * 100
		fmt.Println("Горутина 1 спит:", sleep, "ms")

		time.Sleep(time.Duration(sleep) * time.Millisecond)

		val := rand.Intn(100)
		ch <- val

	}(ch1)

	go func(ch chan int) {

		sleep := (rand.Intn(10) + 1) * 100
		fmt.Println("Горутина 2 спит:", sleep, "ms")

		time.Sleep(time.Duration(sleep) * time.Millisecond)

		val := rand.Intn(100)
		ch <- val

	}(ch2)

	time.Sleep(500 * time.Millisecond)

	select {

	case v := <-ch1:
		fmt.Println("Получено из ch1:", v)

	case v := <-ch2:
		fmt.Println("Получено из ch2:", v)

	default:
		fmt.Println("Сработал default — ни один канал не готов")
	}
}