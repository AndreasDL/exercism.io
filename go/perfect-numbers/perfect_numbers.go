package perfect

import(
	"errors"
	"sort"
	"math"
)




type Classification int
const (
	ClassificationDeficient = iota
	ClassificationPerfect   = iota
	ClassificationAbundant  = iota
)

var ErrOnlyPositive = errors.New("input should be positive !")


func Classify(number int64) (Classification,error){
	if number <= 0 { 
		return -1, ErrOnlyPositive 
	} else if number == 1 { 
		return ClassificationDeficient, nil 
	}

	primeFactors := factors(number)
	primeFactors = primeFactors[:len(primeFactors)-1] //remove number itself

	sum := int64(0)
	for _, p := range primeFactors {
		sum += p
	}

	if sum == number { 
		return ClassificationPerfect, nil 
	} else if sum < number { 
		return ClassificationDeficient, nil 
	} else {
		return ClassificationAbundant, nil 
	}
}


func factors(n int64) []int64{

	pfs := map[int64]bool{ //unique factors !
		1: true,
		n: true,
	}

	for i := int64(2) ; i <= int64(math.Sqrt(float64(n))) ; i++ {
		
		if n % i == 0 {
			pfs[i] = true
			pfs[n/i] = true
		}
	}

	//convert keys to array
	res := []int64{}
	for k := range pfs {
		res = append(res, k)
	}

	sort.Slice(res, func(i,j int)bool{
		return res[i] < res[j]
	})

	return res
}