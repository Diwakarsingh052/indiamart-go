package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	RateLimitSec    = 5 // duration in seconds
	MaxRequestsRate = 1 // M, max requests per RateLimitSec duration
)

func main() {
	var wg sync.WaitGroup

	// Ticker for rate limit
	//the formula time.Second * RateLimitSec / MaxRequestsRate is calculating
	//the time duration between each tick, in seconds.
	//This allows you to control the rate at which requests are processed,
	//as each request can only be processed after a tick from the ticker.
	ticker := time.NewTicker(time.Second * RateLimitSec / MaxRequestsRate)

	// Simulate requests
	for i := 0; i < 100; i++ {
		wg.Add(1)

		// If the requests come in faster than the ticker, they will be
		// blocked here until the next tick.

		// A request can be processed.
		go func(i int) {
			defer wg.Done()
			processRequest(i)
		}(i)
		fmt.Println(<-ticker.C, "ticker ticked")
	}

	// Wait for all requests to finish before exiting.
	wg.Wait()

	// Stop the ticker to release its resources.
	ticker.Stop()
}

func processRequest(i int) {
	// Simulate request processing
	fmt.Printf("Processing request %d at %s\n", i, time.Now())
	//time.Sleep(1 * time.Second)
	fmt.Printf("Finished processing request %d\n", i)
}
