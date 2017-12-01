package prime


func Nth(nth int) (int, bool){
	if nth <= 0 {return -1, false}

	i := 0
	v := 1
	for  i < nth {
		v++
		if isPrime(v){ i++ }
	}

	return v, true
}


func isPrime(v int) bool{
	//inspired by ../prime-factors


	//%2 => afterwars all primes are even so we can progress with f+= 2 instead of f++	
	for v % 2 == 0 {
		return v == 2
	}

	//if we find divisor that is not f => not prime 
	for f := 3 ; f * f <= v ; {
		if v % f == 0{
			return v == f
		} else {
			f+= 2
		}
	}

	return true
}