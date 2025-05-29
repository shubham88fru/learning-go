package main

import "fmt"

// generic function
func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1] // remove the last item
	return item, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.items) == 0
}

func (s Stack[T]) PrintAll() {
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	for i, item := range s.items {
		fmt.Printf("Item %d: %v\n", i, item)
	}
}

func main() {
	x, y := 1, 2

	x, y = swap(x, y)
	fmt.Println("x:", x, "y:", y)

	s, t := "hello", "world"
	s, t = swap(s, t)
	fmt.Println("s:", s, "t:", t)

	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.Push(4)
	intStack.PrintAll()
	poppedItem, ok := intStack.Pop()
	fmt.Println("Popped item:", poppedItem, "Success:", ok)

	stringStack := Stack[string]{}
	stringStack.PrintAll()
	stringStack.Push("hello")
	stringStack.Push("world")
	stringStack.PrintAll()
}
