package main

import (
	"fmt"
	"math/rand"
)

func ExerciseFourQ1() {
	numbers := make([]int, 0, 100)
	for range 100 {
		numbers = append(numbers, rand.Intn(100))
	}
}

func ExerciseFourQ2() {
	numbers := make([]int, 0, 100)
	for range 100 {
		numbers = append(numbers, rand.Intn(100))
	}
	for _, v := range numbers {
		switch {
		case v%2 == 0 && v%3 == 0:
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func ExerciseFourQ3() {
	var total int
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total)
	}
	// The bug is that for every iteration in the loop a new
	// total variable is declared which shadows the total
	// variable outside the for block. So for every iteration
	// total value is the value of i while the total variable
	// outside the for block remains dormant.
}
