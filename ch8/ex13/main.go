package main

import (
	"fmt"
	"io/ioutil"
)

// read the contents of a file and display them on a terminal
func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
