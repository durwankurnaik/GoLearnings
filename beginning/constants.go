package main

import (
	"fmt"
)

//const (
//	a = iota
//	b = iota
//	c = iota
//)

// this block is same as the above block
const (
	a = iota
	b
	c
)

const (
	a1 = iota // here, the iota is reset to 0 because iota is only changing in blocked scope
)

func main() {
	//const num int = 42
	//const anotherNum = 33
	//fmt.Printf("%v, %T\n", num, num)
	//fmt.Printf("%v, %T\n", anotherNum, anotherNum)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%v\n", a1)
}
