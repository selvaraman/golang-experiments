package main

import "fmt"

func longestPalindromeFrom(str string, left, right int) string {
	for left >= 0 && len(str) > right && str[left] == str[right]  {
		left--
		right++
	}
	return str[left+1:right]
}

func longestPalindrome(str string) string {
	var longestStr string
	for i:=0; i<len(str); i++ {
		oddLengthStr := longestPalindromeFrom(str, i, i)
		evenLengthStr := longestPalindromeFrom(str, i, i+1)
		var temp string
		if len(oddLengthStr) > len(evenLengthStr) {
			temp = oddLengthStr
		}else{
			temp = evenLengthStr
		}
		if len(temp) > len(longestStr) {
			longestStr = temp
		}
	}
	return longestStr
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
