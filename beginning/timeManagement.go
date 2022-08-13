package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()

	createdTime := time.Date(2002, time.March, 14, 16, 34, 45, 23, time.Local)

	fmt.Println(presentTime, "is your current time")
	fmt.Println(createdTime, "is your created time")
}
