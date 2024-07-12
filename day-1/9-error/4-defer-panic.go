package main

import "fmt"

func main() {
	//"defer"  function's execution is postponed until the surrounding function returns.
	// Defer maintains a stack, meaning functions are executed in Last-In-First-Out (LIFO) order.

	defer fmt.Println(1)
	defer fmt.Println(2)

	doFunc()

	fmt.Println("end of the main")

}

func doFunc() {
	panic("some kind of msg")
}
