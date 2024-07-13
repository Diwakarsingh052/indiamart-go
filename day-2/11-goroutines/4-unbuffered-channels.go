package main

import "fmt"

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv

func main() {
	// unbuffered channel, because no size is specified
	ch := make(chan int) // unbuffered channel
	go func() {
		ch <- 20 // send signal
	}()

	// recv is a blocking call until there is no sender
	x := <-ch // recv signal
	fmt.Println(x)
}
