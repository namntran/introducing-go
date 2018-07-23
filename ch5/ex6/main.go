package main

import "fmt"

func main() {
	arr := [5]float64{1, 2, 3, 4, 5} //create a slice with an underlying float64 array of length 5
	x := arr[0:4]                    //create a slice using the [low : high] expression - index
	fmt.Println(x)
}
