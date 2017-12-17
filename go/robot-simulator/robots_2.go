package robot

const (
	R Command = iota
	L
	A
)
type Command byte // valid values are 'R', 'L', 'A'

type RU int
type Pos struct{ Easting, Northing RU }
type Rect struct{ Min, Max Pos }
type Step2Robot struct {
	Dir
	Pos
}

const (
	advance Action = iota
	rotateLeft
	rotateRight
	done
)
type Action int


func StartRobot(cmd chan Command, act chan Action){
	for {
		c, ok := <-cmd

		if !ok {
			close(act)
			return
		}

		switch c {
		case 'R':
			act <- rotateRight
		case 'L':
			act <- rotateLeft
		case 'A':
			act <- advance
		}
	}
}

func Room(r Rect, robot Step2Robot, act chan Action, rep chan Step2Robot){
	for {

		a, ok := <-act

		if !ok {
			rep <- robot
			close(rep)
			return
		}

		switch a {
		case advance:
			robot.advance(r)
		case rotateLeft:
			robot.rotateLeft()
		case rotateRight:
			robot.rotateRight()
		}
	}
}

func (r *Step2Robot) advance(room Rect){
	switch r.Dir{
		case N:
			if r.Northing < room.Max.Northing { r.Northing++ }
		case E:
			if r.Easting  < room.Max.Easting  { r.Easting++  }

		case S:
			if r.Northing > room.Min.Northing {	r.Northing-- }
		case W:
			if r.Easting  > room.Min.Easting  { r.Easting--  }
	}
}
func (r *Step2Robot) rotateLeft(){
	r.Dir = (r.Dir + 3) % 4
}
func (r *Step2Robot) rotateRight(){
	r.Dir = (r.Dir + 1) % 4		
}