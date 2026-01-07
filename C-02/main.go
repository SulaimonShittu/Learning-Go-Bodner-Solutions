package main

import (
	"fmt"
	"math"
)

func main() {
	// E-01
	var i = 20
	var f = float64(i)
	fmt.Println(i, f)

	// E-02
	const value = 10
	i = value
	f = value

	// E-03
	var b byte
	var smallI int32
	var bigI uint64
	b = math.MaxInt8
	smallI = math.MaxInt32
	bigI = math.MaxUint64
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
}
