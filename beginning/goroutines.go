package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to go routines in GoLang")

	go greeter("Durwankur")
	greeter("Naik")
}

func greeter(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println("Hii", name)
	}
}
