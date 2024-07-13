package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func slowFunction() int {
	time.Sleep(5 * time.Second)
	fmt.Println("slow fn ran")
	return 42
}
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	// creating an empty container
	ctx := context.Background()
	//setting a timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	// clean the resources taken up by the context
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case ch <- slowFunction():
			fmt.Println("value sent from the goroutine")
		case <-ctx.Done():
			fmt.Println("can't send the values")
			return
		}

	}()

	select {
	//checking if context was timed out
	case <-ctx.Done():
		fmt.Println("Operation timed out!", ctx.Err())
		return

		//recv over the channel
	case res := <-ch:
		fmt.Println("Got value from slow function:", res)

	}

	wg.Wait()

}
