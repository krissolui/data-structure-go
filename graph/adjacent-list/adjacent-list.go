package adjacentlist

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Graph struct {
	nodes map[int][]int
}

func NewGraph(bufSize int) Graph {
	return Graph{make(map[int][]int, bufSize)}
}

func (g *Graph) addPair(root, pair int) bool {
	if pairs, ok := g.nodes[root]; ok {
		if slices.Contains(pairs, pair) {
			return false
		}
		g.nodes[root] = append(g.nodes[root], pair)
		return true
	}

	g.nodes[root] = []int{pair}
	return true
}

func (g *Graph) PairNodes(node1, node2 int) bool {
	if ok := g.addPair(node1, node2); !ok {
		return false
	}
	g.addPair(node2, node1)
	return true
}

func (g *Graph) Print() {
	for node, pairs := range g.nodes {
		fmt.Printf("\t%d --> %v\n", node, pairs)
	}
}

func (g *Graph) Size() int {
	return len(g.nodes)
}
