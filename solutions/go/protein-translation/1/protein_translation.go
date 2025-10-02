package protein

import (
    "fmt"
)

var ErrInvalidBase = fmt.Errorf("no more tea available")
var ErrStop = fmt.Errorf("stop")

func FromRNA(rna string) ([]string, error) {
    var proteins []string
	for i :=0; i < len(rna); i += 3 {
        protein, err := FromCodon(rna[i:i+3])
        if err == ErrStop {
            return proteins, nil
        }
        if err == ErrInvalidBase {
            return proteins, ErrInvalidBase
        }
        proteins = append(proteins, protein)
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
        	return "", ErrStop
        default:
        	return "", ErrInvalidBase
    }
}
