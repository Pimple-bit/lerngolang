package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gardener(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	workTime := rand.Intn(500) + 500 // 500–1000 ms
	fmt.Printf("Gardener %d started working\n", id)

	time.Sleep(time.Duration(workTime) * time.Millisecond)

	fmt.Printf("Gardener %d finished after %d ms\n", id, workTime)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	gardenersCount := rand.Intn(5) + 1 // случайное кол-во
	fmt.Println("Gardeners working:", gardenersCount)

	for i := 1; i <= gardenersCount; i++ {
		wg.Add(1)
		go gardener(i, &wg)
	}

	wg.Wait()

	fmt.Println("All gardeners finished. Program completed.")
}