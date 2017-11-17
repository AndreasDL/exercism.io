package accumulate


func Accumulate(given []string, converter func(string) string) []string {
	res := make([]string, len(given))

	for i, s := range given{
		res[i] = converter(s) 
	}

	return res
}