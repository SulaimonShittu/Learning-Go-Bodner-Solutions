package main

import (
	"fmt"
	"math"
)

func main() {

}

func exerciseTwoQ1() {
	// Question one
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
}

func exerciseTwoQ2() {
	// Question two
	const value = 12
	var i int = value
	var f float64 = value
	fmt.Println(i, f)
}

func exerciseTwoQ3() {
	// Question Three
	var (
		b      byte   = math.MaxUint8
		smallI int32  = math.MaxInt32
		bigI   uint64 = math.MaxUint64
	)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
