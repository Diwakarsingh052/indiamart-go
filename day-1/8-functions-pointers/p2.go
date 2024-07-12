package main

import "fmt"

// we are doing a failed update using pointer
func main() {
	var p *int // nil
	fmt.Println(&p)
	updateValue(p)
	fmt.Println(p) // x80687
}

func updateValue(p1 *int) {
	//*p1 = 1000
	fmt.Println(&p1)
	x := 10
	p1 = &x // p1 is different from p from the main function, changing the address here does not update p

}
