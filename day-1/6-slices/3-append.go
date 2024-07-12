package main

import "fmt"

func main() {
	i := []int{10, 20, 30, 40}

	i = append(i, 10, 20, 30)
	fmt.Println(i)
}
