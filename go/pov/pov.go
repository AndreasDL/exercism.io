package pov

//import "fmt"

type Graph struct{
	//maps => O(1) lookups
	conn map[string]map[string]bool //conn[from] => map[to]
	nodes map[string]bool
}

func New() *Graph{
	return &Graph{
		conn : map[string]map[string]bool{},
		nodes: map[string]bool{},
	}
}

func (g *Graph) AddNode(nodelabel string){
	g.conn[nodelabel] = map[string]bool{}
	g.nodes[nodelabel] = true
}

func (g *Graph) AddArc(from, to string){

	//init labels if needed
	if _, fromExists := g.nodes[from] ; !fromExists { g.AddNode(from) }
	if _, toExists   := g.nodes[to]   ; !toExists   { g.AddNode(to  ) }

	g.conn[from][to] = true
}

func (g *Graph) ArcList() []string{
	
	res := []string{}
	for from, toMap := range g.conn {
		for to, _ := range toMap{
			res = append(res, from +  " -> " + to)
		}
	}

	return res
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	return nil
}