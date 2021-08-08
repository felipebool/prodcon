package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			select {
			// msg from other goroutine finish
			case <-ctx.Done():
				fmt.Println("done 1")
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			// msg from other goroutine finish
			case <-ctx.Done():
				fmt.Println("done 2")
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		// your operation
		// call cancel when this goroutine ends
		cancel()
	}()
	wg.Wait()
}
