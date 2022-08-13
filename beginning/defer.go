package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer in GoLang")

	//Kind of like LIFO, not kind of, but literally LIFO
	defer fmt.Println("Entered first")
	defer fmt.Println("Entered second")
	defer fmt.Println("Entered third")
	defer fmt.Println("Entered fourth")
	defer fmt.Println("Entered fifth")

	fmt.Println("Last line of the program")
}
