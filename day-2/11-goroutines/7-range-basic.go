package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // we cant send values to the closed channel
		// we can recv the reaming values from the closed channel
		//ch <- 11 // this would panic because chan is closed
	}()

	// recv over the channel // it would run infinitely until the channel is close
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Done")
}
