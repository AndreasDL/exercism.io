package kindergarten

import (
	"strings"
	"sort"
	"errors"
)

var symbolToPlant = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

type Garden struct{
	lines []string
	children []string
}

func checkNames(children []string) bool{
	childCtr := map[string]int{}

	for _, c := range children { 
		childCtr[c]++ 
		if childCtr[c] > 1 { return false }
	}
	
	return true
}

func checkPlants(lines []string) bool{

	for _, line := range lines {
		for _, p := range line {
			if _, ex := symbolToPlant[p] ; !ex { return false }
		}
	}

	return true
}


func NewGarden(diagram string, children []string) (*Garden, error){
	lines := strings.Split(diagram, "\n")
	if len(lines) != 3 { 
		return nil, errors.New("there are suppose to be two lines with leading newline, so ... 3")
	} else if len(children)*2 != len(lines[1]){
		return nil, errors.New("children diagram mismatch")
	} else if !checkNames(children){
		return nil, errors.New("duplicate names detected!")
	} else if !checkPlants(lines[1:]){
		return nil, errors.New("unknown plants")
	}

	childs := make([]string, len(children))
	for i, c := range children { childs[i] = c }
	sort.Strings(childs)

	g := Garden{
		lines   : lines[1:],
		children: childs,
	}


	return &g, nil
}

func (g *Garden) positionOnRow(child string) (int,int){
	i :=0
	for ; i < len(g.children) && g.children[i] != child ; i++ {}

	if i == len(g.children) { return -1, -1 }
	return 2*i, 2*i+2
}

func (g *Garden) Plants(child string) ([]string, bool){
	res := []string{}

	b, e := g.positionOnRow(child)
	if b < 0 || e < 0 { return []string{}, false }

	for _, line := range g.lines {
		for _, p := range line[b:e] { res = append(res, symbolToPlant[p]) }
	}

	return res, true
}