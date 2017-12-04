package cipher


import (
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}


func NewCaesar() Cipher{ 
	return NewShift(3) 
}

func NewShift(v int) Cipher{ 
	if v <= -26 || v >= 26 || v == 0 { return nil }

	return cipher{ []int{v} }
}

func NewVigenere(key string) Cipher{ 

	if !isValid(key) { return nil }

	k := []int{}
	for _, c := range key { 
		
		c = unicode.ToLower(c)

		k = append(k, int(c - 'a'))
	}

	return cipher{k}
}

func isValid(key string)bool {
	if len(key) == 0 { return false }

	valid := false
	for _, c := range key {

		if !unicode.IsLetter(c) { 
			return false 
		} else if !unicode.IsLower(c){
			return false			
		} else if c != 'a' { 
			valid = true 
		}

	}
	
	return valid
}


type cipher struct{
	key []int
}
func(cipher cipher) Encode(s string) string{
	return cipher.applyMap(
				s, 
				func(a, b int) rune { return rune(a + b) },
			)
}
func(cipher cipher) Decode(s string) string{
	return cipher.applyMap(
				s,
				func(a,b int) rune { return rune(a - b) },
			)
}
func (cipher cipher) applyMap(s string, f func(int,int) rune) string{
	res := ""
	
	i := 0
	for _, c := range s {

		if !unicode.IsLetter(c) { continue }
		c = unicode.ToLower(c)

		dist := cipher.key[i % len(cipher.key)]
		e := f( int(c), dist )

		if e < 'a' {
			e += ('z' - 'a' + 1)
		} else if e > 'z'{
			e -= ('z' - 'a' + 1)
		}

		res += string(e)
		i++
	}
	return res
}