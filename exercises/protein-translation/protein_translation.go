package protein

import "errors"

// ErrStop indicates one of the stop codons.
var ErrStop = errors.New("stop")

// ErrInvalidBase indicates an unrecognized codon.
var ErrInvalidBase = errors.New("invalid base")

// FromRNA translate RNA sequences into proteins.
func FromRNA(rna string) ([]string, error) {
	var proteins []string = []string{}

	lenRNA := len(rna)
	if lenRNA == 0 || lenRNA%3 != 0 {
		return proteins, ErrInvalidBase
	}

	for i := 3; i <= lenRNA; i += 3 {
		codon, err := FromCodon(rna[i-3 : i])
		if err == ErrStop {
			return proteins, nil
		}
		if err != nil {
			return []string{}, err
		}
		proteins = append(proteins, codon)
	}

	return proteins, nil
}

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
		return "STOP", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
