package romannumerals

import "errors"

var aToR = []struct{
	arabic int
	roman string
}{
	{1   ,"I"}, {4  ,"IV"}, {5  ,"V"}, {9  ,"IX"},
	{10  ,"X"}, {40 ,"XL"}, {50 ,"L"}, {90 ,"XC"},
	{100 ,"C"}, {400,"CD"}, {500,"D"}, {900,"CM"},
	{1000,"M"},
}


func ToRomanNumeral(a int) (string, error){

	if a <= 0 || a > 3000 { return "", errors.New("number out of range")}

	res := ""
	for i := len(aToR) - 1 ; i >= 0 ; i-- {

		for a >= aToR[i].arabic { 
			a -= aToR[i].arabic 
			res += aToR[i].roman
		}

	}

	return res, nil
}