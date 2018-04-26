package scanner

import (
	"bytes"
)

const (
	// ACTION : the action
	ACTION = "ACTION"

	// EXPECT : the expect
	EXPECT = "EXPECT"

	// GROUP : the group
	GROUP = "GROUP"

	// DEPS : the group
	DEPS = "DEPS"

	// PRIORITY : is something
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
	Name     bytes.Buffer
	Deps     bytes.Buffer
	Priority int
	Tests    ActionExpectCollection
}

func write(b *bytes.Buffer, s string) {
	b.Write([]byte(s + " "))
}

// Add new action/expect test on the stack.
func addNewExpectTest(group *Document) {
	newTest := &ActionExpect{}
	group.Tests = append(group.Tests, newTest)
}

func consumeBuffer(b *bytes.Buffer, group *Document, lastToken *string, isBr bool) {
	word := b.String()
	switch word { // Check to see if word is identifier.
	case GROUP:
		*lastToken = GROUP
		b.Reset()
	case ACTION:
		*lastToken = ACTION
		addNewExpectTest(group)
		b.Reset()
	case EXPECT:
		*lastToken = EXPECT
		b.Reset()
	case DEPS:
		*lastToken = DEPS
		b.Reset()
	case PRIORITY:
		*lastToken = PRIORITY
		b.Reset()
	default: // The full word isn't an indentifier.
		// If the last rune was a line break, preserve it.
		if isBr {
			word = word + "\n"
		}

		switch *lastToken {
		case GROUP:
			write(&group.Name, word)
			b.Reset()
		case ACTION:
			write(&group.Tests[len(group.Tests)-1].Action, word)
			b.Reset()
		case EXPECT:
			write(&group.Tests[len(group.Tests)-1].Expect, word)
			b.Reset()
		case DEPS:
			write(&group.Deps, word)
			b.Reset()
		}
	}
}

// Scan the text
func Scan(content string) *Document {
	// The group that will be populated.
	group := &Document{}

	// Store the last identifier/token we care about so we know where to append
	// the non-ident chars.
	var lastToken string

	// Current set of runes that are being looped.
	var currentBuffer bytes.Buffer

	for _, rune := range content {
		if rune == '\n' || rune == '\r' {
			consumeBuffer(&currentBuffer, group, &lastToken, true)
		} else if rune == ' ' {
			consumeBuffer(&currentBuffer, group, &lastToken, false)
		} else {
			currentBuffer.Write([]byte(string(rune)))
		}
	}

	return group
}
