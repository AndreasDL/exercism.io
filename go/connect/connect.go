package connect

type point struct{ x, y int }
var DIRECTIONS = []point{ 
	{  0,  1},
	{  1,  0},
	{  0, -1},
	{ -1,  0},
	{ -1,  1}, //valid diag!
	{  1, -1}, //valid diag!
}

type Board struct {
	locations [][]byte
}
func NewBoard(s []string) *Board{

	b := Board{ locations: make([][]byte, len(s)) }

	for y, line := range s {
		b.locations[y] = make([]byte, len(line))
		for x, c := range []byte(line) {
			b.locations[y][x] = c
		}
	}

	return &b
}
func (b *Board) traceHor(c byte) (hor bool) {
	//init the todos 
	todos := stack{ &point{0, 0} }
	for i := 1 ; i < len(b.locations)-1 ; i++ { todos.push(&point{0, i}) }  //left side, no edges
	
	seen := map[point]bool{}
	for pos := todos.pop() ; pos != nil && !hor ; pos = todos.pop(){
		p := *pos

		//check if we passed a border
		if p.x >= len(b.locations[0]) { hor = true }
		
		//outside board, seen or char not correct => skip
		if p.y < 0 || p.x < 0 || 
			p.y >= len(b.locations) || p.x >= len(b.locations[0]) || 
			seen[p] || b.locations[p.y][p.x] != c {
			continue 
		}
		seen[p] = true

		//push neighbors
		for _, delta := range DIRECTIONS {
			todos.push(&point{
				x: p.x+delta.x,
				y: p.y+delta.y,
			})
		}
	}

	return
}
func (b *Board) traceVer(c byte) (ver bool) {
	//init the todos 
	todos := stack{ &point{0, 0} }
	for i := 1 ; i < len(b.locations)-1 ; i++ { todos.push(&point{i, 0}) }

	seen := map[point]bool{}
	for pos := todos.pop() ; pos != nil && !ver ; pos = todos.pop(){
		p := *pos

		//check if we passed a border
		if p.y >= len(b.locations){ ver = true }
		
		//outside board, seen or char not correct => skip
		if p.y < 0 || p.x < 0 || 
			p.y >= len(b.locations) || p.x >= len(b.locations[0]) || 
			seen[p] || b.locations[p.y][p.x] != c {
			continue 
		}
		seen[p] = true

		//push neighbors
		for _, delta := range DIRECTIONS {
			todos.push(&point{
				x: p.x+delta.x,
				y: p.y+delta.y,
			})
		}
	}
	
	return
}

func ResultOf(board []string) (string, error) {
	b := *NewBoard(board)

	//get scores	
	horX := b.traceHor('X')
	verX := b.traceVer('X')
	x := 0 ; if horX { x++ } ; if verX { x++ }

	horO := b.traceHor('O')
	verO := b.traceVer('O')
	o := 0 ; if horO { o++ } ; if verO { o++ }

	//Return winner
	if x > o { 
		return "X", nil
	} else if x < o {
		return "O", nil
	}

	return "", nil
}



type stack []*point
func (s *stack) pop() *point{
	l := len(*s)
	if l == 0 {
		return nil
	} else { 
		c := (*s)[l-1]
		*s = (*s)[:l-1] 
		return c
	}
}
func (s *stack) push(c *point) { *s = append(*s, c) }