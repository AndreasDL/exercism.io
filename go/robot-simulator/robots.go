package robot

import "fmt"

//this exercise is designed HORRIBLY, why global var Step1Robot ?
//Advance() => advance what ?

const (
	N Dir = iota
	E
	S
	W
)
type Dir int
var _ fmt.Stringer = Dir(1729)

var Step1Robot struct {
	X, Y int
	Dir
}

func Advance(){
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}
func Left(){
	Step1Robot.Dir = (Step1Robot.Dir + 3 ) % 4
}
func Right(){
	Step1Robot.Dir = (Step1Robot.Dir + 1 ) % 4
}

func (d Dir) String() string{
	switch Step1Robot.Dir {
	case N:
		return "N"
	case E:
		return "E"
	case S:
		return "S"
	case W:
		return "W"
	}
	return ""
}