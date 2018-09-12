package main

import (
	"fmt"
	"strings"
)

func main() { // determine if a bigger string starts with a smaller string, use the HasPrefix function
	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))
	// => true
}
