//program that converts feet into meters

package main

import "fmt"

func main() {
	fmt.Println("Enter a number to convert feet into meters")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println(output)
}
