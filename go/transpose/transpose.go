package transpose 

import (
	"strings"
)

func Transpose(input []string) []string{
	if len(input) == 0 { return []string{} }
	

	padLinesInPlace(&input)

	res := transpose(&input)

	return res
}

//pads wit additional spaces
//returns the width of the input string == hieght of the transposed version
func padLinesInPlace(input *[]string) {

	maxLength := 0

	//go right to left => so we pad in right direction
	for i := len( *input )-1; i >= 0 ; i-- {

		row := (*input)[i]

		if l := len(row) ; l > maxLength{
			maxLength = l
		} else {
			(*input)[i] = row + strings.Repeat(" ", maxLength - len(row))
		}
	}
}

func transpose(input *[]string) []string{
	data := *input

	//init output table
	output := make([][]byte, len(data[0]))
	for i := range output {
		output[i] = make([]byte, len(data))
	}

	//fill output table
	for i := 0 ; i < len(data) ; i ++ {
		for j := 0 ; j < len(data[i]) ; j++{	
			output[j][i] = data[i][j]
		}
	}

	//tostring
	res := make([]string, len(output))
	for i, r := range output {
		res[i] = strings.Trim( string(r), "\x00") //trim away nil bytes !
	}

	return res
}