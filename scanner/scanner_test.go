package scanner

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func printDoc(d *Document) {
	fmt.Printf("Name: %v\n", d.Name.String())
	fmt.Printf("Action 1: %v\n", d.Tests[0].Action.String())
	fmt.Printf("Expect 1: %v\n", d.Tests[0].Expect.String())
	fmt.Printf("Action 2: %v\n", d.Tests[1].Action.String())
	fmt.Printf("Expect 2: %v\n", d.Tests[1].Expect.String())
	fmt.Printf("Priority: %v\n", d.Priority.String())
	fmt.Printf("Deps: %v\n", d.Deps.String())
}

func readFile(f string) string {
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func TestScanner(t *testing.T) {
	a := Scan(readFile("../test/fixtures/sample-a.qa"))
	b := Scan(readFile("../test/fixtures/sample-b.qa"))

	// @TODO write tests.
	printDoc(a)
	printDoc(b)
}
