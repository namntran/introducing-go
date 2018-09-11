package main

import (
	"fmt"
	"strings"
)

// count the number of times a smaller string occurs in a bigger string, use the Count function:
func main() {
	// func Count(s, sep string) int
	fmt.Println(strings.Count("test", "t"))
	// => 2
}