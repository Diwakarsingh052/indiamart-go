package main

import "fmt"

// This function only sends to the channel
func sender(c chan<- string) {
	c <- "Hello, world!"
	close(c)
}

// This function only receives from the channel
func receiver(c <-chan string) {
	msg := <-c

	fmt.Println(msg)
}

func main() {
	// Make a channel
	messageChan := make(chan string)

	// Start a goroutine with sender, give it the channel
	go sender(messageChan)

	// Start a goroutine with receiver, give it the channel
	receiver(messageChan)

	fmt.Println("Main function has completed")
}
