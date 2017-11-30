package strain

type Ints []int
type Lists [][]int
type Strings []string


func(values *Ints) Keep(f func(int)bool) Ints{
	var res Ints

	for _, v := range *values{
		if f(v) { res = append(res, v) }
	}

	return res
}

func(values *Ints) Discard(f func(int)bool) Ints{
	var res Ints

	for _, v := range *values{
		if !f(v) { res = append(res, v) }
	}

	return res
}

func(values *Lists) Keep(f func([]int)bool) Lists{
	var res Lists

	for _, v := range *values{
		if f(v) { res = append(res, v) }
	}

	return res
}

func(values *Strings) Keep(f func(string)bool) Strings{
	var res Strings

	for _, v := range *values{
		if f(v) { res = append(res, v) }
	}

	return res
}