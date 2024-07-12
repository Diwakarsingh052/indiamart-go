package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Attempt to open the file "text.txt", using `os.Open()`. This function returns two values: a pointer to a `File` object, and an `error`.
	f, err := os.Open("text.txt")
	if err != nil {
		log.Println(err)
		return
	}

	// `defer` is used to ensure that `f.Close()`, which is called to close the file, is executed after all other code in the function has finished running.
	// Regardless of what happens inside the function, the file will be closed before the function ends.
	defer func() {
		fmt.Println("running defer block")
		err := f.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()
}
