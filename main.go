package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// List all .qa files in a directory.
func List(dir string) []string {
	extRe := regexp.MustCompile("\\.qa$")
	fileList := make([]string, 0)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if extRe.MatchString(path) {
			fileList = append(fileList, path)
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return fileList
}

func main() {
	files := List(os.Args[1])
	fmt.Println(files)
}
