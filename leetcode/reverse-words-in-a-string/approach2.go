package main

import (
	"fmt"
)
//1. Split string return Array
//2. Reverse Array return array
//3. Join Array return string

func split(str string) []string {
	var result []string
	var word []rune 
	strRune := []rune(str)
	for _, char := range strRune {
		if char != '.' {
			word = append(word, char)
		} else {
			result = append(result, string(word))
			word = nil 
		}
	}
	if len(word) > 0 {
		result = append(result, string(word))
	}
	return result
}

func reverse(s []string) []string {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

func join(s []string) string {
	result := s[0]
	for i:=1; i<len(s); i++ {
		result += "." + s[i]
	}
	return result
}

func main() {
	s := "i.like.this.program.very.much"
	afterSplit := split(s)
	fmt.Println("After Split:", afterSplit)
	afterReverse := reverse(afterSplit)
	fmt.Println("After Reverse:", afterReverse)
	fmt.Println(join(afterReverse))
}
