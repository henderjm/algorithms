package graph

type Graph struct {
	V   int
	E   int
	Bag map[int][]int
}

func (g *Graph) New(v int) {
	g.V = v
	g.Bag = make(map[int][]int, g.V)
	g.addEdge(0, 2)
	g.addEdge(0, 1)
	g.addEdge(0, 5)
	g.addEdge(1, 2)
	g.addEdge(3, 5)
	g.addEdge(2, 4)
	g.addEdge(3, 4)
	g.addEdge(2, 3)

	//6
	//8
	//0 5
	//2 4
	//2 3
	//1 2
	//0 1
	//3 4
	//3 5
	//0 2
}

func (g *Graph) addEdge(v, w int) {
	g.Bag[v] = append(g.Bag[v], w)
	g.Bag[w] = append(g.Bag[w], v)
	g.E += 1
}
