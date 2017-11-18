package etl

import "strings"

func Transform(input map[int][]string ) map[string]int{
	res := map[string]int{}

	for k, v := range input {

		for _, i := range v {
			res[strings.ToLower(i)] = k
		}
	}


	return res
}