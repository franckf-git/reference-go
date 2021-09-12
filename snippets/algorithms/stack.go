package main

import "fmt"

type Stack struct {
	slice []int
}

func (s *Stack) Push(add int) {
	s.slice = append(s.slice, add)
}

func (s *Stack) Pop() (first int) {
	if s.Size() == 0 {
		return -1
	}
	first = s.Peek()
	s.slice = s.slice[:s.Size()-1]
	return
}

func (s *Stack) Peek() int {
	if s.Size() == 0 {
		return -1
	}
	return s.slice[s.Size()-1]
}

func (s *Stack) Size() int {
	return len(s.slice)
}

func main() {
	var myST1 Stack
	// or like :
	var myST2 *Stack = new(Stack)
	myST1.Push(1)
	myST1.Push(2)
	myST2.Push(10)
	myST2.Push(5)
	fmt.Println(myST2.Peek())
	myST2.Push(6)
	myST2.Push(2)
	myST2.Push(1)
	myST1.Push(3)
	myST2.Push(2)
	fmt.Println(myST2.Peek())
	myST2.Push(1)
	myST2.Push(8)
	fmt.Println(myST2.Peek())

	fmt.Println(myST1.slice)
	fmt.Println(myST2)
	fmt.Println(myST2.slice)

	myST1.Pop()
	myST2.Pop()
	myST2.Pop()

	fmt.Println(myST1.slice)
	fmt.Println(myST2.slice)

	myST1.Pop()
	myST1.Pop()
	myST1.Pop()
	myST1.Pop()
	myST1.Pop()
	fmt.Println(myST1.Size())
	fmt.Println(myST1.Peek())
}
