package csv

import (
	"fmt"

	"github.com/LevInteractive/qa/document"
)

func headerRow() []string {
	return []string{"", "Expected Behavior", "Status", "Notes"}
}

func docHeaderRow(title string) []string {
	return []string{title, "", "", ""}
}

func docRow(action string, expect string) []string {
	return []string{action, expect, "", ""}
}

// Gen converts Documents to CSV.
func Gen(docmap document.Docmap) {
	records := [][]string{}
	records = append(records, headerRow())

	for k, v := range docmap {
		records = append(records, docHeaderRow(k))

		for _, doc := range v {
			for _, test := range doc.Tests {
				fmt.Printf("name within: %v\n", test.Expect.String())
			}
		}
	}
}
