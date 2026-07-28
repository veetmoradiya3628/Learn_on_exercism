package nucleotidecount

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range d {
		if _, exists := h[nucleotide]; !exists {
			return nil, errors.New("invalid nucleotide")
		}
		h[nucleotide]++
	}

	return h, nil
}   