package main

import "fmt"

func sum(input ...int) int { // sum is a function that takes a slice of numbers and adds them together
	sum := 0

	for _, i := range input {
		sum += i
	}
	fmt.Println("sum is", sum)
	return sum
}

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8)
}
