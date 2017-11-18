package pascal


func Triangle(endLine int) [][]int{
	res := [][]int{}

	for n := 0 ; n < endLine ; n++{
		
		line := make([]int, n+1)

		for k := 0; k <= n ; k++ {
			//line[k] =  getCoef(n, k, factorialClassic) //slow method
			line[k] =  getCoef(n, k, factorialFast) //Fast Method

		}

		res = append(res, line)
	}

	return res
}


func getCoef(n, k int, factorial func(int) int) int {
	if k == 0 || n - k == 0 {
		return 1
	}

	return factorial(n) / ( factorial(k) * factorial(n-k) )

}

func factorialClassic(n int) int{ //room for optimization ! => see below !
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else {
		return n * factorialClassic(n-1)
	}
}

//keep track of previously calculated factorials => speedup using a cache => 66% reduction in execution time :D
var facts []int
func init(){
	facts = make([]int, 2)
	facts[0] = 1
	facts[1] = 1
}


func factorialFast(n int) int{
	
	if n < len(facts){
		return facts[n]
	}else {

		for i := len(facts) ; i <= n ; i ++ {
			facts = append(facts, i * facts[i-1] )
		}
		
		return facts[n]
	}
}
