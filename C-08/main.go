package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := NewSLL(1)
	x.Add(2)
	x.Add(4)
	x.Add(6)
	x.Add(8)

	fmt.Println(x.Index(2))
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

func Doubler[T Number](val T) T {
	return 2 * val
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

func Printer[P Printable](value P) {
	fmt.Println(value.String())
}

type num int

func (n num) String() string {
	return strconv.Itoa(int(n))
}

type SingleLinkedList[T comparable] struct {
	Value T
	Next  *SingleLinkedList[T]
}

func NewSLL[T comparable](val T) *SingleLinkedList[T] {
	return &SingleLinkedList[T]{
		Value: val,
	}
}

func (s *SingleLinkedList[T]) Add(val T) {
	if s == nil {
		s = &SingleLinkedList[T]{
			Value: val,
		}
		return
	}
	s.Next.Add(val)
}

func (s *SingleLinkedList[T]) Insert(val T, pos int) {
	element := s
	for range pos {
		element = element.Next
	}
	element.Next = &SingleLinkedList[T]{
		Value: val,
	}
}

func (s *SingleLinkedList[T]) Index(val T) int {
	var count int
	element := s
	if val == element.Value {
		return count
	}
	for element.Next != nil {
		element = element.Next
		count++
		if val == element.Value {
			return count
		}
	}
	return -1
}
