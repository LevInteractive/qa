package csv

import (
	"fmt"

	"github.com/LevInteractive/qa/scanner"
	"github.com/LevInteractive/qa/transform"
)

// CSV converts Documents to CSV.
func CSV(documents scanner.Documents) {
	docmap := make(transform.DocmapType)
	transform.Gen(documents, docmap)

	for k, v := range docmap {
		fmt.Printf("Key: %v\n", k)
		fmt.Printf("Value: %v\n--\n", v)
	}
}
