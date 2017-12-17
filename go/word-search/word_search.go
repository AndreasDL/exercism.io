package wordsearch

import (
	"fmt"
	"strings"
	"errors"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error){
	positions := &map[string][2][2]int{}

	//horizontal search => easy
	findHorizontal(&puzzle, words, positions)
	filterWords(&words, positions)
	fmt.Println(positions)
	fmt.Println(words)
	fmt.Println()

	//vertical search => field


	//any remaining ?
	if len(words) > 0 {
		return *positions, errors.New("not all words were found !")
	}

	return *positions, nil
}

func filterWords(words *[]string, positions *map[string][2][2]int){
	for k, _ := range *positions {

		i := 0
		for i < len(*words) && (*words)[i] != k { i++ }
		
		if len(*words) > 0 && (*words)[i] == k {
			//delete word without preserving order
			(*words)[i] = (*words)[len(*words)-1] //replace with last element
			(*words) = (*words)[:len(*words)-1] //resize
		}
	}
}
func reverse(s string) string{
	res := []rune(s)
	for b, e := 0, len(s)-1 ; b < e ; b,e = b+1, e-1{
		res[b], res[e] = res[e], res[b]
	}
	return string(res)
}

func findHorizontal(fld *[]string, words []string, positions *map[string][2][2]int) {
	field := *fld
	pos := *positions

	for _, w := range words {
		for y, line := range field {
			
			if x := strings.Index(line, w) ; x >= 0 { 
				pos[w] = [2][2]int{
					[2]int{x,y},
					[2]int{x+len(w)-1,y},
				}
			} else if x := strings.Index(line, reverse(w)) ; x >= 0 {
				pos[w] = [2][2]int{
					[2]int{x+len(w)-1,y},
					[2]int{x,y},
				}
			}

		}
	}
}
