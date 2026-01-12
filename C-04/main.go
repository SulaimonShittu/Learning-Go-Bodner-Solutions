package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// E-01
	randNumbers := make([]int, 0, 100)
	for range 100 {
		randNumbers = append(randNumbers, rand.IntN(101))
	}

	// E-02
	for _, v := range randNumbers {
		switch {
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		case v%6 == 0:
			fmt.Println("Six!")
		default:
			fmt.Println("Never mind")
		}
	}

	// E-03
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
}
