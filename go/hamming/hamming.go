package hamming

import "errors"

func Distance(a, b string) (int, error) {

	len_a := len(a)
	len_b := len(b)

	if len_a != len_b	{
		return -1, 
			errors.New("Hamming distance is not defined for strings of different length")
	}

	cnt := 0
	for i := 0 ; i < len_a ; i++ {
		if a[i] != b[i] {
			cnt++
		}
	}

	return cnt, nil
}
