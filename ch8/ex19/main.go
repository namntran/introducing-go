package main

import (
	"fmt"
)

// Dereferencing a pointer means accessing the value of the variable which the pointer points to.
// *variableA is the syntax to deference variableA.

func main() {
	variableB := 255
	variableA := &variableB
	fmt.Println("address of variable b is", variableA)
	fmt.Println("value of variable b is (derefernce variable A and print the value of it)", *variableA) // deference variableA and print the value of it. As expected it prints the value of variableB.
	fmt.Println("value of veriable b", variableB)
}
