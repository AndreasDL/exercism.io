package lsproduct


import (
	"strconv"
	"errors"
)


func LargestSeriesProduct(input string, span int) (int, error){
	
	if span < 0 { return -1, errors.New("span should be >= 0") }

	ptr, err := parseString(&input)
	if err != nil {
		return -1, err
	}

	list := *ptr
	if span > len(list) {
		return -1, errors.New("span must not be wider than the string itself")
	}

	return generateSubstrings(ptr, span), nil
}


func parseString(input *string) (*[]int, error){
	s := *input

	res := make([]int, len(s))
	for i := 0 ; i < len(s) ; i++ {

		d, err := strconv.Atoi( string(s[i]) )
		if err != nil {
			return nil, errors.New("invalid input")
		}

		res[i] = d
	}

	return &res, nil
}

func generateSubstrings(in *[]int, span int) int{
	list := *in
	
	max := 0
	for start, stop := 0, span ; stop <= len(list) ; start, stop = start+1, stop+1 {

		prod := calculateProduct(list[start:stop])

		if prod > max {
			max = prod
		}
	}

	return max
}

func calculateProduct(list []int) int{
	
	res := 1
	for i := 0 ; i < len(list) ; i++ {
		res *= list[i]
	}

	return res
}