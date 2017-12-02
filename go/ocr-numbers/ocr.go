package ocr

import (
	"strings"
)

var CHAR_WIDTH  int = 3 //width  + 1 for easy indexing
var CHAR_HEIGHT int = 4

//O(1) lookups ! :D
var inputTostring = map[string]string{
	` _ | ||_|   `: "0",
	`     |  |   `: "1",
	` _  _||_    `: "2",
	` _  _| _|   `: "3",
	`   |_|  |   `: "4",
	` _ |_  _|   `: "5",
	` _ |_ |_|   `: "6",
	` _   |  |   `: "7",
	` _ |_||_|   `: "8",
	` _ |_| _|   `: "9",
}

func Recognize(s string) []string {

	lines := strings.Split(s, "\n")	
	res := []string{}
	for i, j := 1, 1+CHAR_HEIGHT ; j <= len(lines) ; i, j = i+CHAR_HEIGHT, j+CHAR_HEIGHT {
		line := lines[i:j]

		//cut line into pieces
		noChars := int( len(line[0])/CHAR_WIDTH )
		keys := make([]string, noChars)
		for _, l := range line{
			for k, b, e := 0, 0, CHAR_WIDTH ; e <= len(l) ; k, b, e = k+1, b+CHAR_WIDTH, e+CHAR_WIDTH {
				keys[k] += l[b:e]
			}
		}

		//process pieces
		lineResult := ""
		for _, k := range keys { 
			lineResult += recognizeDigit(k) 
		}

		//store line output
		res = append(res, lineResult)
	}


	return res
}

func recognizeDigit(s string) string{

	//flatten (if needed)
	s = strings.Replace(s, "\n", "", -1)


	res, exists := inputTostring[s]

	if !exists { res = "?" }

	return res
}

