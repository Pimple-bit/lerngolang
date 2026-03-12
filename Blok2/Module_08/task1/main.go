package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "stopped:", ctx.Err())
			return
		default:
			fmt.Println(name, "working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func startGroup(ctx context.Context, groupName string, count int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go worker(ctx, fmt.Sprintf("%s-%d", groupName, i), wg)
	}
}

func main() {

	var wg sync.WaitGroup

	parentCtx, cancelParent := context.WithCancel(context.Background())
	middleCtx, cancelMiddle := context.WithCancel(parentCtx)
	childCtx, cancelChild := context.WithCancel(middleCtx)

	startGroup(parentCtx, "parentGroup", 2, &wg)
	startGroup(middleCtx, "middleGroup", 2, &wg)
	startGroup(childCtx, "childGroup", 2, &wg)

	time.Sleep(2 * time.Second)

	fmt.Println("\nCancel CHILD")
	cancelChild()

	time.Sleep(2 * time.Second)

	fmt.Println("\nCancel MIDDLE")
	cancelMiddle()

	time.Sleep(2 * time.Second)

	fmt.Println("\nCancel PARENT")
	cancelParent()

	wg.Wait()
}