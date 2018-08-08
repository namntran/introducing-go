// variables must be passed to functions
// functions don't have access to anything in the calling function unless it's passed in explicitly
package main

import "fmt"

func f(x int) {
	fmt.Println(x)
}

func main() {
	x := 5
	f(x)
}
