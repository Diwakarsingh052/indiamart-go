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
	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		c1 <- 1
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		c2 <- 2
	}()
	go func() {
		defer wg.Done()
		c3 <- 3
		c3 <- 4
	}()

	// the loop has a fixed number of iterations, we are guessing, which is wrong!!
	for i := 1; i <= 3; i++ {
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

		}
	}
	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)

	wg.Wait()
}
