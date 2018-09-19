package main

import "os"
// create a file in the working directory 
func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("Lorem ipsum dolor sit amet, modo cetero dissentiet eos ei.")
}
