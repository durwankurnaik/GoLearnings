package main

import (
	"fmt"
	"runtime"
)

//var waitGrp = sync.WaitGroup{}
//var mtx = sync.RWMutex{}
//var counter = 0

func main() {
	// first go routine
	//msg1 := "message1"
	//waitGrp.Add(2)
	//go func(msg string) {
	//	fmt.Println(msg)
	//	waitGrp.Done()
	//}(msg1)
	//
	//// second go routine
	//msg2 := "message2"
	//go func(msg string) {
	//	fmt.Println(msg)
	//	waitGrp.Done()
	//}(msg2)

	//for i := 0; i < 10; i++ {
	//	waitGrp.Add(2)
	//	mtx.RLock()
	//	go greet()
	//	mtx.Lock()
	//	go increment()
	//}
	//	// in the above example, because we are syncing the application my using mutex lock and unlock,
	//	// the concept of concurrency and parallelism have been completely destroyed
	//	// actually this application would perform even slower than normal execution of the program because
	//	// we are constantly locking and unlocking mutexes
	//	// so if we have to do something like this, we are better off not using the goroutines, but to do some tasks, this is very helpful
	//
	//waitGrp.Wait() // this line will wait for the above calls to be completed

	// time.Sleep(100 * time.Millisecond)  // using time decreases the performance because we don't exactly know how much time will go take to complete its execution
	// hence we may wait for more time that would decrease the performance
	// we may also wait for less time, that would lead to program not being executed properly
	// so, we can use wait groups from the sync package to do this work for us and finish the program just in time

	fmt.Println("Threads number on the system is", runtime.GOMAXPROCS(-1))
	// runtime.GOMAXPROCS() helps us to set the maximum number of threads in our go application
	// we can do it by like runtime.GOMAXPROCS(n) where n is the number of threads to be set, don't go too high
	// that would cause performance issues because of the scheduler need to maintain all those threads information
	// but, we can access that number by using -1 as the argument, that way is won't change the max number of threads
	// and will give us our result
}

//func greet() {
//	fmt.Printf("Hello go routines %v\n", counter)
//	mtx.RUnlock()
//	waitGrp.Done()
//}
//
//func increment() {
//	counter++
//	mtx.Unlock()
//	waitGrp.Done()
//}
