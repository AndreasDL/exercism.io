package say

import (
	"strconv"
	"strings"
//	"fmt"
)


var exceptions = map[string]string{
	"0" : "zero",
	"11": "eleven",
	"12": "twelve",
	"13": "thirdteen",
	"14": "fourteen",
	"15": "fifteen",
	"16": "sixteen",
	"17": "seventeen",
	"18": "eighteen",
	"19": "nineteen",
}


func units(digit byte) string{
	return map[rune]string{
		'1' : "one"  ,
		'2' : "two"  ,
		'3' : "three",
		'4' : "four" ,
		'5' : "five" ,
		'6' : "six"  ,
		'7' : "seven",
		'8' : "eight",
		'9' : "nine" ,
	}[rune(digit)]
}

func dozens(digit byte) string{
	switch rune(digit) {
		case '0' : return ""
		case '1' : return "teen"
		case '2' : return "twenty"
		case '3' : return "thirty"
		case '4' : return "forty"
		case '5' : return "fifty"
		case '8' : return "eighty"
		default : return units(digit) + "ty"
	}
}

func hundreds(digit byte) string{
	if rune(digit) == '0' { return "" }
	return units(digit) + " hundred"
}

func unitTypes(position int) string{
	if position > 4 { return "" }
	return []string{
		"", 
		"thousand", 
		"million", 
		"billion", 
		"trillion",
	}[position]
}

func handlePiece(nbr string) string{
	l := len(nbr)
	
	res := ""
	if l >= 3 { res += hundreds(nbr[l - 3]) + " "}

	if exception, exists := exceptions[nbr] ; exists { 
		res += exception
	} else {
		if l >= 2 { 
			res += dozens(nbr[l - 2]) 

			if nbr[l-1] != byte('0') && nbr[l-2] != byte('0') {
				res += "-"
			}
		}
		
		res += units(nbr[l-1])
	}

	return strings.Trim(res, " ")
}

func cutInPieces(nbr string) []string{
	res := []string{}

	start , stop := 0, len(nbr) % 3
	if stop > 0 {
		res = append(res, nbr[start:stop])
	}

	for start, stop = stop, stop + 3 ; stop <= len(nbr) ; start, stop = stop, stop + 3 {
		res = append(res, nbr[start:stop])
	}

	return res
}

func Say(number int64) (string, bool){
	if number < 0 { return "Out of range", false 
	} else if number > 999999999999 { return "Out of range", false }
	
	nbr := strconv.Itoa(int(number)) //usage of int64 is useless, since we don't allow numbers > 999 999 999 999

	pieces := cutInPieces(nbr)
	//fmt.Println(pieces)

	res := ""
	for i, p := range pieces {
		
		pText := handlePiece(p)
		uText := unitTypes(len(pieces) - 1 - i)

		//fmt.Println("\t", i, " ", p, " => *", pText, "* *", uText, "*")

		res += pText

		if pText != "" && uText != "" {
			res += " " + uText + " "
		}
	}

	return strings.Trim(res," ") , true	
}