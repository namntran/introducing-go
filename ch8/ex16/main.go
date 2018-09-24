package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Walk function in the path/filepath package: recursively walk a folder (read the folder's contents, all the subfolders, all the sub-folders)
func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error { //the function passed to Walk is called for every file and folder in the root folder
		//Walk function iw passed 3 arguments: path which is the path to the ffile; info, which is the information for the file; and err, which is any error received while walking the directory
		fmt.Println(path)
		return nil
	})
}
