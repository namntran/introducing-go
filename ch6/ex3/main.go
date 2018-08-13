package main

import "fmt"

func f() (int, int, int) { //go can return multiple values from a function, this function returns 3 integers
	return 5, 6, 7
}
func main() {
	x, y, z := f()
	fmt.Println(x, y, z)
}
