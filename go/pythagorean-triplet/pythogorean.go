package pythagorean

import (
	"sort"
)


type Triplet [3]int

func Range(min, max int) []Triplet{
	ts := []Triplet{}

	for i := min ; i <= max ; i++ {
		for j := min ; j <= i ; j++ {
			for k := min ; k <= j ; k++ {

				if isTriplet(k,j,i) {
					ts = append(ts, Triplet{k,j,i})
				}

			}
		}
	}

	return ts	
}

func isTriplet(a,b,c int) bool{
	
	sides := []int{a*a,b*b,c*c}
	sort.Ints(sides)

	return sides[0] + sides[1] == sides[2]
}

func Sum(p int) []Triplet{

	res := []Triplet{}

	r := Range(1,p)
	for i := len(r) -1 ; i >= 0 ; i-- {

		if sumTriplet(r[i]) == p {
			res = append(res, r[i])
		}

	}

	return res
}

func sumTriplet(t Triplet) int {
	return t[0] + t[1] + t[2]
}