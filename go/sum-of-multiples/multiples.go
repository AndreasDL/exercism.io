package summultiples


func SumMultiples(limit int, divisors ...int) int{

	multiples := map[int]int{} //map => unique keys => unique multiples
	for _, div := range divisors {		
		i := div 
		for i < limit {
			
			multiples[i]++
			i += div
		}
	}


	sum := 0
	for k, _ := range multiples {
		sum += k
	}

	return sum
}