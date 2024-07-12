package main

import "fmt"

func updateVal(px *int) {
	abc := 20
	px = &abc
	*px = 100 //
}
func main() {
	x := 10

	updateVal(&x)
	fmt.Println(x)
}
