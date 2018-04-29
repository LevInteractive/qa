package transform

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/LevInteractive/qa/document"
)

// Make create fresh Docmap which will be used by transformers. It
// depends on a collection of documents returned from the scanner. These can be
// raw and unsorted.
func Make(docs document.Documents) document.OrderedGroups {
	groups := document.OrderedGroups{}
	build(docs, &groups, 0)
	return groups
}

func build(docs document.Documents, groups *document.OrderedGroups, lastCount int) {
	unsatisfied := make(document.Documents, 0)

	for _, doc := range docs {
		// Check if the doc is read to be apart of the map based on the dependencies
		// of the current doc.
		satisfied := satisfiesDocDeps(doc, groups)

		// If doc does not satisfy, save it for the next loop. If it does, add it to
		// where it should go on the map's loop.
		if satisfied == false {
			unsatisfied = append(unsatisfied, doc)
		} else {
			s := doc.Group.String()
			addedToExistingGroup := false

			for idx, group := range *groups {
				if group.Name == s {
					group.Docs = append(group.Docs, doc)
					sort.Slice(group.Docs, func(i, j int) bool {
						p1, err := strconv.Atoi(group.Docs[i].Priority.String())
						if err != nil {
							p1 = 1
						}

						p2, err := strconv.Atoi(group.Docs[j].Priority.String())
						if err != nil {
							p2 = 1
						}

						return p1 < p2
					})
					(*groups)[idx] = group
					addedToExistingGroup = true
					break
				}
			}

			// Create a new group
			if addedToExistingGroup == false {
				*groups = append(*groups, document.DocGroup{
					Name: s,
					Docs: document.Documents{doc},
				})
			}
		}
	}

	newCount := len(unsatisfied)

	// If there are unsatisified docs and has the same amount as last time, then
	// throw a fatal.
	if newCount > 0 && newCount == lastCount {
		log.Fatalf("There is a unsatisfied doc: %v", newCount)
	}

	// There are still some unsatisified docs, evaluate again.
	if newCount > 0 {
		build(unsatisfied, groups, newCount)
	}
}

// Based on the doc's dep string, return a trimed array of names.
func normalizedDeps(doc *document.Document) []string {
	deps := make([]string, 0)

	for _, part := range strings.Split(doc.Deps.String(), ",") {
		p := strings.TrimSpace(part)
		if len(p) > 0 {
			deps = append(deps, p)
		}
	}

	return deps
}

func satisfiesDocDeps(doc *document.Document, groups *document.OrderedGroups) bool {
	deps := normalizedDeps(doc)
	// fmt.Printf("Doc group: %v | deps: %v | deps size: %v\n", doc.Group.String(), deps, len(deps))

	for _, dep := range deps {

		satisfiedDep := false

		for _, group := range *groups {
			if group.Name == dep {
				satisfiedDep = true
			}
		}

		// Could not satisfy a dep, return now.
		if satisfiedDep == false {
			return false
		}
	}

	return true
}
