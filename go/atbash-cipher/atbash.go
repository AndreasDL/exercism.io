package atbash

const CIPHER_GROUP_LENGTH = 5

var mapper = map[rune]string{ //lazy but convienent
	'a' : "z",
	'b' : "y",
	'c' : "x",
	'd' : "w",
	'e' : "v",
	'f' : "u",
	'g' : "t",
	'h' : "s",
	'i' : "r",
	'j' : "q",
	'k' : "p",
	'l' : "o",
	'm' : "n",
	'n' : "m",
	'o' : "l",
	'p' : "k",
	'q' : "j",
	'r' : "i",
	's' : "h",
	't' : "g",
	'u' : "f",
	'v' : "e",
	'w' : "d",
	'x' : "c",
	'y' : "b",
	'z' : "a",
	
	'A' : "z",
	'B' : "y",
	'C' : "x",
	'D' : "w",
	'E' : "v",
	'F' : "u",
	'G' : "t",
	'H' : "s",
	'I' : "r",
	'J' : "q",
	'K' : "p",
	'L' : "o",
	'M' : "n",
	'N' : "m",
	'O' : "l",
	'P' : "k",
	'Q' : "j",
	'R' : "i",
	'S' : "h",
	'T' : "g",
	'U' : "f",
	'W' : "e",
	'V' : "d",
	'X' : "c",
	'Y' : "b",
	'Z' : "a",

	'1' : "1",
	'2' : "2",
	'3' : "3",
	'4' : "4",
	'5' : "5",
	'6' : "6",
	'7' : "7",
	'8' : "8",
	'9' : "9",
	'0' : "0",
}

func Atbash(plain string) string{

	cipher := ""

	i := 0
	for _, c := range plain {

		if i > 0 && i % CIPHER_GROUP_LENGTH == 0 { 
			cipher += " "
			i = 0
		}

		
		if char, exists := mapper[c] ; exists {
			cipher += char 
			i++
		}
	}


	if i == 0 { //remove trailing " "
		cipher = cipher[:len(cipher)-1]
	}

	return cipher
}