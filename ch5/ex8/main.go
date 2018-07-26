package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3} // slices are typically used to represent lists of items, particularly when you need to
	// access the nth item quickly - e.g. player #33, or the the 10th most popular search query.
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

// after running the program slice1 has [1 2 3] and slice2 has [1 2]. the contenct of slice1 and copied into slice2,
// but because slice2 has room for only two elements, only the first two elements of slice1 are copied.
