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

//teams[teamname] ==> returns team stats
func parseInput(scanner *bufio.Scanner) (*map[string]*team, error) {
	teams := map[string]*team{} //teams[teamname] ==> returns team stats

	for scanner.Scan() {

		split := strings.Split(scanner.Text(), ";") // split = [ team1 , team2 , outcome ]
		if len(split) == 3 {

			if (split[2] != "win" && split[2] != "loss" && split[2] != "draw"){
				return nil, errors.New("outcome should be either win, loss or draw. ( scanner.Text() )")
			}

			if _, ok := teams[split[0]] ; ! ok {
				teams[split[0]] = &team{}
			}

			if _, ok := teams[split[1]] ; ! ok {
				teams[split[1]] = &team{}
			}

			(*teams[split[0]]).updateTeam(split[0], split[2])
			(*teams[split[1]]).updateReverse(split[1], split[2])	
		} else if len(split) != 3 && len(split) != 1 {
			return nil, errors.New("line has more than 3 arguments ( " + scanner.Text() + ")")
		}
	}

	return &teams, nil
}

func generateOutput(writer *bufio.Writer, teams *map[string]*team){
	//bit ugly, but it does the trick
	
	//maps score => list of teams with that score
	score_to_team := map[int][]string{}
	for k, v := range *teams {
		score_to_team[v.p] = append(score_to_team[v.p], k)
	}

	//get the different scores
	scores := []int{}
	for s, _ := range score_to_team {
		scores = append(scores, s)
	}

	//sort scores
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	writer.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, score := range scores {
		//sort the list of team names alphabetically
		sort.Strings( score_to_team[score] )

		for _, teamname := range score_to_team[score] {
			team := (*teams)[teamname]
			writer.WriteString(
				fmt.Sprintf("%-30s |%3d |%3d |%3d |%3d |%3d\n",
					teamname,
					team.mp,
					team.w,
					team.d,
					team.l,
					team.p,
				),
			)
		}
	}
	//writer.WriteString("\n")
}

func Tally(r io.Reader,w io.Writer) error{

	scanner := bufio.NewScanner(r)	
	teams, err := parseInput(scanner)

	if err != nil {
		return err
	}

	writer := bufio.NewWriter(w)
	generateOutput(writer, teams)
	writer.Flush()


	return nil
}