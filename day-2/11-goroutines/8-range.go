package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	// keeping track of number goroutines spawned
	wgWorker := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			wgWorker.Add(1)
			go func(id int) {
				defer wgWorker.Done()
				ch <- id
			}(i)
		}
		//we need to block our goroutine before closing the channel
		//because we want to make sure all the work
		// is done and finished // closing a channel will stop the for range loop
		wgWorker.Wait() // waiting until the worker goroutines are not finished
		close(ch)       //sending is finished over the channel ch
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		//ranging until the channel is not closed
		//range would receive all the remaining values even after the channel is closed
		for v := range ch {
			fmt.Println("recv", v)
		}

	}()

	wg.Wait()
}
