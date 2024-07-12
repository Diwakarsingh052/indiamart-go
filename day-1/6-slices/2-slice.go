package main

import "fmt"

func main() {
	//var i []int // nil
	i := []int{10, 20, 30, 40}
	fmt.Println(i)
	// this will cause panic as length is not available to store the value
	i[6] = 100 //update ops // it would not grow the slice
	fmt.Println(i)
	fmt.Printf("%#v", i)
}
