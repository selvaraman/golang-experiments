package main

import (
	"fmt"
	"slices"
)

func sortByFreq(s []int) []int {
	freq := make(map[int]int)
	for _, ch := range s {
		freq[ch]++
	}
	slices.SortFunc(s, func(a, b int) int {
		if freq[a] != freq[b] {
			return freq[b] - freq[a]
		}
		return a - b
	})
	return s 
}

func main() {
	input := []int{1, 2, 1, 3, 4, 2, 5, 7, 1, 3}
	fmt.Println(sortByFreq(input))
}
