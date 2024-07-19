package main

import "fmt"

func main() {

	ch := make(chan int)

	close(ch)
	// ok would be false if channel is closed
	a, ok := <-ch
	if !ok {
		fmt.Println(a, "closed")
		return
	}
	fmt.Println(a)

}
