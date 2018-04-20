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

func TestTokenize(t *testing.T) {
	sampleA := Tokenize(getFile("sample-a.qa"))
	for _, t := range sampleA {
		fmt.Println("group")
		fmt.Println(t.group.String())
		fmt.Println("expect")
		fmt.Println(t.expect.String())
		fmt.Println("action")
		fmt.Println(t.action.String())
		fmt.Println("deps")
		fmt.Println(t.deps.String())
	}
	fmt.Println(sampleA)
	if len(sampleA) != 5 {
		t.Fatalf("expected sample-a to have 5 tokens. found %v", len(sampleA))
	}
}
