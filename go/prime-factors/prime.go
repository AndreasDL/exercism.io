package prime


func Factors(v int64)[]int64{
	factors := []int64{}
	
	//%2 => afterwars all primes are even so we can progress with f+= 2 instead of f++	
	for v % int64(2) == 0 {
		v /= int64(2)
		factors = append(factors, int64(2))
	}


	for f := int64(3) ; f * f <= v ; {
		if v % f == 0{
			v /= f
			factors = append(factors, f)
		} else {
			f+= 2
		}
	}

	//this is also a prime factor ! 
	if v != 1 { 
		factors = append(factors, v)
	}


	return factors
}