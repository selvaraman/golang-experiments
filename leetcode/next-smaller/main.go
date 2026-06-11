package main

import "fmt"

func main() {
	input := []int{1, 3, 2, 4}
	output := []int{-1, -1, -1, -1}
	stack := []int{}
	for index, element := range input {
		for len(stack) > 0  && input[stack[len(stack)-1]] > element {
			top := stack[len(stack)-1]
			output[top] = element
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	fmt.Println("Input:", input)
	fmt.Println("Output:", output)
}
