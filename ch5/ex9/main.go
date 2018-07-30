package main

import "fmt"

func main() {
	// var x map[string]int //x is a map of strings to ints
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
}
	
