package main

import "fmt"

func main() { //closure - a function with nonlocal variables it references.  increment and x form the "closure".
	x := 0
	increment := func() int { //increment adds 1 to the variable x defined by main function's scope. the x variable can be accessed and modiffied by the increment function.
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
