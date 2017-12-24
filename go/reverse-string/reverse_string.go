package reverse

func String(s string) string {

	res := []byte(s)

	for i,j := 0, len(s)-1 ; i<j ; i,j= i+1, j-1 {
		res[i], res[j] = res[j],res[i]
	}

	return string(res)
}