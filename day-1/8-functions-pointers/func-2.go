package main

import "fmt"

type addition func(x, y int) int

func add(x, y int) addition {
	//fmt.Println(x + y)
	return func(x, y int) int {
		return x + y
	}

}

func operate(op func(a, b int) int, x, y int) {
	fmt.Println(op(x, y))
}

func main() {
	operate(add(10, 20), 10, 20)
}
