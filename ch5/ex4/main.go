package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83} //first, create an array of length 5 to hold our test scores, then fill each element with a grade

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
