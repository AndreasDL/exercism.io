package sublist

type Relation string


func Sublist(a, b []int) Relation {

	//quick checks
	if len(a) == 0 && len(b) == 0 {
		return "equal"
	} else if len(a) == 0 {
		return "sublist"
	} else if len(b) == 0 {
		return "superlist"
	}

	//sublist => short in long one
	short := *(&a)
	long  := *(&b)
	swapped := false
	if len(b) < len(a) {
		short = *(&b)
		long  = *(&a)
		swapped = true
	}
	
	res := isSublist(short, long)
	if swapped && res == "sublist"{
		res = "superlist"
	}

	return res
}

func isSublist(short, long []int) Relation{
	i := 0
	j := 0
	for i < len(long) {

		ti := i
		tj := j
		for tj < len(short) && ti < len(long) && long[ti] == short[tj] {
			ti++
			tj++
		}

		if tj == len(short) {
			if len(short) == len(long){
				return "equal"
			}
			return "sublist"
		}

		i++
		j = 0
	}

	return "unequal"

}