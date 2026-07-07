package main

import "fmt"

func main() {
	str := "golang"
	fmt.Println(string(str[0]))
	str[0] = 't' //won't work
}
