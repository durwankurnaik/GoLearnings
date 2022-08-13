package main

import "fmt"

func main() {
	//fmt.Println("Lets talk about the pointers")
	//
	//age := 20
	//
	//var ptr *int = &age
	//fmt.Println(ptr)
	//
	//var ptrOfPtr **int = &ptr
	//fmt.Println(ptrOfPtr)
	//
	//doTheMath(age)
	//fmt.Println(age)

	var customPtr *custom
	customPtr = new(custom)
	(*customPtr).name = "Harshit Naik"
	customPtr.name = "Durwankur Naik"
	fmt.Println((*customPtr).name)
	fmt.Println(customPtr.name) // note that here if we see, we are trying to access name field of a pointer
	// but go this is actually a simpler way to dereference the pointer first and then access its field name
	// it is a convenience feature
}

//func doTheMath(num int) {
//	num *= 2
//}

type custom struct {
	name string
}
