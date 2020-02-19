package protein

import "errors"

var (
	// ErrStop indicates a STOP codon
	ErrStop        error = errors.New("STOP codon")
	// ErrInvalidBase indicates an invalid codon was found
	ErrInvalidBase error = errors.New("Invalid codon")
)

// FromCodon returns amino acid names, given a codon string
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

// FromRNA cuts up an RNA string into codons and builds a list of
// amino acids from them
func FromRNA(rna string) ([]string, error) {
	out := make([]string, 0, len(rna)/3)
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		}
		if err != nil {
			return out, err
		}

		out = append(out, protein)
	}

	return out, nil
}
