package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	wg := sync.WaitGroup{}
	// we would close the done channel when all goroutines are finished sending
	done := make(chan struct{})
	// To keep track of if the goroutine work is finished or not,
	// we need another goroutine to close the channel (done).
	// This pattern is useful in case of multiple producers sending data on
	// the channel and single consumer reading from it. Closing the channel
	// signals the consumer about the completion of data sending by all producers.
	wgWorker := sync.WaitGroup{}

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		time.Sleep(2 * time.Second)
		c1 <- 1

	}()
	go func() {
		defer wgWorker.Done()
		time.Sleep(1 * time.Second)
		c2 <- 2
		c2 <- 4

	}()
	go func() {
		defer wgWorker.Done()
		c3 <- 3

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		wgWorker.Wait()
		//close is signaling without data , it would send the recv signal to the select
		close(done)
	}()

	wg.Add(1)
	go func() {
		// Infinite for loop with select is useful here because we don't know how many value
		// we would receive from the channels
		// whatever channel sends the data first, that case will execute.
		// This gives us concurrency and we don't have to wait for a particular channel to finish.
		// Also, it allows the easy addition of more channels in the future.
		defer wg.Done()
		// the loop has a fixed number of iterations, we are guessing, which is wrong!!
		for {
			select {
			// whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			case x := <-c1:
				fmt.Println(x)
			case x := <-c2:
				fmt.Println(x)
			case x := <-c3:
				fmt.Println(x)
				time.Sleep(6 * time.Second)
			case <-done:
				fmt.Println("all goroutines finished sending")
				return
			}
		}
	}()
	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)

	wg.Wait()
}
