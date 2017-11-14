package tournament


import (
	"io"
	"bufio"
	"strings"
	"sort"
	"fmt"
	"errors"
)

type team struct {
	name string
	mp, w, d, l, p int
}

func (t *team) updateTeam(teamname string, outcome string){
	if outcome == "draw"{
		t.mp++
		t.d++
		t.p++
	} else if outcome == "win" {
		t.mp++
		t.w++
		t.p += 3
	} else if outcome == "loss" {
		t.mp++
		t.l++
	}
}

//team1 wins => team 2 loses
func (t *team) updateReverse(teamname string, outcome string){
	if outcome == "win" {
		t.updateTeam(teamname, "loss")
	} else if outcome == "loss" {
		t.updateTeam(teamname, "win")
	} else {
		t.updateTeam(teamname, outcome)
	}
}

//returns sorted list of input
func getInputList(scanner *bufio.Scanner) (*[]*team, error) {
	teams := map[string]*team{} //teams[teamname] ==> returns team stats
	//https://stackoverflow.com/questions/13101298/calling-a-pointer-method-on-a-struct-in-a-map => relevance++

	for scanner.Scan() {

		split := strings.Split(scanner.Text(), ";") // split = [ team1 , team2 , outcome ]
		if len(split) == 3 {

			if (split[2] != "win" && split[2] != "loss" && split[2] != "draw"){
				return nil, errors.New("outcome should be either win, loss or draw. ( scanner.Text() )")
			}

			if _, ok := teams[split[0]] ; ! ok {
				teams[split[0]] = &team{name : split[0]}
			}

			if _, ok := teams[split[1]] ; ! ok {
				teams[split[1]] = &team{name : split[1]}
			}

			(*teams[split[0]]).updateTeam(split[0], split[2])
			(*teams[split[1]]).updateReverse(split[1], split[2])	
		} else if len(split) != 3 && len(split) != 1 {
			return nil, errors.New("line has more than 3 arguments ( " + scanner.Text() + ")")
		}
	}

	lteams := []*team{}
	for _, v := range teams {
		lteams = append(lteams, v)
	}

	sort.Slice(lteams, func(i, j int) bool {
			if lteams[i].p > lteams[j].p {
				return true
			} else if lteams[i].p < lteams[j].p {
				return false
			} else {
				return lteams[i].name < lteams[j].name
			}
	})

	return &lteams, nil
}

func generateOutput(writer *bufio.Writer, teams *[]*team){
	
	writer.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, tp := range *teams {
		team := *tp //dereference
		writer.WriteString(
			fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n",
				team.name,
				team.mp,
				team.w,
				team.d,
				team.l,
				team.p,
			),
		)
	}
}

func Tally(r io.Reader,w io.Writer) error{

	scanner := bufio.NewScanner(r)	
	teams, err := getInputList(scanner)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(w)
	generateOutput(writer, teams)
	writer.Flush()


	return nil
}