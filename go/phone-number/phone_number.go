package phonenumber

import (
	"unicode"
	"errors"
)

func Number(s string) (string, error){

	if len(s) < 10 { return "", errors.New("A valid number should be at least 10 chars long")}
	
	//remove all non alphenumerical keys
	res := ""
	for _, c := range s {

		if unicode.IsDigit(c) {
			res += string(c)
		}
	}

	if len(res) == 11 && res[0] == '1' { res = res[1:] } //remove area code when present
	
	if len(res) != 10 { //check length
		return "", errors.New("A valid number should be 10 digits long, possibly starting with the country code ")
	}

	if res[0] == '1' || res[0] == '0' { return "", errors.New("area code should be 2-9")}
	if res[3] == '1' || res[3] == '0' { return "", errors.New("exchange code should be 2-9")}

	return res, nil
}

func AreaCode(s string) (string,error){
	res, err := Number(s)
	
	if err != nil { return "", err }
	
	return res[:3], err
}

func Format(s string) (string, error){

	res, err := Number(s)

	if err != nil { return "", err }

	return "(" + res[:3] + ") " + res[3:6] + "-" + res[6:], nil
}