package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func Build(records []Record) (*Node, error){

	if len(records) == 0 || records == nil {
		return nil, nil
	}

	children, err := createChildrenMap(records)
	if err != nil {
		return nil, err
	}

	return createNode(0, children), nil
}

//builds a map that has following structure: 
//         children[parentid] = list of children id
//also does some checks on input
func createChildrenMap(records []Record) (*map[int][]int, error){
	children := map[int][]int{} //children[parentid] = list of children id
	minID, maxID := 0, 0 //we'll check these ones later
	
	for _, r := range records {
		if r.Parent > r.ID {
			return nil, errors.New("Parent id must be <= that the node's ID")

		} else if r.Parent < r.ID {
			if r.ID < minID {
				minID = r.ID
			} else if r.ID > maxID {
				maxID = r.ID
			}

			children[r.Parent] = append(children[r.Parent], r.ID)	
		} else if r.Parent == r.ID && r.ID != 0 {
			return nil, errors.New("Only the root node can have id == parentid")
		}
	}

	if minID < 0 || maxID >= len(records){
		return nil, errors.New("ID of node should be between 0 and the number of nodes")
	}

	return &children, nil
}

//recursively create nodes
func createNode(i int, children *map[int][]int) *Node{

	var childList []*Node = nil

	if len( (*children)[i] ) > 0 {
		childList = []*Node{}

		sort.Ints( (*children)[i] )
		for _, c := range (*children)[i] {
			childList = append(childList, createNode(c,children))
		}
	}

	return &Node{
		i,
		childList,
	}
}