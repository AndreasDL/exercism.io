package wordcount

import (
	"strings"
	"regexp"
)


type Frequency map[string]int

func WordCount(phrase string) Frequency{
	freq := Frequency{}

	//remove uppercase
	phrase = strings.ToLower(phrase)

	//cleanup , & newlines	
	reg, _ := regexp.Compile("[^a-z0-9'\"]+")
	phrase = reg.ReplaceAllString(phrase, " ")

	//cleanup double spaces
	reg, _ = regexp.Compile("  +")
	phrase = reg.ReplaceAllString(phrase, " ")

	//cleanup quotations
	reg, _ = regexp.Compile("'([^' ]+)'")
	phrase = reg.ReplaceAllString(phrase, "$1")	

	//magic below
	for _, w := range strings.Split(phrase, " ") {
		if len(w) > 0 && w != " " {
			freq[w]++
		}
	}

	return freq
}