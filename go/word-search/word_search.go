package wordsearch

import (
	//"fmt"
	"strings"
	"errors"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error){
	positions := &map[string][2][2]int{}

	//horizontal search => easy
	find(&puzzle, 
		words, 
		positions, 
		func(x,y,w int,r bool)[2][2]int{return getLocations(x,y,w,false,r)},
	)
	filterWords(&words, positions)

	//vertical search => field
	puzzleTransposed := transposePuzzle(&puzzle)
	find(puzzleTransposed, 
		words, 
		positions, 
		func(x,y,w int,r bool)[2][2]int{return getLocations(x,y,w,true,r)},
	)
	filterWords(&words, positions)

	//diagonal search
	puzzleDiag := diagonalize(&puzzle)
	find(puzzleDiag, 
		words, 
		positions, 
		func(x,y,w int,r bool)[2][2]int{return getDiagLocations(x,y,w,len(puzzle),r) },
	)
	filterWords(&words, positions)

	//rev diagonal search
	puzzleDiag = revDiagonalize(&puzzle)
	find(puzzleDiag, 
		words, 
		positions, 
		func(x,y,w int,r bool)[2][2]int{return getRevDiagLocations(x,y,w,len(puzzle),r) },
	)
	filterWords(&words, positions)

	//any remaining ?
	if len(words) > 0 {
		//fmt.Println(words)
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

func find(fld *[]string, words []string, positions *map[string][2][2]int, 
	locator func(int,int,int,bool)[2][2]int) {
	field := *fld
	pos := *positions

	for _, w := range words {
		for y, line := range field {
			
			if x := strings.Index(line, w) ; x >= 0 { 
				pos[w] = locator(x,y,len(w), false)
			} else if x := strings.Index(line, reverse(w)) ; x >= 0 {
				pos[w] = locator(x,y,len(w), true)
			}

		}
	}
}
func getLocations(x,y,w int, transposed, reversed bool) [2][2]int{

	xMin, xMax := x, x+w-1
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

func diagonalize(puzzle *[]string) *[]string{
	p := *puzzle

	res := make([]string, 2*len(p))
	i := 0
	for startX := len(p)-1; startX >= 0 ; startX, i = startX-1, i+1 {
		
		line := []byte{}
		for y, x := 0, startX ; y < len(p) && x < len(p) ; y,x = y+1, x+1 {
			line = append(line, p[y][x])
		}
		res[i] = string(line)
	}

	for startY := 1 ; startY < len(p) ; startY, i = startY+1, i+1 {

		line := []byte{}
		for y, x := startY, 0 ; y < len(p) && x < len(p) ; y,x = y+1, x+1 {
			line = append(line, p[y][x])
		}
		res[i] = string(line)
	}

	return &res
}
func getDiagLocations(x,y,w,d int, reversed bool) [2][2]int{
	xMin, yMin := x, y
	if y < d {
		xMin = d - y -1
		yMin = x
	} else {
		yMin = y - d +1 +x
	}

	xMax, yMax := xMin+w-1, yMin+w-1

	if reversed {
		xMin, xMax = xMax, xMin
		yMin, yMax = yMax, yMin
	}

	return [2][2]int{
		{xMin, yMin},
		{xMax, yMax},
	}
}

func revDiagonalize(puzzle *[]string) *[]string{
	p := *puzzle

	res := make([]string, 2*len(p))
	i := 0
	for startX := 0; startX < len(p) ; startX, i = startX+1, i+1 {
		
		line := []byte{}
		for y, x := 0, startX ; y < len(p) && x >= 0 ; y,x = y+1, x-1 {
			line = append(line, p[y][x])
		}
		res[i] = reverse(string(line))
	}

	for startY := 1 ; startY < len(p) ; startY, i = startY+1, i+1 {

		line := []byte{}
		for y, x := startY, len(p)-1 ; y < len(p) && x >= 0 ; y,x = y+1, x-1 {
			line = append(line, p[y][x])
		}
		res[i] = reverse(string(line))
	}

	return &res
}
func getRevDiagLocations(x,y,w,d int, reversed bool) [2][2]int{
	xMin, yMin := x, y
	if y < d {
		yMin = y - x
	} else {
		xMin = x + (y-d) +1
		yMin = d - 1 - x
	}

	xMax, yMax := xMin+w-1, yMin-w+1

	if reversed {
		xMin, xMax = xMax, xMin
		yMin, yMax = yMax, yMin
	}

	return [2][2]int{
		{xMin, yMin},
		{xMax, yMax},
	}
}