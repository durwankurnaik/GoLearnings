package main

import "fmt"

func main() {
	num := 10000000
	// it works upto 10^7, is varies according to the language as below
	// in C++, it works fine till 10^6
	// in Java, it works fine till 10^5
	// in Python, it is set to 1000 but can be easily changed upto 10^4 but not recommended
	fmt.Println(doItRecursively(num))
}

func doItRecursively(num int) int {
	if num == 1 {
		return 1
	}
	return doItRecursively(num - 1)
}
