package wordsearch

import (
	"fmt"
	"strings"
	"errors"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error){
	positions := &map[string][2][2]int{}

	//horizontal search => easy
	find(&puzzle, words, positions, false)
	filterWords(&words, positions)

	//vertical search => field
	puzzleTransposed := transposePuzzle(&puzzle)
	find(puzzleTransposed, words, positions, true)
	filterWords(&words, positions)

	//diagonal search


	//any remaining ?
	if len(words) > 0 {
		fmt.Println(words)
		return *positions, errors.New("not all words were found !")
	}

	return *positions, nil
}

func filterWords(words *[]string, positions *map[string][2][2]int){
	for k, _ := range *positions {

		i := 0
		for i < len(*words) && (*words)[i] != k { i++ }
		
		if i < len(*words) {
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
func transposePuzzle(puzzle *[]string) *[]string{
	p := *puzzle

	//input empty field
	res := make([][]byte, len(p[0]))
	for i := range res { res[i] = make([]byte, len(p)) }

	for x, line := range p {
		for y, c := range []byte(line) { res[y][x] = c }
	}

	output := make([]string, len(res))
	for i, line := range res { output[i] = string(line) }

	return &output
}

func find(fld *[]string, words []string, positions *map[string][2][2]int, transposed bool) {
	field := *fld
	pos := *positions

	for _, w := range words {
		for y, line := range field {
			
			if x := strings.Index(line, w) ; x >= 0 { 
				pos[w] = getLocations(x,y,w, transposed, false)
			} else if x := strings.Index(line, reverse(w)) ; x >= 0 {
				pos[w] = getLocations(x,y,w, transposed, true)
			}

		}
	}
}
func getLocations(x,y int, w string, transposed, reversed bool) [2][2]int{

	xMin, xMax := x, x+len(w)-1
	yMin, yMax := y, y
	
	if transposed {
		xMin, yMin = yMin, xMin
		xMax, yMax = yMax, xMax
	}

	if reversed {
		xMin, xMax = xMax, xMin
		yMin, yMax = yMax, yMin
	}

	
	return [2][2]int{
		{xMin, yMin},
		{xMax, yMax},
	}
}