package palindrome

import (
	"strconv"
	"errors"
	"fmt"
)

type Product struct {
	Product int
	Factorizations [][2]int
}


func Products(fmin, fmax int) (Product, Product, error){
	//fmt.Println(fmin, " ", fmax)
	if fmin > fmax { return Product{}, Product{}, errors.New("fmin > fmax") }

	palindromes := map[int]*Product{} 
	//e.g. palindromes[9] => Product{9, [][2]int{1,9}, {3,3}}
	//=> allows us to store different factorizations at right location

	for i := fmin ; i <= fmax ; i++ {
		for j := i ; j <= fmax ; j++ {

			p := i * j
			if isPalindrome(p) {
				
				_, exists := palindromes[p]
				if ! exists {
					palindromes[p] = new(Product)
					palindromes[p].Product = p
				}
				
				palindromes[p].Factorizations = append(palindromes[p].Factorizations, [2]int{i,j})
			}
		}
	}

	//printMap(&palindromes)
	if len(palindromes) == 0 { return Product{}, Product{}, errors.New("no palindromes") }
	
	pmin, pmax := &Product{} , &Product{}
	initMin, initMax := true, true //fill with first value
	for k, v := range palindromes {
		if k < pmin.Product || initMin {
			pmin = v
			initMin = false
		}
		if k > pmax.Product || initMax {
			pmax = v
			initMax = false
		}
	}

	return *pmin, *pmax, nil
}

//debug
func printMap(ptr *map[int]*Product){
	m := *ptr

	for k, v := range(m){
		fmt.Println("\t", k, " => ", *v)
	}
}


func isPalindrome(i int) bool{
	if i < 0 { i = -i } //bonus exercise

	s := strconv.Itoa(i)
	for i, j := 0, len(s) -1 ; i < int(len(s) / 2) ; i, j = i+1, j-1 {
		if s[i] != s[j] { return false }
	}

	return true
}