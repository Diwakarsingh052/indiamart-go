package main

import (
	"fmt"
)

// go mod init github.com/username/reponame
// global scope
var someData string

func main() {
	var s string // default value of string ""
	var i int = 10
	var name = "bob" // go is a statically compiled language
	//name = 100 not allowed
	sName, sAge := "raj", 12 // shorthand operator // it creates and assign the value

	fmt.Printf("%q\n", s)
	fmt.Println("hello world")
	fmt.Println(i, name, sName, sAge)

	var (
		//camelCase
		uName  string
		uAge   int = 15
		uMarks float64
	)

	fmt.Println(uName, uAge, uMarks)
	//time.Second // peek into it for design pattern
	//http.StatusInternalServerError
	//os.O_APPEND
}
