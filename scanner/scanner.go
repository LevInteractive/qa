package scanner

import (
	"bytes"
	"strings"
)

const (
	ACTION   = "ACTION"
	EXPECT   = "EXPECT"
	GROUP    = "GROUP"
	DEPS     = "DEPS"
	PRIORITY = "PRIORITY"
)

// ActionExpect represents each word.
type ActionExpect struct {
	Action bytes.Buffer
	Expect bytes.Buffer
}

// ActionExpectCollection is a list
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

func write(b *bytes.Buffer, s string) {
	b.Write([]byte(strings.Trim(s, " ") + " "))
}

// Add new action/expect test on the stack.
func addNewExpectTest(document *Document) {
	newTest := &ActionExpect{}
	document.Tests = append(document.Tests, newTest)
}

func consumeBuffer(b *bytes.Buffer, document *Document, lastToken *string, isBr bool) {
	word := b.String()
	switch word { // Check to see if word is identifier.
	case GROUP, ACTION, EXPECT, DEPS, PRIORITY:
		*lastToken = word
		b.Reset()
	default: // The full word isn't an indentifier.
		// If the last rune was a line break, preserve it.
		if isBr {
			word = word + "\n"
		}

		switch *lastToken {
		case GROUP:
			write(&document.Name, word)
			b.Reset()
		case PRIORITY:
			write(&document.Priority, word)
			b.Reset()
		case DEPS:
			write(&document.Deps, word)
			b.Reset()
		case ACTION:
			addNewExpectTest(document)
			write(&document.Tests[len(document.Tests)-1].Action, word)
			b.Reset()
		case EXPECT:
			write(&document.Tests[len(document.Tests)-1].Expect, word)
			b.Reset()
		}
	}
}

// Scan the text
func Scan(content string) *Document {
	// The document that will be populated.
	document := &Document{}

	// Store the last identifier/token we care about so we know where to append
	// the non-ident chars.
	var lastToken string

	// Current set of runes that are being looped.
	var currentBuffer bytes.Buffer

	for _, rune := range content {
		if rune == '\n' || rune == '\r' {
			consumeBuffer(&currentBuffer, document, &lastToken, true)
		} else if rune == ' ' {
			consumeBuffer(&currentBuffer, document, &lastToken, false)
		} else {
			currentBuffer.Write([]byte(string(rune)))
		}
	}

	return document
}
