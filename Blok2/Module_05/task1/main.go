package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			fmt.Println("Горутина 1: сообщение", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			fmt.Println("Горутина 2: сообщение", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			fmt.Println("Горутина 3: сообщение", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	wg.Wait()
}