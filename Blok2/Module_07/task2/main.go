package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomString(n int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	res := make([]rune, n)
	for i := range res {
		res[i] = chars[rand.Intn(len(chars))]
	}
	return string(res)
}

func interviewer(ch chan string, N int) {
	for i := 0; i < N; i++ {
		delay := rand.Intn(400) + 300 // 300–700 ms
		time.Sleep(time.Duration(delay) * time.Millisecond)

		opinion := randomString(5)
		ch <- opinion
	}
	close(ch)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	N := rand.Intn(10) + 1 // случайное количество жителей
	fmt.Println("Будет опрошено жителей:", N)

	ch := make(chan string)

	go interviewer(ch, N)

	for opinion := range ch {
		fmt.Println("Получено мнение:", opinion)
	}

	fmt.Println("Все мнения собраны, программа завершена")
}