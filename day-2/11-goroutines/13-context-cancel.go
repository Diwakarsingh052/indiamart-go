package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func slowFunction() int {
	time.Sleep(5 * time.Second)
	return 42
}
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- slowFunction()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Operation timed out!")

	case res := <-ch:
		fmt.Println("Got value from slow function:", res)
	}

	wg.Wait()

}
