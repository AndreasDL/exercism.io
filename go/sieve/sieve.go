package sieve

func Sieve(limit int) []int{
	if limit < 2 {
		return []int{}
	}
	
	candidates := make([]bool, limit+1) 
	//default false => prime
	//visited => true meaning not a prime

	candidates[0] = true
	candidates[1] = true
	
	primes := []int{}		
	for p := 2 ; p <= limit ; p++ {
		
		if ! candidates[p] { //prime number => remove others
			markMultiplesInPlace(&candidates, p)
			primes = append(primes, p)
		}
	}

	return primes
}


func markMultiplesInPlace(ptr *[]bool, p int){
	candidates := *ptr

	for i := p+p ; i < len(candidates) ; i += p {
		candidates[i] = true
	}
}