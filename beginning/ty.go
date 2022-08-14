package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

func bigFunction() {
	time.Sleep(2 * time.Second)
	fmt.Println("Operation done")
	waitGroup.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		waitGroup.Add(1)
		go bigFunction()
	}
	waitGroup.Wait()
}

// best use-case of goroutines I have seen so far in golang
