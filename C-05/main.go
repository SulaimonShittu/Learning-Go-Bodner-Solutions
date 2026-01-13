package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	// E-01
	x1, err := div(5, 2)
	if err == nil {
		fmt.Println("The quotient of 5 / 2 is : ", x1)
	} else {
		log.Println(err)
	}
	x2, err := div(5, 0)
	if err == nil {
		fmt.Println("The quotient of 5 / 0 is : ", x2)
	} else {
		log.Println(err)
	}

	// E-02
	fileLength, err := fileLen("name.txt")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(fileLength)
	}

	// E-03
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

func div(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("Division by zero")
	}
	return num / den, nil
}

func fileLen(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	fStat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return int(fStat.Size()), nil
}

func prefixer(value string) func(string) string {
	return func(s string) string {
		return fmt.Sprintf("%s %s", value, s)
	}
}
