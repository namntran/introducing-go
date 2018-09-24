package main

import (
	"fmt"
	"os"
)
// get the contents of a directory
func main() {
	dir, err := os.Open(".") //use os.Open function to get the contents of a directory and give it a directory path instead of a filename
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // Readdir takes a single argument that limits the number of entries returned.
	// Passing -1 returns all of the entries
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
