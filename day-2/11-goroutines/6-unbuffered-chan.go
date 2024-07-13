package main

import "sync"

var wg = &sync.WaitGroup{}

func main() {
	//channel helps to send signals and data from one go routine boundary to another go routine
	c := make(chan int) // unbuffered chan
	//wg.Add(3)

	go addNum(10, 20, c)
	go mult(10, 10, c)
	go sub(100, 90, c)

}
func addNum(a, b int, c chan int) {
	//defer wg.Done()
	sum := a + b

	// send operation signal on the channel  // signaling with data
}

func sub(a, b int, c chan int) {
	//defer wg.Done()
	diff := a - b

	// send

}

func mult(a, b int, c chan int) {
	//defer wg.Done()
	prod := a * b

	// send
}

func calcAll(c chan int) {

	//recv from the channel
}
