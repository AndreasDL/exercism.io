package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	if nucleotide != byte('A') && 
	   nucleotide != byte('C') &&
	   nucleotide != byte('G') &&
	   nucleotide != byte('T') {
	   	return -1, errors.New(string(nucleotide) + " Is not a valid nucleotide!")
	}


	hist, err := d.Counts()
	if err != nil { return -1, err }

	v, e := hist[nucleotide]
	if !e { v = 0 }

	return v, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	hist := Histogram{
		byte('A'): 0,
		byte('C'): 0,
		byte('G'): 0,
		byte('T'): 0,
	}

	for _, c := range d {
		
		if _, exists := hist[byte(c)] ; !exists { 
			return nil, errors.New(string(c) + " Is not a valid nucleotide!") 
		}

		hist[byte(c)]++
	}

	return hist, nil
}
