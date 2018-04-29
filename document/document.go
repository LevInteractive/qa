package document

import (
	"bytes"
	"fmt"
)

// ActionExpect represents each test within a document.
type ActionExpect struct {
	Action bytes.Buffer
	Expect bytes.Buffer
}

// ActionExpectCollection is a list of tests.
type ActionExpectCollection []*ActionExpect

// Document is
type Document struct {
	ID       int64
	Priority bytes.Buffer
	Deps     bytes.Buffer
	Name     bytes.Buffer
	Tests    ActionExpectCollection
}

// Documents are a collection of *Documents.
type Documents []*Document

// Docmap is a map of documents grouped by group name and properly ordered
// by the priority number.
type Docmap map[string][]*Document

// Print a document for debugging purposes.
func Print(d *Document) {
	fmt.Printf("Name: %v\n", d.Name.String())
	for idx, test := range d.Tests {
		fmt.Printf("Action %d: %v\n", idx, test.Action.String())
		fmt.Printf("Expect %d: %v\n", idx, test.Expect.String())
	}
	fmt.Printf("Priority: %v\n", d.Priority.String())
	fmt.Printf("Deps: %v\n", d.Deps.String())
}
