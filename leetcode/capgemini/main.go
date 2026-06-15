package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	window := 3
	var result [][]int
	for i := 0; i < len(input); i += window {
		end := i + window
		if end > len(input) {
			end = len(input)
		}
		result = append(result, input[i:end])
	}
	fmt.Println("Input:", input)
	fmt.Println("Output: ", result)
}
