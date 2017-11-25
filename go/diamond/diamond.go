package diamond

import "errors"
//import "fmt"

var STARTBYTE = byte('A')

func genSpaces(spaces byte) string{
	res := ""
	for ; spaces > byte(0) ; spaces-- { 
		res += " " 
	}

	return res
}
func firstLastLine(spaces byte) string {
	res := genSpaces(spaces)
	res += string(STARTBYTE)
	res += genSpaces(spaces)
	res += "\n"

	return res
}

func otherLine(totalSpaces, spaces, currByte byte) string{
	res := genSpaces(spaces)
	res += string(currByte)
	res += genSpaces(totalSpaces - spaces - spaces)
	res += string(currByte)
	res += genSpaces(spaces)
	res += "\n"

	return res
}



func Gen(c byte) (string,error){

	if c < 'A' || c > 'Z' { return "", errors.New("Char out of range")}

	spaces := c - STARTBYTE
	totalSpaces := spaces * 2 - 1

	res := firstLastLine(spaces)
	spaces--
	
	currByte := STARTBYTE+1
	for ; currByte < c ; currByte++ {
		res += otherLine(totalSpaces, spaces, currByte)
		spaces--
	}

	if c != byte('A') {	res += otherLine(totalSpaces, spaces, currByte) }
	currByte--
	spaces++

	for ; currByte > STARTBYTE ; currByte-- {
		res += otherLine(totalSpaces, spaces, currByte)
		spaces++
	}


	if c != byte('A') {	res += firstLastLine(spaces) }
/*
	fmt.Println(res)
	fmt.Println()
*/
	return res, nil
}