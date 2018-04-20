package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	actionLeader = "ACTION:"
	expectLeader = "EXPECT:"
	groupLeader  = "GROUP:"
	depsLeader   = "DEPS:"
)

var (
	breakRe  = regexp.MustCompile("^-$")
	extRe    = regexp.MustCompile("\\.qa$")
	actionRe = regexp.MustCompile("^ACTION:")
	expectRe = regexp.MustCompile("^EXPECT:")
	groupRe  = regexp.MustCompile("^GROUP:")
	depsRe   = regexp.MustCompile("^DEPS:")
)

// Token represents a single action w/ expection. This is a building block for
// the complete sheet.
type Token struct {
	sectionType string
	data        bytes.Buffer
}

// type sheet struct {
// 	groups []group
// }
//
// type group struct {
// 	name string
// 	rows []row
// }

// type Row struct {
// 	action string
// 	expect string
// }

func cleanString(str string, lead string) string {
	return strings.Trim(strings.Replace(str, lead, "", 1), " ")
}

// Used to determine if two tokens are the exact same.
func tokensMatch(t1 *Token, t2 *Token) bool {
	return t1.sectionType == t2.sectionType &&
		t1.data.String() == t2.data.String()
}

// Tokenize will take in a file path and turn it into tokens.
func Tokenize(file string) []*Token {
	inFile, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	var tokens []*Token
	var section string

	tokenRow := &Token{}

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ln := scanner.Text()

		if groupRe.MatchString(ln) {
			tokenRow.section = groupLeader
		} else if actionRe.MatchString(ln) {
			tokenRow.section = actionLeader
		} else if expectRe.MatchString(ln) {
			tokenRow.section = expectLeader
		} else if depsRe.MatchString(ln) {
			tokenRow.section = depsLeader
		} else if breakRe.MatchString(ln) {
			tokens = append(tokens, tokenRow)
			tokenRow = &Token{}
			continue
		}

		str := cleanString(ln, section)

		if len(str) > 0 {
			switch section {
			case groupLeader:
				tokenRow.group.WriteString(str)
				tokenRow.group.WriteString(" ")
			case actionLeader:
				tokenRow.action.WriteString(str)
				tokenRow.action.WriteString(" ")
			case expectLeader:
				tokenRow.expect.WriteString(str)
				tokenRow.expect.WriteString(" ")
			case depsLeader:
				tokenRow.deps.WriteString(str)
				tokenRow.deps.WriteString(" ")
			}
		}
	}

	if len(tokens) == 0 || tokensMatch(tokens[len(tokens)-1], tokenRow) == false {
		tokens = append(tokens, tokenRow)
	}

	return tokens
}

// List all .qa files in a directory.
func List(dir string) []string {
	fileList := make([]string, 0)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if extRe.MatchString(path) {
			fileList = append(fileList, path)
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return fileList
}

// func GenerateCSV() {
// 	for _, file := range fileList {
// 	}
// }
