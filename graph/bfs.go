package graph

import "fmt"

type BFS struct {
	marked map[int]bool
	edgeTo []int
}

func NewBFS(s int) *BFS {
	m := make(map[int]bool, s)
	e := make([]int, s+1)
	return &BFS {
		marked: m,
		edgeTo: e,
	}
}

func (b *BFS) Bfs(g Graph, s int) {
	queue := make(chan int, 100)
	b.marked[s] = true
	fmt.Println(g.Bag)
	queue <- s
	for len(queue) > 0 {
		v := <-queue
		for _, w := range g.Bag[v] {
			if ok := b.marked[w]; !ok {
				b.edgeTo[w] = v
				b.marked[w] = true
				queue <- w
			}
		}
	}
	b.pathTo(3)
}

func (b *BFS) pathTo(v int) {
	path := []int{}
	for x := v; x != 0; x = b.edgeTo[x] {
		path = append([]int{x},path...)
	}
	path = append([]int{0},path...)
	fmt.Println(path)
}
