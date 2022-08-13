package main

import "fmt"

func main() {
	//fmt.Println("Welcome to arrays in GoLang")
	//
	//var fruits [5]string
	//
	//fruits[0] = "Banana"
	//fruits[1] = "Apple"
	//fruits[2] = "Mango"
	//fruits[3] = "Guava"
	//fruits[4] = "Jack-fruit"
	//
	//fmt.Println(fruits)
	//
	//vegetables := [1]string{"beetroot"}
	//
	//fmt.Println(vegetables)

	grades := [5]int{3, 5, 6}

	fmt.Printf("%v, %T\n", grades, grades)
	fmt.Println(len(grades))
	fmt.Println(cap(grades))
}
