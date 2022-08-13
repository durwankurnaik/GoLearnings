package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in GoLang")

	durwankur := User{"Durwankur Naik", 20, "Coding", 50_000}
	fmt.Println(durwankur.Name)

	durwankur.getName()
	durwankur.setName()
	fmt.Println(durwankur.Name)
}

type User struct {
	Name   string
	Age    int
	Hobby  string
	Salary int
}

func (u User) getName() {
	fmt.Println("The user name of current user is", u.Name)
}

func (u User) setName() {
	u.Name = "Harshit Naik"
	fmt.Println("Changed name in the function is", u.Name)
}
