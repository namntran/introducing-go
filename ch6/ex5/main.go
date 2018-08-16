package main

import "fmt"

func main() {
	add := func(x, y int) int { //create functions inside of functions. add is a function that takes 2 ints and returns an int.got
		return x + y
	}
	fmt.Println(add(1, 1))
}
