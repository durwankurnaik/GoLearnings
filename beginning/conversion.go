package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the our pizza app")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please rate our pizza between 1 to 5")
	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Please enter a number")
	} else {
		fmt.Println("Thanks for rating us", numRating+1)
	}
}
