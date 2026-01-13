package main

import "fmt"

func main() {
	// E-02
	a := []string{"aa", "bb", "cc", "ee"}
	fmt.Println(a)
	UpdateSlice(a, "dd")
	GrowSlice(a, "ee")
	fmt.Println(a)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(FirstName, LastName string, Age int) Person {
	return Person{
		FirstName: "Sulaimon",
		LastName:  "Cole",
		Age:       10,
	}
}

func MakePersonPointer(FirstName, LastName string, Age int) *Person {
	return &Person{
		FirstName: "Lanre",
		LastName:  "Shittu",
		Age:       15,
	}
}

func UpdateSlice(data []string, value string) {
	data[len(data)-1] = value
	fmt.Println(data)
}

func GrowSlice(data []string, value string) {
	data = append(data, value)
	fmt.Println(data)
}

func PersonBuilder() {
	persons := make([]Person, 10000000)
	for range 10000000 {
		persons = append(persons, Person{
			FirstName: "Lil",
			LastName:  "Baby",
			Age:       40,
		})
	}
}
