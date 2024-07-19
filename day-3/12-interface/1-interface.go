package main

import (
	"fmt"
)

type Reader interface {
	Read(b []byte) (int, error)
	// we need to implement all of the methods to impl the interface
	//abc()
}

type file struct {
	fileName string
}

func (f file) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them")
	return 0, nil
}

type json struct {
	fileName string
}

func (j json) Read(b []byte) (int, error) {
	fmt.Println("reading json files and processing them")
	return 0, nil
}
func DoWork(r Reader) {
	r.Read(nil)
}

func main() {
	var f file
	var j json

	DoWork(f)
	DoWork(j)

	//log.New()
	//io.ReadAll()

}
