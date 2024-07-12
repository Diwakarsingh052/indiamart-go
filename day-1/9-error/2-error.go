package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrFileNotFound = errors.New("not in the root directory")

func main() {
	_, err := openFile("abc")
	if err != nil {
		if errors.Is(err, ErrFileNotFound) {
			fmt.Println("File not found in not found in root directory")
			// take some actions to fix issues or do specific logging for the error
			// to add more context
			log.Println(err)
			return
		}
		log.Println(err)
	}
}

func openFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		//%w wrap the additional error context message
		return nil, fmt.Errorf("%w %w", err, ErrFileNotFound)
	}
	return f, nil
}
