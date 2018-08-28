package main

import "fmt"

func makeOddGenerator() func() uint { // function that generates odd numbers
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // prints 1
	fmt.Println(nextOdd()) // prints 3
	fmt.Println(nextOdd()) // prints 5
	fmt.Println(nextOdd()) // prints 7
}

// using makeEvenGenerator program as an example

// package main

// import "fmt"

// func makeEvenGenerator() func() uint { //closure example - write a function that returns another function.
// 	i := uint(0)
// 	return func() (ret uint) {
// 		ret = i
// 		i += 2
// 		return
// 	}
// }
// func main() {
// 	nextEven := makeEvenGenerator() //(nextEven)returns a (makeEvenGenerator)function that generates even numbers.
// 	fmt.Println(nextEven()) // returns 0
// 	fmt.Println(nextEven()) // returns 2
// 	fmt.Println(nextEven()) // returns 4
// }
