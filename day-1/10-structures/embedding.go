package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) updateEmail(email string) {
	u.email = email
}

type admin struct {
	user // embedding // anonymous field
	role []string
}

func (a admin) show() {
	fmt.Println(a)
}

func main() {
	a := admin{
		user: user{
			name:  "ajay",
			email: "ajay@email.com",
		},
		role: []string{"admin"},
	}
	a.updateEmail("abc@email.com")

	a.show()
}
