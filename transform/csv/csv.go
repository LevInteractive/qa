package csv

import (
	"fmt"

	"github.com/LevInteractive/qa/transform"
)

// Gen converts Documents to CSV.
func Gen(docmap transform.Docmap) {
	for k, v := range docmap {
		fmt.Printf("Key: %v\n", k)
		fmt.Printf("Value: %v\n--\n", v)
	}
}
