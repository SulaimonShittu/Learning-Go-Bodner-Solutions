package main

import (
	"fmt"
)

func main() {
	// E-01
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sub1 := greetings[:2]
	sub2 := greetings[1:4]
	sub3 := greetings[3:]
	sub4 := greetings[3:5]
	fmt.Println(greetings, sub1, sub2, sub3, sub4)

	// E-02
	message := "Hi 👧 and 👦"
	msgRunes := []rune(message)
	fourthRune := msgRunes[4]
	fmt.Println(string(fourthRune))

	// E-03
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	firstInstance := Employee{
		"Sulaimon",
		"Shittu",
		22,
	}
	secondInstance := Employee{
		firstName: "Zami",
		lastName:  "Shittu",
		id:        17,
	}
	var thirdInstance Employee
	thirdInstance.firstName = "Sameeha"
	thirdInstance.firstName = "Shittu"
	thirdInstance.id = 14
}
