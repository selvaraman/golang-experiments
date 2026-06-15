package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	window := 3
	n := len(slice)
	fmt.Println("I/P:", slice)
	for i := 0; i < n; i += window {
		left := i
		var right int
		if i+window < n-1 {
			right = i + window - 1
		} else {
			right = n - 1
		}
		for left < right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}
	fmt.Println("O/P:", slice)
}
