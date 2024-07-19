package main

import (
	"log"
)

type user struct {
	name  string
	email string
}

func main() {
	var u user
	l := log.New(u, "sales: ", log.Lshortfile)
	l.Println("hello")
}
