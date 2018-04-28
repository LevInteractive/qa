package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/LevInteractive/qa/scanner"
	"github.com/LevInteractive/qa/transform"
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
	var dir string

	if len(os.Args) != 2 {
		dir = "./"
	} else {
		dir = os.Args[1]
	}

	files := List(dir)

	documents := make([]*scanner.Document, 0)

	for _, file := range files {
		dat, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		documents = append(documents, scanner.Scan(string(dat)))
	}

	transform.CSV(documents)
}
