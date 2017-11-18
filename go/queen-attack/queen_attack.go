package queenattack

import (
	"errors"
	"strconv"
)

func CanQueenAttack(w, b string) (bool, error){

	if len(w) != 2 || len(b) != 2 || w == b {
		return false, errors.New("invalid input")
	}

	x1, y1 := toXY(w)
	x2, y2 := toXY(b)
	
	if x1 < 1 || x1 > 8 || y1 < 1 || y1 > 8 ||
			x2 < 1 || x2 > 8 || y2 < 1 || y2 > 8 {
		return false, errors.New("Position invalid !")
	}

	return y1 == y2 || //same row
			x1 == x2 || //same col
			diagonal(x1, y1, x2, y2), 
		nil
}


var charToNum map[byte]int
func init(){
	charToNum = map[byte]int {
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
	}
}

func toXY(s string) (x,y int){
	x = charToNum[s[0]]
	y, _ = strconv.Atoi(string(s[1]))

	return
}

func diagonal(x1, y1, x2, y2 int) bool{
	dx := x1 - x2
	dy := y1 - y2

	if dx < 0 { dx = -dx }
	if dy < 0 { dy = -dy }

	return dx == dy
}