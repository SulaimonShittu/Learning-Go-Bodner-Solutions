package main

import (
	"fmt"
)

func ExerciseThreeQ1() {
	greetings := []string{
		"Hello",
		"Hola",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}
	greetings2 := greetings[:2]
	greetings3 := greetings[1:4]
	greetings4 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(greetings2)
	fmt.Println(greetings3)
	fmt.Println(greetings4)
}

func ExerciseThreeQ2() {
	message := "Hi👧🏼 and 👦🏼"
	fourthRune := rune(message[3])
	fmt.Println(string(fourthRune))
}

func ExerciseThreeQ3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	sula := Employee{
		"Sulaimon",
		"Shittu",
		21,
	}
	zami := Employee{
		firstName: "Zami",
		lastName:  "Shittu",
		id:        16,
	}
	var anu Employee = Employee{}
	anu.firstName = "Anuoluwapo"
	anu.lastName = "Shittu"
	anu.id = 13

	fmt.Println(sula)
	fmt.Println(zami)
	fmt.Println(anu)
}
