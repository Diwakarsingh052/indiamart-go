package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
}

func greet() {
	data := os.Args[1:]
	if len(data) != 3 {
		log.Println("please provide name , age , marks")
		return
	}

	name := data[0]
	ageString := data[1]
	//marksString := data[2]
	fmt.Println(name, ageString)
	//var err error
	//fmt.Println(err)
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println("error in conversion", err)
		return
	}
	fmt.Println(age)

}
