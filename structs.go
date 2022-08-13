package main

import "fmt"

// Person structs can be defined anywhere in the program, no need to strictly put them at the top of main function
type Person struct {
	name      string
	age       int
	hobbies   []string
	isMarried bool
}

func main() {
	durwankur := Person{
		name:      "Durwankur Naik",
		age:       20,
		hobbies:   []string{"Coding", "Problem Solving", "Typing", "Volleyball"},
		isMarried: false,
	}

	rahul := durwankur
	rahul.name = "Rahul Sawant"

	fmt.Println(rahul)
	fmt.Println(durwankur)

	// This is how a anonymous struct is been defined and initialized
	//aEngineer := struct {
	//	name string
	//	age  int
	//}{
	//	name: "Durwankur",
	//	age:  20,
	//}
	//
	//fmt.Println(aEngineer)
}
