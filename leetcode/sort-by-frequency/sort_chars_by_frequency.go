package main

import (
	"fmt"
	"slices"
)

func frequencySort(s string) string {
	runeStr := []rune(s)
	freq := make(map[rune]int)
	for _, ch := range runeStr {
		freq[ch]++
	}
	slices.SortFunc(runeStr, func(a, b rune) int {
		if freq[b] != freq[a] {
			return freq[b] - freq[a]
		}
		return int(a) - int(b)
	})
	return string(runeStr)
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}

/*
	Note: slices.SortFunc uses Pattern-defeating Quicksort
*/
