package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Nice to see you gofer"

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name сука блять:")
	name, _ := reader.ReadString('\n')

	fmt.Println(welcome, name)
	fmt.Printf("Type of the reader is %T\n", reader)
}
