package luhn

import (
	"strings"
	"strconv"
	"unicode"
	"fmt"
)



/*
4539 1488 0343 6467

The first step of the Luhn algorithm is to double every second digit,
starting from the right. We will be doubling

4_3_ 1_8_ 0_4_ 6_6_

If doubling the number results in a number greater than 9 then subtract 9
from the product. The results of our doubling:

8569 2478 0383 3437
 
Then sum all of the digits:

8+5+6+9+2+4+7+8+0+3+8+3+3+4+3+7 = 80

If the sum is evenly divisible by 10, then the number is valid. This number is valid!
## Example 2: invalid credit card number

8273 1232 7352 0569

Double the second digits, starting from the right

7253 2262 5312 0539

Sum the digits

7+2+5+3+2+2+6+2+5+3+1+2+0+5+3+9 = 57

57 is not evenly divisible by 10, so this number is not valid.
*/

func containsLetters( s string) bool{
	for _, letter := range s{
		if unicode.IsLetter(letter){
			return true
		}
	}
	return false
}

func toNumbers(s string) []int{
	numbers := []int{}
	for _, c := range s {
		num, _ := strconv.Atoi(string(c))
		numbers = append(numbers, num)
	}

	return numbers
}

func doubleEverySecondInPlace(n *[]int){	
	for i := len(*n) - 2 ; i >= 0 ; i -= 2{

		(*n)[i] = ( (*n)[i]*2 ) 
		if (*n)[i] > 9 {
			(*n)[i] -= 9
		}
	}
}

func sumAllDigits(n []int) int{
	sum := 0

	for _, i := range n {
		sum += i
	}

	return sum
}


func Valid(s string) bool{
	
	s = strings.Replace(s, " ", "", -1)

	if len(s) <= 1 {
		return false
	} else if containsLetters(s) {
		return false
	}

	n := toNumbers(s)
	fmt.Println(n)

	doubleEverySecondInPlace(&n)
	fmt.Println(n)

	res := sumAllDigits(n) % 10 == 0
	fmt.Println(res)
	fmt.Println("")
	return res
}