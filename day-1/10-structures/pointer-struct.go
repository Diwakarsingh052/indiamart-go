package main

import "fmt"

type list struct {
	num []int
}

func (l1 *list) addToList() {
	l1.num = append(l1.num, 40, 50)
}

func main() {

	l := &list{num: []int{10, 20, 30}} // l = 3, c=3

	l.addToList()

	fmt.Println(l)
}
