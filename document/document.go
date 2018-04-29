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

// Document represents a file of tests. They are the heart and soul.
type Document struct {
	// ID is only for internal use (testing). Doesn't actually do anything be help
	// identify documents.
	ID       string
	Priority bytes.Buffer
	Deps     bytes.Buffer
	Group    bytes.Buffer
	Tests    ActionExpectCollection
}

// Documents are a collection of *Documents.
type Documents []*Document

// DocGroup represents a collection of ordered documents.
type DocGroup struct {
	Name string
	Docs Documents
}

// OrderedGroups are a sorted collection of groups.
type OrderedGroups []DocGroup

// Print a document for debugging purposes.
func Print(d *Document) {
	fmt.Printf("Group: %v\n", d.Group.String())
	for idx, test := range d.Tests {
		fmt.Printf("Action %d: %v\n", idx, test.Action.String())
		fmt.Printf("Expect %d: %v\n", idx, test.Expect.String())
	}
	fmt.Printf("Priority: %v\n", d.Priority.String())
	fmt.Printf("Deps: %v\n", d.Deps.String())
}
