package main

import (
	"fmt"
	"sync"
	"time"
)

var cab int = 1

func main() {
	var wg = &sync.WaitGroup{}
	var m = &sync.Mutex{}
	names := []string{"a", "b", "c", "d"}
	for _, name := range names {
		wg.Add(1)
		go bookCab(name, wg, m)

	}

	wg.Wait()

}

func bookCab(name string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	fmt.Println("welcome to the website", name)
	fmt.Println("some offers for you", name)
	//until the lock is not released
	//any read , write from other goroutines would not be allowed after lock is acquired
	m.Lock()
	defer m.Unlock()
	//critical section
	// we are protecting critical section from concurrent writes
	if cab >= 1 {
		fmt.Println("car is available for", name)
		time.Sleep(1 * time.Second)
		fmt.Println("booking confirmed", name)
		cab--
	} else {
		fmt.Println("car is not available for", name)
	}

}
