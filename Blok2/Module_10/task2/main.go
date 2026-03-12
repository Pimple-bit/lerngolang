package main

import (
	"fmt"
	"sync"
)

func main() {

	letters := []string{}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for j := 0; j < 100; j++ {

				letter := fmt.Sprintf("Letter from fan %d", id)

				mu.Lock()
				letters = append(letters, letter)
				mu.Unlock()
			}

		}(i)
	}

	wg.Wait()

	fmt.Println("Letters received:", len(letters))
}