package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumber(ch chan int) {

	num := rand.Intn(100)

	ch <- num
}

func main() {

	rand.Seed(time.Now().UnixNano())

	const workers = 5

	ch := make(chan int)

	for i := 0; i < workers; i++ {
		go generateNumber(ch)
	}

	for i := 0; i < workers; i++ {
		num := <-ch
		fmt.Println("Получено число:", num)
	}
}