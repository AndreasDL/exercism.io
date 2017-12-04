package pov

//import "fmt"

type Graph struct{
	//maps => O(1) lookups
	conn map[string]map[string]bool //conn[from] => map[to]
	nodes map[string]bool
}

func New() *Graph{
	return &Graph{
		conn    : map[string]map[string]bool{},
		nodes   : map[string]bool{},
	}
}

func (g *Graph) AddNode(nodelabel string){
	g.conn[nodelabel]    = map[string]bool{}
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
	if oldRoot == newRoot { return g }

	//fmt.Println(g.ArcList())
	
	path := g.findPath(oldRoot, newRoot)
	//fmt.Println(oldRoot, newRoot, path)
	
	g.reversePath(path)
	//fmt.Println(g.ArcList())

	//in place !
	return g
}

func (g *Graph) findPath(from, to string) []string{

	if from == to {
		return []string{to}
	}

	for child, _ := range g.conn[from] {

		if path := g.findPath(child, to) ; path != nil {
			return append(path, from)
		}
	}

	return nil
}

func (g *Graph) reversePath(path []string){

	for c, p := 0, 1 ; p < len(path) ; c, p = c+1, p+1 {
		child, parent := path[c], path[p]

		g.reverseEdge(parent, child)
	}
}

//from -> to ==>> to -> from
func (g *Graph) reverseEdge(from, to string){
	g.AddArc(to, from)
	g.removeArc(from, to)
}

func (g *Graph) removeArc(from, to string){
	delete(g.conn[from], to)
}