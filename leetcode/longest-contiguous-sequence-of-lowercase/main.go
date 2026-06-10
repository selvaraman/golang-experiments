package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "AAabcBBxyZpqRSabcdT"

	s := []rune(str)

	start := -1
	longest := ""

	for i, ch := range s {
		if unicode.IsUpper(ch) {
			if start != -1 {
				sub := string(s[start+1 : i])
				if len(sub) > len(longest) {
					longest = sub
				}
			}
			start = i
		}
	}

	fmt.Println(longest)
}
