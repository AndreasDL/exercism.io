package anagram

import (
	"sort"
	"strings"
)



func Detect(word string, candidates []string) []string {


	//sort all letters of word => use as lookup key
	w := []rune(strings.ToLower(word))
	sort.Slice(w, func(i,j int) bool{ return w[i] < w[j] })


	//go over candidates and compared 'sorted words'
	res := []string{}
	for _, cw := range candidates {

		c := []rune(strings.ToLower(cw))
		sort.Slice(c, func(i,j int) bool { return c[i] < c[j] }) //sort chars

		if string(c) == string(w) && strings.ToLower(word) != strings.ToLower(cw) { //word itself is not an anagram !
			res = append(res, cw)
		}

	}

	return res
}