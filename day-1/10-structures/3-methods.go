package main

import "fmt"

// First, and most important, does the method need to modify the receiver? If it does, the receiver must be a pointer.
// Second is the consideration of efficiency. If the receiver is large, a big struct for instance, it may be cheaper to use a pointer receiver.
//
// Next is consistency. If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used.
// model data using structs
type user struct {
	Name     string // fields
	Password string
	age      int
	marks    []int
}

func (u *user) show() { // func(receiver)methodName(Args)returnTypes {}
	fmt.Printf("%+v\n", u)
}

func (u *user) updatePassword(password string) {
	u.Password = password
}

func main() {
	u1 := user{
		Name:     "raj",
		Password: "raj",
		age:      30,
		marks:    []int{10, 20, 40, 50},
	}
	u1.updatePassword("abc")
	u1.show()
}
