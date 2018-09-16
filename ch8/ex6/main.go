package main

import (
	"fmt"
	"strings"
)

// function that takes a list of strings and joins them together in a single string separated by another string
func main() {
	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	// => "a-b"
}
