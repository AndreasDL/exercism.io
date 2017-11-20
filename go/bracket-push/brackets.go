package brackets


var openers = map[rune]rune{
	']': '[',
	')': '(',
	'}': '{',
}

type stack []rune

func (s *stack) pop() rune{
	l := len(*s)
	if l == 0 {
		return 'a'
	} else { 
		c := (*s)[l-1]
		*s = (*s)[:l-1] 
		return c
	}
}

func (s *stack) push(c rune) {
	*s = append(*s, c)
}


func Bracket(s string) (bool, error) {
	stck := stack{}

	for i := 0 ; i < len(s) ; i++ {

		c := rune(s[i])
	
		
		if c == '[' || c == '(' || c == '{' {  //opening		
			stck.push(c)
		} else if c == ']' || c == ')' || c == '}' { //closing

			shouldbe := openers[c]
			actual := stck.pop()

			if shouldbe != actual {
				return false, nil
			}
		}


	}

	return len(stck) == 0 , nil
}