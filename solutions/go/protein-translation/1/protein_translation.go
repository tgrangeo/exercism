package protein

import "errors"

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine",
	"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine",
	"UGU": "Cysteine", "UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	res := []string{}
	for i := 0; i+3 <= len(rna); i += 3 {
		codon := rna[i : i+3]
		protein, err := FromCodon(codon)
		if err != nil {
			return nil, err
		}
		if protein == "STOP" {
			break
		}
		res = append(res, protein)
	}
	return res, nil
}

func FromCodon(codon string) (string, error) {
	protein, ok := codonMap[codon]
	if !ok {
		return "", errors.New("invalid codon")
	}
	return protein, nil
}
