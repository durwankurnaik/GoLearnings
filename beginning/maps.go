package main

import "fmt"

func main() {
	//fmt.Println("Welcome to Maps in GoLang")
	//
	//languages := make(map[string]string)
	//
	//languages["JS"] = "Javascript"
	//languages["Jv"] = "Java"
	//languages["Py"] = "Python"
	//languages["Go"] = "GoLang"
	//languages["Cpp"] = "C++"
	//
	//fmt.Println(languages)
	////fmt.Println(languages["Go"])
	//
	//delete(languages, "Cpp")
	//
	//fmt.Println(languages)
	//
	//for key, val := range languages {
	//	fmt.Printf("Key is %v and the value is %v\n", key, val)
	//}

	langProficiency := map[string]string{
		"Java":       "Very good",
		"Python":     "Good",
		"Javascript": "Good",
		"Typescript": "Beginner",
		"GoLang":     "Beginner",
	}

	fmt.Println(langProficiency)
}
