package main

import "fmt"

func half(x int) (int, bool) { // function takes integer and halves it and returns true if even, or false if odd
	halved := x / 2

	if x%2 == 0 {
		return halved, true
	} else {
		return halved, false
	}
}

func main() {
	fmt.Println(half(1)) // half (1) should return 0, false
	fmt.Println(half(2)) // half (2) should return 1, true
	fmt.Println(half(3))
	fmt.Println(half(4))
	fmt.Println(half(5))
	fmt.Println(half(6))
	fmt.Println(half(7))
	fmt.Println(half(8))
	fmt.Println(half(9))
	fmt.Println(half(10))
}
