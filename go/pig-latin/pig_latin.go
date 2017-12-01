package igpay

import "strings"


func isVowel(c byte) bool{
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'x' || c== 'y'
}

func PigLatin(input string) string{
	//No i'm not writings or / else if statements !
	res := ""
	for _, s := range strings.Split(input, " ") {

		if s[:3] == "squ" || s[:3] == "thr" || s[:3] == "sch" {
			s = s[3:] + s[:3]
		
		} else if s[:2] == "ch" || s[:2] == "qu"	|| s[:2] == "th"{
			s = s[2:] + s[:2]
		
		} else if ! isVowel(s[0]) {
			s = s[1:] + string(s[0])
		
		} else if (s[0] == 'y' || s[0] == 'x') && isVowel(s[1]) {
			s = s[1:] + string(s[0])

		}

		res += s + "ay "
	}

	 

	//always add "ay"
	return res[:len(res)-1]
}