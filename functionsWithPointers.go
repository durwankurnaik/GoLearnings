package main

import "fmt"

func main() {
	//name := "Durwankur"
	//doggos := [2]string{"Ruby", "Simba"}
	//
	//fmt.Println(name, doggos)
	//changeNames(&name, &doggos)
	//fmt.Println(name, doggos)

	//takeAll(1, 2, 3, 4, 5, 6, 7, 8)

	num1 := 4.69
	num2 := 2.54
	fmt.Println(divideThoseFloats(num1, num2))

	ans, err := divideIt(5.0, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ans)
}

//func changeNames(name *string, doggos *[2]string) {
//	*name = "Harshit"
//	(*doggos)[0] = "Oscar"
//	(*doggos)[1] = "Ruby"
//	// here, we can also change the values like doggos[0] = "oscar"
//	// because compiler will actually do the same thing that we did it here
//	// this is just a syntactic sugar for ease of developers so that they don't have to always put in parenthesis and dereference it
//	// but when its not so complex declaration like for int and string, we need to explicitly dereference those pointers
//	// we can't do something like name = "harshit", that would just give us compile time error
//
//	// this behavior is not available for slices and maps because actually under the hood, they are using arrays
//	// so they point towards memory location of the starting index of the array
//	// so they are always passed in as a pass by reference and not pass by value
//}

// variatic parameters example
// here, it would take all the values and then it would convert it into slice and give it the name that we specify
// also the variatic parameters, in this case the values parameter should be at the end, additional parameters
// can be defined before the variatic parameters like func takeAll(message string, values ...int) {...}
//func takeAll(values ...int) {
//	var result int
//	for _, val := range values {
//		result += val
//	}
//	fmt.Println(result)
//}

// this type is called named return, this is done to make the code look cleaner but don't do it if the function are big
// because it would make it hard to read what you are returning
//func sumAll(values ...int) (result int) {
//	for _, val := range values {
//		result += val
//	}
//	return
//}

// functions with multiple return types
func divideThoseFloats(num1, num2 float64) float64 {
	return num1 / num2
}

// to tackle something like zero division error, we can do multiple return types
func divideIt(num1, num2 float64) (float64, error) {
	if num2 == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}
