package scrabble

import "strings"
/*
## Examples
"cabbage" should be scored as worth 14 points:

- 3 points for C
- 1 point for A, twice
- 3 points for B, twice
- 2 points for G
- 1 point for E

And to total:

- `3 + 2*1 + 2*3 + 2 + 1`
- = `3 + 2 + 6 + 3`
- = `5 + 9`
- = 14
*/

func getScore(c rune) int{
	
	if strings.ContainsRune("AEIOULNRST", c) {
		return 1
	} else if strings.ContainsRune("DG", c){
		return 2
	} else if strings.ContainsRune("BCMP", c){
		return 3
	} else if strings.ContainsRune("FHVWY", c){
		return 4
	} else if strings.ContainsRune("K", c){
		return 5
	} else if strings.ContainsRune("JX", c){
		return 8
	} else if strings.ContainsRune("QZ", c){
		return 10
	}
	return 0
}

func Score(s string) int {
	
	s = strings.ToUpper(s) 

	score := 0
	for _, c := range s {
		score += getScore(c)
	}

	return score
}