package main

import (
	"fmt"
	"testing"
)

const fixturePath = "test/fixtures"

func getFile(s string) string {
	return fmt.Sprintf("%s/%s", fixturePath, s)
}

func TestList(t *testing.T) {
	files := List(fixturePath)

	if len(files) != 3 {
		t.Fatalf("bad number of returned qa files: %v", len(files))
	}
	if files[0] != getFile("sample-a.qa") {
		t.Fatalf("wrong file name: %v - should be sample-a.qa", files[0])
	}
	if files[1] != getFile("sample-b.qa") {
		t.Fatalf("wrong file name: %v - should be sample-b.qa", files[1])
	}
	if files[2] != getFile("sample-c.qa") {
		t.Fatalf("wrong file name: %v - should be sample-c.qa", files[2])
	}
}
