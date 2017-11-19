package protein

const CODONLENGTH int = 3
var codonToProtein = map[string]string{
	"AUG" : "Methionine",
	"UUU" : "Phenylalanine",
	"UUC" : "Phenylalanine",
	"UUA" : "Leucine",
	"UUG" : "Leucine",
	"UCU" :	"Serine",
	"UCC" : "Serine",
	"UCA" : "Serine",
	"UCG" : "Serine",
	"UAU" : "Tyrosine",
	"UAC" : "Tyrosine",
	"UGU" : "Cysteine",
	"UGC" : "Cysteine",
	"UGG" : "Tryptophan",
	"UAA" : "STOP",
	"UAG" : "STOP",
	"UGA" : "STOP",
}

func FromCodon(input string) string{
	return codonToProtein[input]
}


func FromRNA(input string) []string{
	
	res := []string{}

	for start, stop := 0, CODONLENGTH ; stop <= len(input) ; start, stop = start+CODONLENGTH, stop+CODONLENGTH {

		codon := input[start:stop]
		protein := FromCodon(codon)

		if protein == "STOP"{ return res }

		res = append(res, protein)
	}

	return res
}