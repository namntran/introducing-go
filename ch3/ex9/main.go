//program that converts Fahrenheit into Celsius (C=(F-32)* 5/9)

package main

import "fmt"

// func main() {
// 	fmt.Println("Enter a number")
// 	var input float64
// 	fmt.Scanf("%f", &input)

// 	output := input * 2

// 	fmt.Println(output)

func main() {
	fmt.Println("Enter a number to convert degrees Fahrenheit into degrees Celsius ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5 / 9

	fmt.Println(output)

}
