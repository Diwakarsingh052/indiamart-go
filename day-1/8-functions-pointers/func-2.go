package main

import "fmt"

func add(x, y int) int {
	//fmt.Println(x + y)
	return x + y

}

func operate(op func(a, b int) int, x, y int) {
	fmt.Println(op(x, y))
}

func main() {
	operate(add(), 10, 20)
}
