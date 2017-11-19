package cryptosquare

import (
	"strings"
	"regexp"
	"math"
)

func Encode(s string) string{

	plain := normalize(s)
	cipher := encode(plain)

	return cipher
}


func normalize(s string) string{
	/*
		If man was meant to stay on the ground, god would have given us roots.
	is normalized to:
		ifmanwasmeanttostayonthegroundgodwouldhavegivenusroots
	*/

	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-z1-9]+")
    s = reg.ReplaceAllString(s, "")

    return s
}

func encode(s string) string{

	c:= int(math.Ceil(math.Sqrt(float64(len(s)))))
	lines := make([]string, c)
	for i := 0 ; i < c ; i++{

		res := ""
		for j := i; j < len(s) ; j += c {
			res += string(s[j])
		}

		lines[i] = res
	}

	return strings.Join(lines, " ")
}