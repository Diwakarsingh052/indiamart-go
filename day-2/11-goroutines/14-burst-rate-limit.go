package main

import (
	"fmt"
	"sync"
	"time"
)

// MaxGoroutines is the maximum number of concurrent goroutines
const MaxGoroutines = 2

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// increment wg counter
		wg.Add(1)

		go func(i int) {
			// Schedule the call to Done to tell goroutine has completed
			defer wg.Done()

			// do your work
			fmt.Println("started goroutine ", i)
			time.Sleep(2 * time.Second)

		}(i)
	}
	// Wait for all goroutines to complete.
	wg.Wait()
}
