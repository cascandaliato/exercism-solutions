package protein

import (
	"errors"
)

type protein string

const (
	methionine    protein = "Methionine"
	phenylalanine protein = "Phenylalanine"
	leucine       protein = "Leucine"
	serine        protein = "Serine"
	tyrosine      protein = "Tyrosine"
	cysteine      protein = "Cysteine"
	tryptophan    protein = "Tryptophan"
	stop          protein = "STOP"
)

var codonToProtein = map[string]protein{
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

// ErrStop happens when trying to convert a 'STOP' codon to protein.
var ErrStop = errors.New("cannot convert a 'STOP' codon into a protein")

// ErrInvalidBase happens when an invalid codon value is specified.
var ErrInvalidBase = errors.New("cannot convert an invalid codon into a protein")

// FromCodon converts a codon into the corresponding protein.
func FromCodon(codon string) (string, error) {
	p, ok := codonToProtein[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if p == stop {
		return "", ErrStop
	}
	return string(p), nil
}

// FromRNA converts an RNA sequence into a sequence of proteins.
func FromRNA(rna string) ([]string, error) {
	var proteins []string

	var codon string
	var i int
	for _, ch := range rna {
		codon += string(ch)

		if (i+1)%3 == 0 {
			protein, err := FromCodon(codon)
			if err == ErrInvalidBase {
				return proteins, err
			}
			if err == ErrStop {
				break
			}
			proteins = append(proteins, protein)
			codon = ""
		}

		i++
	}
	return proteins, nil
}
