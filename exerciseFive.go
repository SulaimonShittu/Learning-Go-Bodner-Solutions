package main

import (
	"errors"
	"io"
	"log"
	"os"
)

func Div(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero")
	}
	return dividend / divisor, nil
}

func fileLen(fileName string) (int, error) {
	fil, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer fil.Close()
	size, err := io.Copy(io.Discard, fil) // Copies content to Discard, just counting bytes
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return int(size), nil
}

func prefixer(value string) func(string) string {
	return func(s string) string {
		return value + s
	}
}
