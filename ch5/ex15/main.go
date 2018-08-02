package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := x[0] // assume first value is the smallest

	for _, value := range x {
		if value < min {
			min = value // found another smaller value, replace previous value in min
		}
	}
	fmt.Println("The smallest value is : ", min)
}
