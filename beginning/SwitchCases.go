package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch cases in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("The value is one, Fucked up luck")
	case 2:
		fmt.Println("The value is two, Not so good luck")
	case 3:
		fmt.Println("The value is three. OK")
	case 4:
		fmt.Println("The value is four, noice")
	case 5:
		fmt.Println("The value is five, nice luck I guess")
	case 6:
		fmt.Println("The value is six, you can roll again")
	default:
		fmt.Println("Something went wrong")
	}
}
