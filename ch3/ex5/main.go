package main

import "fmt"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x += " & second"
	fmt.Println(x)
	x += " & third"
	fmt.Println(x)
}
