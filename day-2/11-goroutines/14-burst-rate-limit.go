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
	// buffered channel to limit concurrent goroutines
	semaphore := make(chan struct{}, MaxGoroutines)

	for i := 0; i < 10; i++ {
		// increment wg counter
		wg.Add(1)

		// acquire a semaphore slot before spawning a new goroutine
		semaphore <- struct{}{}
		go func(i int) {
			// Schedule the call to Done to tell goroutine has completed
			defer wg.Done()
			defer func() {
				// it might be useful in case any statement in this func might cause panic
				<-semaphore
			}()

			// do your work
			fmt.Println("started goroutine ", i)
			time.Sleep(2 * time.Second)
			time.Now()

		}(i)
	}
	// Wait for all goroutines to complete.
	wg.Wait()
}
