package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5 //the zero function will not modify the original x variable in the main function
	zero(x) //call the zero function that takes an argument, that argument is copied to the function
	fmt.Println(x) //x is still 5
}
