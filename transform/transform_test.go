package transform

import (
	"bytes"
	"testing"

	"github.com/LevInteractive/qa/document"
)

func TestCreateDocmapOrder(t *testing.T) {
	food := &document.Document{
		ID:    "human food",
		Group: *bytes.NewBufferString("human food"),
	}
	bokchoy := &document.Document{
		ID:    "bok choy",
		Group: *bytes.NewBufferString("veges"),
		Deps:  *bytes.NewBufferString("human food"),
	}
	broccolini := &document.Document{
		ID:    "broccolini",
		Group: *bytes.NewBufferString("veges"),
		Deps:  *bytes.NewBufferString("human food"),
	}
	apple := &document.Document{
		ID:       "apple",
		Group:    *bytes.NewBufferString("fruits"),
		Priority: *bytes.NewBufferString("2"),
		Deps:     *bytes.NewBufferString("human food"),
	}
	cherry := &document.Document{
		ID:       "cherry",
		Group:    *bytes.NewBufferString("fruits"),
		Priority: *bytes.NewBufferString("1"),
		Deps:     *bytes.NewBufferString("human food"),
	}
	mango := &document.Document{
		ID:       "mango",
		Group:    *bytes.NewBufferString("fruits"),
		Priority: *bytes.NewBufferString("5"),
		Deps:     *bytes.NewBufferString("human food"),
	}
	screwdriver := &document.Document{
		ID: "screwdriver",
	}

	docs := document.Documents{
		mango,
		bokchoy,
		broccolini,
		apple,
		screwdriver,
		cherry,
		food,
	}

	list := Make(docs)

	if len(list) != 4 {
		t.Fatalf("Expected list to have length of 4. Got %v", len(list))
	}
	if list[0].Name != "" {
		t.Fatalf("Expected first doc to be ''. Got %v", list[0].Name)
	}
	if len(list[0].Docs) != 1 {
		t.Fatalf("Expected first group to have lenth of 1. Got %v", len(list[0].Docs))
	}
	if list[0].Docs[0].ID != "screwdriver" {
		t.Fatalf("Expected first doc to be screwdriver. Got %v", list[0].Docs[0].ID)
	}
	if list[1].Name != "human food" {
		t.Fatalf("Expected seconds doc to be 'human food'. Got %v", list[1].Name)
	}
	if len(list[1].Docs) != 1 {
		t.Fatalf("Expected human food to have lenth of 1. Got %v", len(list[1].Docs))
	}
	if list[2].Name != "fruits" {
		t.Fatalf("Expected group 'fruits'. Got %v", list[2].Name)
	}
	if len(list[2].Docs) != 3 {
		t.Fatalf("There should be 3 fruits. Got %v", len(list[2].Docs))
	}
	if list[3].Name != "veges" {
		t.Fatalf("Expected group 'veges'. Got %v", list[3].Name)
	}
	if len(list[3].Docs) != 2 {
		t.Fatalf("There should be 2 fruits. Got %v", len(list[3].Docs))
	}
}
