package main

import "github.com/henderjm/algorithms/graph"

func main() {
	g := graph.Graph{}
	g.New(5)
	b := graph.NewBFS(g.V)
	b.Bfs(g,0)

}
