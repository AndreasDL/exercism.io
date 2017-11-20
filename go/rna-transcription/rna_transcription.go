package strand

var dnaToRna = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

//Unsafe !
func ToRNA(dna string) string {
	res := ""

	for _, c := range dna {

		res += dnaToRna[c]
	}

	return res
}