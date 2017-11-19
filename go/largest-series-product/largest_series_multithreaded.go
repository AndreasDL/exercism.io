package lsproduct


/*
import (
	"strconv"
	"errors"
)


func LargestSeriesProduct(input string, span int) (int, error){
	
	if span < 0 { return -1, errors.New("span should be >= 0") }

	ptr, err := parseString(&input)
	if err != nil { return -1, err }

	list := *ptr
	if span > len(list) { return -1, errors.New("span must not be wider than the string itself") }

	
	ch := make(chan int, 100)
	go func(){
		generateSubstrings(ptr, span, ch)
	}()

	max := 0
	for prod := range ch {
		if prod > max {
			max = prod
		}
	}

	return max, nil
}


func parseString(input *string) (*[]int, error){
	s := *input

	res := make([]int, len(s))
	for i := 0 ; i < len(s) ; i++ {

		d, err := strconv.Atoi( string(s[i]) )
		if err != nil {
			return nil, errors.New("invalid input")
		}

		res[i] = d
	}

	return &res, nil
}

func generateSubstrings(in *[]int, span int, ch chan int){
	list := *in

	//start parallel calculations
	for start, stop := 0, span ; stop <= len(list) ; start, stop = start+1, stop+1 {
		
		res := 1
		for _,d := range list[start:stop] {
			res *= d
		}

		ch <- res
	}

	close(ch)
}
*/