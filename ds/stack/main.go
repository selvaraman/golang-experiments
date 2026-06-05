package main

import "fmt"

type Stack struct {
	data []int
}

func (stack *Stack) Push(data int) {
	stack.data = append(stack.data, data)
}

func (stack *Stack) Pop() (int, bool) {
	if len(stack.data) == 0 {
		return 0, false
	}
	length := len(stack.data)
	data := stack.data[length-1]
	stack.data = stack.data[:length-1]
	return data, true
}

func main() {
	stack := Stack {data: nil}
	for i:=0; i<10; i++ {
		stack.Push(i)
	}
	fmt.Printf("%+v\n", stack)
	data, ok := stack.Pop()
	fmt.Printf("Data: %d, Ok?: %t\n", data, ok)
	fmt.Println("Stack after pop")
	fmt.Printf("%+v\n", stack)
}
