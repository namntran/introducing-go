package main

import "fmt"

func add(args ...int) int { //add is allowed to be called with multiple integers, known as a variadic parameter
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println()
}
