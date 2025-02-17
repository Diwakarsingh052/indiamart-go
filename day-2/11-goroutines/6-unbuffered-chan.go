package main

import (
	"fmt"
	"sync"
)

//A send on an unbuffered channel can proceed if a receiver is ready. send will block until there is no recv
//channel -> unbuffered chan, buffered chan

// A send on a buffered channel can proceed if there is room in the buffer

var wg = &sync.WaitGroup{}

func main() {
	//channel helps to send signals and data from one go routine boundary to another go routine
	c := make(chan int) // unbuffered chan
	//wg.Add(3)

	go addNum(10, 20, c)
	go mult(10, 10, c)
	go sub(100, 90, c)
	calcAll(c)

	//wg.Wait()
}

func addNum(a, b int, c chan int) {
	//defer wg.Done()
	sum := a + b
	// send this to channel
	// in case of an unbuffered chan , receiver must be ready otherwise send will block
	// send operation signal on the channel  // signaling with data
	c <- sum
}

func sub(a, b int, c chan int) {
	//defer wg.Done()
	diff := a - b

	// send
	c <- diff
}

func mult(a, b int, c chan int) {
	//defer wg.Done()
	prod := a * b

	// send
	c <- prod
}

func calcAll(c chan int) {
	//defer wg.Done()
	//recv from the channel // block operation until we don't recv the values
	x, y, z := <-c, <-c, <-c // it blocks until we recv all the three values
	//TODO:// recv all the values
	fmt.Println(x, y, z)

}
