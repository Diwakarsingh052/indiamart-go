package main

import "fmt"

// we are doing a failed update using pointer
func main() {
	var p *int // nil
	fmt.Println(&p)
	updateValue(&p)
	fmt.Println(*p)
}

func updateValue(p1 **int) {
	//*p1 = 1000
	fmt.Println(&p1)
	x := 10

	*p1 = &x
}
