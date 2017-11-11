package isogram

import (
	"strings"
	"regexp"
)


func IsIsogram(s string) bool{

	s = strings.ToUpper(s)
	reg, _ := regexp.Compile("[^a-zA-Z]")
	s = reg.ReplaceAllString(s, "")

	freq := map[rune]int{}

	for _, c := range s {
		freq[c]++

		if freq[c] == 2 {
			return false
		}
	}

	return true
}