package main

import (
	"fmt"
	"strconv"
)

func main() {

}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | uint8 | uint16 | uint32 | uint64 | ~float32 | ~float64
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
	if s.Next == nil {
		s.Next = &SingleLinkedList[T]{
			Value: val,
		}
		return
	}
	s.Next.Add(val)
}

func (s *SingleLinkedList[T]) Insert(val T, pos int) {
	if pos == 0 {
		newSLL := *s
		s.Value = val
		s.Next = &newSLL
		return
	}
	element := s
	for range pos - 1 {
		element = element.Next
	}
	newSLL := &SingleLinkedList[T]{
		Value: val,
	}
	newSLL.Next = element.Next
	element.Next = newSLL
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
