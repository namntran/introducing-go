package main

import (
	"fmt"
	"strings"
)

// replace a smaller string in a bigger string with some other string
func main() {
	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2)) // 2 indicates how many times to do the replacement
	// => "bbaa"

}
