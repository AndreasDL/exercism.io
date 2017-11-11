package bob

import "strings"
import "unicode"

//Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//Bob answers 'Sure.' if you ask him a question.
//He answers 'Whoa, chill out!' if you yell at him.
//He says 'Fine. Be that way!' if you address him without actually saying
//anything.
//He answers 'Whatever.' to anything else.

func containsLetters( s string) bool{
	for _, letter := range s{
		if unicode.IsLetter(letter){
			return true
		}
	}
	return false
}

func Hey(remark string) string {
	response_question := "Sure."
	response_yell := "Whoa, chill out!"
	response_nothing := "Fine. Be that way!"
	response_default := "Whatever."


	trimmed_remark := strings.Trim(remark, " \t\n\r")


	if strings.ToUpper(trimmed_remark) == trimmed_remark && containsLetters(trimmed_remark){
		return response_yell
	} else if strings.HasSuffix(trimmed_remark, "?"){
		return response_question
	} else if trimmed_remark == "" {
		return response_nothing
	} else {
		return response_default
	}
}