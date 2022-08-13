package main

import "fmt"

func main() {
	//fmt.Println("Welcome to loop breaks in GoLang")
	//
	//days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	//
	//fmt.Println(days)
	//
	////for i := 0; i < len(days); i++ {
	////	fmt.Println(days[i])
	////}
	//
	//for i := range days {
	//	fmt.Println(days[i])
	//}

	keyValuePair := map[int]int{
		36: 1,
		22: 2,
		46: 3,
		83: 4,
		52: 5,
	}

	// for each loop can be used to iterate over each and every collection in go and the syntax is always like below
	// it returns key and value for each collection, we can take anything according to our convenience
	for _, vals := range keyValuePair {
		fmt.Println(vals)
	}
}
