package main

import "fmt"

func zero(xPtr *int) { // pointer represented using an asterisk(*) followed by type of stored value. xPtr is pointer to an int (*int)
	*xPtr = 7 //store the int 7 in the memory location xPtr refers to.
	// asterisk used to dereference pointer variables, gives us access to the value the pointer points to.
	// pointers reference a location in memory where a value is stored rather than the value itself.
	// by using a pointer (*int), the zero function is able to modify the original variable in the main function.
}

func main() {
	x := 5
	zero(&x) // & operator finds the address of a variable. &x returns *int (pointer to an int) because x is an int.
	// this allows us to modify the original variable. &x in main function and xPtr in zero function refer to the same memory location.
	fmt.Println(x) // x is 7
}
