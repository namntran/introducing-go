package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 5 //store the int 5 in the memory location xPtr refers to.
}

func main() {
	xPtr := new(int) // new takes a type as an argument, allocate enough memory to fit a value of that type, and returns a pointer to it.
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}
