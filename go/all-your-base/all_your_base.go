package allyourbase

import (
	"errors"
	//"fmt"
)

func ConvertToBase(inBase int, input []int, outBase int) ([]int, error){
	if inBase < 2 {
		return nil, errors.New("input base must be >= 2")
	} else if outBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	//to base 10
	base10 := 0
	val := 1 //keep track of exponent
	for i := len(input)-1 ; i >= 0 ; val, i = val*inBase, i-1 {

		if input[i] < 0 || input[i] >= inBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		base10 += input[i] * val
	}
	
	//to output base
	res := []int{}
	for base10 > 0 {
		digit := base10 % outBase
		base10 -= digit
		base10 /= outBase

		res = append(res, digit)
	}
	reverse(&res)

	if len(res) == 0 {
		res = []int{0}
	}
	
	return res, nil
}

func reverse(arr *[]int){
	a := *arr
	for i, j := 0, len(a)-1 ; i < j ; i,j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}