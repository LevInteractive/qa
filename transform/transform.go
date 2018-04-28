package transform

import (
	"strings"

	"github.com/LevInteractive/qa/scanner"
)

// Docmap is a map of documents grouped by group name and properly ordered
// by the priority number.
type Docmap map[string][]*scanner.Document

// CreateDocmap create fresh Docmap which will be used by transformers. It
// depends on a collection of documents returned from the scanner. These can be
// raw and unsorted.
func CreateDocmap(docs scanner.Documents, docmap Docmap) {
	unsatisfied := make([]*scanner.Document, 0)

	for _, doc := range docs {

		// Check if the doc is read to be apart of the map based on the dependencies
		// of the current doc.
		canAppend := satisfiesDocDeps(doc, docmap)

		// If doc does not satisfy, save it for the next loop. If it does, add it to
		// where it should go on the map's loop.
		if canAppend == false {
			unsatisfied = append(unsatisfied, doc)
		} else {
			s := doc.Name.String()
			docmap[s] = append(docmap[s], doc)
		}
	}

	if len(unsatisfied) > 0 {
		CreateDocmap(unsatisfied, docmap)
	}
}

// Check to see if the doc passed can fit somewhere on the docmap
func satisfiesDocDeps(doc *scanner.Document, docmap Docmap) bool {
	deps := strings.Split(strings.ToLower(doc.Deps.String()), ",")
	satisfied := false

	for _, dep := range deps {
		if _, ok := docmap[strings.TrimSpace(dep)]; ok == false {
			satisfied = true
			break
		}
	}

	return satisfied
}
