package main

import "fmt"

func addToList(s []int) {
	fmt.Printf("s %p\n", s)
	s = append(s, 888, 999) // updated the copy
	fmt.Printf("after append s %p\n", s)

}
func main() {
	list := []int{10, 20} // l = 2 , c =2
	fmt.Printf("list %p\n", list)

	addToList(list)
	fmt.Printf("after append list %p\n", list)
	fmt.Println(list)
}
