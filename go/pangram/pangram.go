package pangram

import (
	"strings"
	"regexp"
)

func IsPangram (s string) bool{
	s = strings.ToLower(s)

    reg, _ := regexp.Compile("[^a-z]+")
    s = reg.ReplaceAllString(s, "")


    ctr := 0
    charsSeen := map[rune]bool{}
	for _, c := range s {

		_, exists := charsSeen[c]
		charsSeen[c] = true
		
		if !exists {
			ctr++
			if ctr == 26 {
				return true
			}
		}
	}

	return false
}