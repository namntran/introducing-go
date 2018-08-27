package main

import "fmt"

// write a function with one variadic parameter that finds the greatest number in a list of numbers.
func greatestNumber(args ...int) int { //greatestNumber is allowed to be called with multiple integers, known as a variadic parameter
	max := int(0) //assume first value is the smallest
	for _, arg := range args {
		if arg > max {
			max = arg // found another larger value, replace previous value in max
		}
	}
	return max
}
func main() {
	fmt.Println(greatestNumber(1, 2, 4, 7, 107, 9, 12, 15, 18))
}
