package main

import "fmt"

var age int = 20

func main() {
	//fmt.Println(age)
	//age := 22 // here, the age declaration shadows the package declaration and it is given more precedence
	//fmt.Println(age)

	//specialNum := "69"
	//extractNum, _ := strconv.ParseInt(specialNum, 10, 32)
	//
	//fmt.Printf("%v %T\n", extractNum, extractNum)
	//
	//age := 33
	//fAge := float32(age)
	//
	//fmt.Printf("%v, %T\n", fAge, fAge)

	var didMarried bool = true

	fmt.Printf("%v, %T\n", didMarried, didMarried)

	didMarried = false

	fmt.Println(didMarried)
}
