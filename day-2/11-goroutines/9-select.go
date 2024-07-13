package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	done := make(chan struct{})

	wg.Add(3)
	go func() {
		defer wg.Done()
		c1 <- 1
		c1 <- 4
	}()

	go func() {
		defer wg.Done()
		c2 <- 2

	}()

	go func() {
		defer wg.Done()
		c3 <- 3

	}()

	//recv the values
	go func() {
		defer wg.Done()
		for {
			// whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			select {
			case x := <-c1:
				fmt.Println(x)
			case x := <-c2:
				fmt.Println(x)
			case x := <-c3:
				fmt.Println(x)
			case <-done:
				return
			}
		}

	}()

	wg.Wait()
	fmt.Println("main has to do some important")

	close(done)

	wg.Wait()

}
