package main

import "fmt"

func main() {
	// convert a string to a slice of bytes
	arr := []byte("test")
	string := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr, string)
	// => [116 101 115 116] test
}
