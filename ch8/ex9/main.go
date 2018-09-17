package main

import (
	"fmt"
	"strings"
)

// split a string into a list of strings by separating a string using -
func main() {
	// func Split (s, sep string)[]string
	fmt.Println(strings.Split("a-b-c-d-d", "-"))
	// => [a b c d d]
}
