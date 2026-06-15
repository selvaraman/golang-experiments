package main

import "fmt"

type MyInt int

func (mi MyInt) print() {
	fmt.Println("Success")
}

func main() {
	var nonStructType MyInt
	nonStructType.print()
}
