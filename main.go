package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/LevInteractive/qa/document"
	"github.com/LevInteractive/qa/scanner"
	"github.com/LevInteractive/qa/transform"
	"github.com/LevInteractive/qa/transform/csv"
)

var helptxt = `
Qa

1. Navigate/cd into the root of your project with the .qa files.
2. Run: qa . > qa.csv
3. Done!
`

func main() {
	helpMode := flag.Bool("h", false, "help")
	flag.Parse()

	if *helpMode == true {
		fmt.Println(helptxt)
		return
	}

	var dir string

	if len(flag.Args()) == 0 {
		fmt.Println(helptxt)
		return
	}

	dir = flag.Args()[0]

	files := List(dir)

	config := scanner.Config{
		AllowLinebreaks: false,
	}

	documents := make(document.Documents, 0)

	for _, file := range files {
		dat, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		documents = append(documents, scanner.Scan(string(dat), config))
	}

	groups := transform.Make(documents)

	csv.Gen(groups)
}

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
