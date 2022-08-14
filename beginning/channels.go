package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		//i := <-ch
		//fmt.Println(i)  // here, we are sending two inputs in the sender channel but only receiving one input
		// that would cause us a panic, so we need to take all the inputs, we can do than using for in range loops
		for i := range ch {
			fmt.Println(i)
		}
		// after looping, we will close the channel in the sender channel
		//ch <- 14
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 69
		ch <- 100
		//i := <-ch
		//fmt.Println(i)
		// after doing the for loop, we need to close out channel from the sender channel
		close(ch)
		wg.Done()
	}(ch)
	// in the above example, both functions act as sender as receiver at the same time
	// but, we can strictly prohibit this behavior by giving them the parameter in the function
	// the commented out code shows that how the double way channels would have worked

	wg.Wait()
}
