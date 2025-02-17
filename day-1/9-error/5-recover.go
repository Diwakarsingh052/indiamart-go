package main

import (
	"fmt"
)

func main() {

	doSomething()
	fmt.Println("end of the main, stopped gracefully")
}

func doSomething() {
	defer recoveryFunc()
	//guarantee exec of recovery function because we are using defer
	var i []int
	i[100] = 1000
}

func recoveryFunc() {

	// The built-in `recover` function can stop the process of panicking,
	//if it is called within a deferred function.
	msg := recover()
	if msg != nil {
		// If `recover` captured a panic, it returns the panic value.
		// Here we print it.
		fmt.Println(msg)

		// `debug.Stack` gives us the stack trace at the point of the `recover` call,
		//helping us to identify where the problem occurred.
		//fmt.Println(string(debug.Stack()))

		// Acknowledge recovery by printing a message.
		// After this, the program will continue to run instead of crashing.
		fmt.Println("recovered from the panic")
	}
}
