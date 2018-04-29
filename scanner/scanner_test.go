package scanner

import (
	"io/ioutil"
	"testing"
)

func readFile(f string) string {
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func TestGeneralFormatting(t *testing.T) {
	var tests = []struct {
		in  string
		out string
		c   Config
	}{
		{`GROUP This is a test`, "This is a test", Config{}},
		{"GROUP W1#@rd characters!", "W1#@rd characters!", Config{}},
		{"GROUP My linebreak is preserved!\n", "My linebreak is preserved!\n", Config{AllowLinebreaks: true}},
		{"GROUP My linebreak is NOT preserved!\n", "My linebreak is NOT preserved!", Config{}},
		{`GROUP    w h i t e s p a c e   `, "w h i t e s p a c e ", Config{}},
	}

	for _, test := range tests {
		doc := Scan(test.in, test.c)

		if doc.Group.String() != test.out {
			t.Fatalf("Expected \"%v\", but got \"%v\"", test.out, doc.Group.String())
		}
	}
}

func TestScanner(t *testing.T) {
	a := Scan(readFile("../test/fixtures/sample-a.qa"), Config{})
	var v string

	v = a.Group.String()
	if v != "User Authentication" {
		t.Fatalf("Expected name 'User Authentication', but got '%v'", v)
	}

	v = a.Tests[0].Action.String()
	if v != "Apples, Oranges" {
		t.Fatalf("Expected name 'Apples, Oranges', but got '%v'", v)
	}

	// b := Scan(readFile("../test/fixtures/sample-b.qa"), Config{})

	// document.Print(a)
	// document.Print(b)
}
