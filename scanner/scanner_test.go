package scanner

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestScanner(t *testing.T) {
	dat, err := ioutil.ReadFile("../test/fixtures/sample-a.qa")
	if err != nil {
		panic(err)
	}

	g := Scan(string(dat))
	fmt.Println(g.Name.String())
	fmt.Println(g.Tests[0].Action.String())
	fmt.Println(g.Tests[0].Expect.String())

	fmt.Println(g.Tests[1].Action.String())
	fmt.Println(g.Tests[1].Expect.String())
}
