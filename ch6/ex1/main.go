// write programs that use more than one function
// a function is an independent section of code that maps input parameters to output parameters
package main

import "fmt"

func average(xs []float64) float64 { // take code from our main function and move it into our average function
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs)) // the return statement causes the function to immediately stop and return the value after it to the function that called this one.
}
func main() {
	xs := []float64{98, 93, 77, 82, 83} // xs = list of scores. This is the calling function.
	fmt.Println(average(xs))
}

/* s
func average(xs []float64) float64 {
	panic("Not Implemented")
}
func main() {
	xs := []float64{98, 93, 77, 82, 83} // xs = list of scores

	total := 0.0
	for _, v := range xs {
		total += v
	}
	fmt.Println(total / float64(len(xs)))
}
*/
