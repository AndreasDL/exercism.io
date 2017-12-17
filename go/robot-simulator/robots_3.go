package robot

import (
	"strings"
	"fmt"
	"errors"
)
//got stuck during this part, thanks to https://github.com/macpla/exercism_go/blob/master/robot-simulator/robot_simulator_step3.go for a solution !


type Step3Robot struct {
	Name string
	Step2Robot
}

type Action3 struct{
	Name string
	Action//reuse action from step2
}

func StartRobot3(name, script string, action chan Action3, log chan string){
loop:	
	for _, instr := range strings.Split(script, ""){
		switch instr[0] {
		case 'L':
			action <- Action3{name, rotateLeft  }
		case 'R':
			action <- Action3{name, rotateRight }
		case 'A':
			action <- Action3{name, advance     }
		default:
			log <- "undefined command in script!"
			break loop
		}
	}
	action <- Action3{name, done}
}

func (r *Step2Robot) isInRect(rect Rect) bool {
	if r.Easting < rect.Min.Easting ||
		r.Easting > rect.Max.Easting ||
		r.Northing < rect.Min.Northing ||
		r.Northing > rect.Max.Northing {
		return false
	}
	return true
}

var scene refFrame
func Room3(extent Rect, robots []Step3Robot, action <-chan Action3, report chan<- []Step3Robot, log chan<- string) {
	robotsByName, err := scene.spawn(extent, robots)
	if err != nil {
		log <- err.Error()
	}
	robotsByPosition, err := scene.orient(extent, robotsByName)
	if err != nil {
		log <- err.Error()
	}
	nRunningRobots := len(robots)
	defer func() {
		report <- robots
		close(report)
	}()
	for nRunningRobots > 0 {
		act := <-action
		switch act.Action {
		case advance:
			rob, ok := robotsByName[act.Name]
			if !ok {
				log <- "An action from an unknown robot"
			} else if err := scene.advanceRobot(rob, extent, robotsByPosition); err != nil {
				log <- fmt.Sprint(err)
			}
		case rotateLeft:
			robotsByName[act.Name].rotateLeft()
		case rotateRight:
			robotsByName[act.Name].rotateRight()
		case done:
			nRunningRobots--
		}
	}
}

func (p *Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.Easting, p.Northing)
}

var errAdvanceOnRobot = errors.New("robot attempting to advance into antother robot")
var errAdvanceOnWall = errors.New("robot attempting to advance into a wall")

type refFrame struct{}
func (*refFrame) advanceRobot(robot *Step3Robot, room Rect, robotsByPositions map[Pos]*Step3Robot) error {
	futurePos := robot.Pos
	switch robot.Dir {
	case N:
		futurePos.Northing++
	case E:
		futurePos.Easting++
	case S:
		futurePos.Northing--
	case W:
		futurePos.Easting--
	}
	if _, ok := robotsByPositions[futurePos]; ok {
		return errAdvanceOnRobot
	}
	if futurePos.Easting < room.Min.Easting ||
		futurePos.Easting > room.Max.Easting ||
		futurePos.Northing < room.Min.Northing ||
		futurePos.Northing > room.Max.Northing {
		return errAdvanceOnWall
	}
	delete(robotsByPositions, robot.Pos)
	robot.Pos = futurePos
	robotsByPositions[robot.Pos] = robot
	return nil
}

func (*refFrame) spawn(room Rect, robots []Step3Robot) (rMap map[string]*Step3Robot, err error) {
	rMap = make(map[string]*Step3Robot)
	for i, r := range robots {
		if _, ok := rMap[r.Name]; !ok {
			if len(r.Name) == 0 {
				err = fmt.Errorf("Robot without a name")
			}
			rMap[r.Name] = &robots[i]
		} else {
			err = fmt.Errorf("Duplicate robot names")
		}
	}
	return rMap, err
}

func (*refFrame) orient(room Rect, robotsByName map[string]*Step3Robot) (map[Pos]*Step3Robot, error) {
	robotsByPositions := make(map[Pos]*Step3Robot, len(robotsByName))
	var err error
	for k, v := range robotsByName {
		if _, ok := robotsByPositions[v.Pos]; ok {
			err = fmt.Errorf("robots placed at the same place")
		} else {
			if !v.isInRect(room) {
				err = fmt.Errorf("a robot placed outside of the room")
			}
			robotsByPositions[v.Pos] = robotsByName[k]
		}
	}
	return robotsByPositions, err
}