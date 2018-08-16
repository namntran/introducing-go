package main

import "fmt"

func makeEvenGenerator() func() uint { //closure example - write a function that returns another function.
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
func main() {
	nextEven := makeEvenGenerator() //(nextEven)returns a (makeEvenGenerator)function that generates even numbers.
	fmt.Println(nextEven()) // returns 0
	fmt.Println(nextEven()) // returns 2
	fmt.Println(nextEven()) // returns 4
}
