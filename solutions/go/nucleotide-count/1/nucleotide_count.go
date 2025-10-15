package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in a given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides, which is essentially a string.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	// Iterate through each nucleotide in the DNA strand
	for _, nucleotide := range d {
		switch nucleotide {
		case 'A', 'C', 'G', 'T':
			h[nucleotide]++
		default:
			return nil, errors.New("bad nucleotide found")
		}
	}
	return h, nil
}
