package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s := "i.like.this.program.very.much"
	strSlice := strings.Split(s, ".")
	slices.Reverse(strSlice)
	result := strings.Join(strSlice, ".")
	fmt.Println("I/P:", s)
	fmt.Println("O/P:", result)
}
