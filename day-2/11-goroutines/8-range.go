package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			go func(id int) {
				ch <- id
			}(i)
		}
	}()

	for v := range ch {
		fmt.Println("recv", v)
	}
}
