package scanner

import (
	"bytes"
	"strings"

	"github.com/LevInteractive/qa/document"
)

// Config for main scanner method
type Config struct {

	// AllowLinebreaks will cause the parser to respect line breaks.
	AllowLinebreaks bool
}

type token int

const (
	// ILLEGAL represents a non-identifier.
	ILLEGAL token = iota
	// ACTION should be an action a tester should take.
	ACTION
	// EXPECT should be an expectation a tester should have.
	EXPECT
	// GROUP is a human friendly identifier for the group of tests.
	GROUP
	// DEPS represents other groups that should come before this one.
	DEPS
	// PRIORITY allows for order within a collection of groups.
	PRIORITY
)

// Add new action/expect test on the stack.
func addNewExpectTest(doc *document.Document) {
	newTest := &document.ActionExpect{}
	doc.Tests = append(doc.Tests, newTest)
}

func convertStringToToken(s string) token {
	switch s {
	case "ACTION":
		return ACTION
	case "EXPECT":
		return EXPECT
	case "GROUP":
		return GROUP
	case "DEPS":
		return DEPS
	case "PRIORITY":
		return PRIORITY
	default:
		return ILLEGAL
	}
}

// Called when the buffer needs to be consumed by a step.
func consumeBuffer(
	b *bytes.Buffer,
	doc *document.Document,
	lastToken *token,
	testIndex *int,
) {
	word := b.String()
	tok := convertStringToToken(strings.TrimSpace(word))

	if tok != ILLEGAL {
		if tok == ACTION {
			*testIndex = *testIndex + 1
			addNewExpectTest(doc)
		}

		*lastToken = tok
		b.Reset()
	} else {
		switch *lastToken {
		case GROUP:
			doc.Name.WriteString(word)
			b.Reset()
		case PRIORITY:
			doc.Priority.WriteString(word)
			b.Reset()
		case DEPS:
			doc.Deps.WriteString(word)
			b.Reset()
		case ACTION:
			doc.Tests[*testIndex].Action.WriteString(word)
			b.Reset()
		case EXPECT:
			doc.Tests[*testIndex].Expect.WriteString(word)
			b.Reset()
		}
	}
}

// Scan the text
func Scan(content string, c Config) *document.Document {
	// The document that will be populated.
	doc := &document.Document{}

	// Store the last identifier/token we care about so we know where to append
	// the non-ident chars.
	var lastToken token

	// Current set of runes that are being looped.
	var currentBuffer bytes.Buffer

	testIndex, l, eof, prevRune := -1, len(content), false, rune(0)

	for idx, rune := range content {
		eof = idx == l-1
		if rune == '\n' || rune == '\r' || rune == ' ' { // Consume on whitespace.

			// Add a white space if:
			// 1. If the previous rune was not a white space.
			// 2. If the current rune is a white space.
			// 3. We aren't at the EOF.
			if prevRune != ' ' && rune == ' ' {
				currentBuffer.WriteString(" ")
			} else if c.AllowLinebreaks == true && rune == '\n' || rune == '\r' {
				currentBuffer.WriteString("\n")
			}

			consumeBuffer(
				&currentBuffer,
				doc,
				&lastToken,
				&testIndex,
			)
		} else if eof == true { // Write and consume on EOF.
			currentBuffer.WriteString(string(rune))
			consumeBuffer(&currentBuffer, doc, &lastToken, &testIndex)
		} else {
			currentBuffer.WriteString(string(rune))
		}
		prevRune = rune
	}

	return doc
}
