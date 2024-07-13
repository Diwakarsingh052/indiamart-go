package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait() // We block here until all the goroutines call Done()

}

func work(id int, wg *sync.WaitGroup) int {
	defer wg.Done()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("work", id)
	}()

}
