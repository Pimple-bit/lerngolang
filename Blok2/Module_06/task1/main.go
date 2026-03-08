package main

import (
	"fmt"
	"time"
)

func intProducer(ch chan int) {
	i := 0
	for {
		i++
		ch <- i
		time.Sleep(300 * time.Millisecond)
	}
}

func stringProducer(ch chan string) {
	i := 0
	for {
		i++
		ch <- fmt.Sprintf("msg %d", i)
		time.Sleep(1 * time.Second)
	}
}

func floatProducer(ch chan float64) {
	i := 0
	for {
		i++
		ch <- float64(i) * 1.1
		time.Sleep(5 * time.Second)
	}
}

func main() {

	intCh := make(chan int)
	strCh := make(chan string)
	floatCh := make(chan float64)

	go intProducer(intCh)
	go stringProducer(strCh)
	go floatProducer(floatCh)

	for {
		select {

		case v := <-intCh:
			fmt.Println("Получено из intCh:", v)

		case v := <-strCh:
			fmt.Println("Получено из strCh:", v)

		case v := <-floatCh:
			fmt.Println("Получено из floatCh:", v)
		}
	}
}