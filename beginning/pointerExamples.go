package main

import "fmt"

type Dog struct {
	name     string
	age      int
	breed    string
	cost     int
	strength int
	cuteness int
}

// in below, we can even pass Dg Dog, but that would only be pass by value, but we do want to change the values so we pass a pointer
// if we do something like:
// func(Dg Dog) calculateCuteness() {...}, that would also work, but the thing is that would not change the value of age
// that is being changed by this function below, but in the below case, that would change the actual value of age for that particular dog
// because the pointer to that doy is passed on
func (Dg *Dog) calculateCuteness() { // these are not called functions but are called as methods. they can be accessed by using . operator
	fmt.Printf("%v is %v out of 10 in cuteness\n", Dg.name, Dg.cuteness)
	(*Dg).age = 3 // if we write Dg.age = 3, that would also work but I just wanted to point this out
}

func main() {
	ruby := Dog{
		name:     "Ruby",
		age:      2,
		breed:    "Hybrid",
		cost:     10_000,
		strength: 6,
		cuteness: 9,
	}

	fmt.Println(ruby.age)
	ruby.calculateCuteness()
	fmt.Println(ruby.age)
}
