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
)

// ErrStop happens when trying to convert a 'STOP' codon into an amino acid.
var ErrStop = errors.New("cannot convert a 'STOP' codon into an amino acid")

// ErrInvalidBase happens when an invalid codon value is specified.
var ErrInvalidBase = errors.New("cannot convert an invalid codon into an amino acid")

// FromCodon converts a codon into the corresponding amino acid.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "UGC", "UGU":
		return cysteine, nil
	case "UUA", "UUG":
		return leucine, nil
	case "AUG":
		return methionine, nil
	case "UUC", "UUU":
		return phenylalanine, nil
	case "UCA", "UCC", "UCG", "UCU":
		return serine, nil
	case "UGG":
		return tryptophan, nil
	case "UAC", "UAU":
		return tyrosine, nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
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
			return nil, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
