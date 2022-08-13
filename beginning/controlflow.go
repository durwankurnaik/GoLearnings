package main

import "fmt"

func main() {
	//if false {
	//	fmt.Println("True")
	//} else {
	//	fmt.Println("False")
	//}

	// i := 10  // we can even declare it like this and not have anything between switch keyword and curly braces
	//switch i := 10; {
	//case i <= 10:
	//	fmt.Println("Less than or equal to 10")
	//	fallthrough // in go, the break statement is implicitly available, so to fallthrough, we need to mention it like this
	//	// and even if the condition in the case does not satisfy, the next block of code will get executed
	//case i <= 20:
	//	fmt.Println("Less than or equal to 20")
	//}

	// we can even use type case syntax to compare types of the variables like below
	// it can be compared with any type of data type
	var i interface{} = [2]int{}
	switch i.(type) {
	case int:
		fmt.Println("It is a int")
	case float64:
		fmt.Println("It is a float64")
	case string:
		fmt.Println("It is a string")
	case [3]int:
		fmt.Println("It as a [3]int")
	default:
		fmt.Println("Other than mentioned int, float64 and string")
	}
}
