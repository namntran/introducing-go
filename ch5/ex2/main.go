package main

import "fmt"

func main() {
	var x [5]float64 //first, create an array of length 5 to hold our test scores, then fill each element with a grade
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ { //second, set up a loop to compute the total score
		total += x[i]
	}
	fmt.Println(total / 5) //third, divide the total score by the number of elements to find the average
}
