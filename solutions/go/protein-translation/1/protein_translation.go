package proteintranslation

import "errors"

// Define standard errors expected by the tests
var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

// FromCodon translates a single codon to an amino acid
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA translates an RNA strand into a sequence of amino acids
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	
	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			return proteins, ErrInvalidBase
		}
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err == ErrStop {
			return proteins, nil
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}