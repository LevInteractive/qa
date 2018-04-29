package csv

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/LevInteractive/qa/document"
)

func headerRow() []string {
	return []string{"Action", "Expect", "Status", "Notes"}
}

func docHeaderRow(title string) []string {
	return []string{title, "", "", ""}
}

func docRow(action string, expect string) []string {
	return []string{action, expect, "", ""}
}

// Gen converts Documents to CSV.
func Gen(groups document.OrderedGroups) {
	records := [][]string{}
	records = append(records, headerRow())

	for _, g := range groups {
		records = append(records, docHeaderRow(g.Name))

		for _, doc := range g.Docs {
			for _, test := range doc.Tests {
				records = append(
					records,
					docRow(test.Action.String(), test.Expect.String()),
				)
			}
		}
	}
	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
