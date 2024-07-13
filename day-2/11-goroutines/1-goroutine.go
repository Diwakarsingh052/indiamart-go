package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	// counter to add number of goroutines
	// that we are starting or spinning up
	wg.Add(1)
	go hello(wg) // spawning a goroutine
	fmt.Println("end of the main")
	// this is the worst idea to wait
	//time.Sleep(time.Second)
	wg.Wait() //block until the counter is reset to 0
}

func hello(wg *sync.WaitGroup) {
	//decrement the counter
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("hello from the hello func")
}
