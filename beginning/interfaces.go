package main

import "fmt"

func main() {
	//obj := ConsoleWriter{}
	//obj.Write([]byte("This is using interface"))

	myInt := IntCounter(0)
	var inc Increment = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// interfaces contains the method declaration only and not the method body
// we don't need to explicitly tell go to implement, it automatically does that by checking our method signature
// to do that, we actually need a struct because we can't create methods without structs

//type Writer interface {
//	WriteIt([]byte) (int, error)
//}
//
//type ConsoleWriter struct{}
//
//func (cw ConsoleWriter) Write(data []byte) (int, error) {
//	n, err := fmt.Println(string(data))
//	return n, err
//}

// we can also implement interface without using structs like below

type Increment interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
