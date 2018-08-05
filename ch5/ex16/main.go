package main

import "fmt"

func main() {
	x := []int{ //create a slice of ints
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	max := x[0] //assume first value is the smallest

	for _, value := range x {
		if value > max {
			max = value // found another larger value, replace previous value in max
		}
	}
	fmt.Println("The largest value is ", max)

}
