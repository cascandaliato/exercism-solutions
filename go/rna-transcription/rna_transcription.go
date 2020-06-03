// Package strand provides utility functions to manipulate DNA and RNA strands.
package strand

var dnaToRna = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA converts a DNA strand into a RNA strand.
func ToRNA(dna string) string {
	rna := ""
	for _, v := range dna {
		rna += dnaToRna[string(v)]
	}
	return rna
}
