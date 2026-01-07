package main

import "fmt"

func main() {
	// E-01
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstSub := greetings[0:2:2]
	secondSub := greetings[1:4:4]
	thirdSub := greetings[3:]
	fmt.Println(greetings, firstSub, secondSub, thirdSub)

	// E-02

	// E-03
	type Employee struct {
		FirstName string
		LastName  string
		Id        int
	}
	sula := Employee{
		"Sulaimon",
		"Shittu",
		21,
	}
	zami := Employee{
		FirstName: "Zami",
		LastName:  "Shittu",
		Id:        17,
	}
	var anu Employee
	anu.FirstName = "Anuoluwapo"
	anu.LastName = "Shittu"
	anu.Id = 14
	fmt.Println(sula, zami, anu)
}
