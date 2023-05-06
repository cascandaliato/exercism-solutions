package protein

import (
	"errors"
)

const (
	methionine    = "Methionine"
	phenylalanine = "Phenylalanine"
	leucine       = "Leucine"
	serine        = "Serine"
	tyrosine      = "Tyrosine"
	cysteine      = "Cysteine"
	tryptophan    = "Tryptophan"
	stop          = "STOP"
)

var codonToProtein = map[string]string{
	"AUG": methionine,
	"UUU": phenylalanine,
	"UUC": phenylalanine,
	"UUA": leucine,
	"UUG": leucine,
	"UCU": serine,
	"UCC": serine,
	"UCA": serine,
	"UCG": serine,
	"UAU": tyrosine,
	"UAC": tyrosine,
	"UGU": cysteine,
	"UGC": cysteine,
	"UGG": tryptophan,
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

// ErrStop happens when trying to convert a 'STOP' codon into a protein.
var ErrStop = errors.New("cannot convert a 'STOP' codon into a protein")

// ErrInvalidBase happens when an invalid codon value is specified.
var ErrInvalidBase = errors.New("cannot convert an invalid codon into a protein")

// FromCodon converts a codon into the corresponding protein.
func FromCodon(codon string) (string, error) {
	protein, ok := codonToProtein[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == stop {
		return "", ErrStop
	}
	return protein, nil
}

// FromRNA converts an RNA sequence into a sequence of proteins.
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				break
			}
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
