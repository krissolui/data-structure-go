package edgelist

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Graph struct {
	nodes [][]int
}

func NewGraph(bufSize int) Graph {
	return Graph{make([][]int, bufSize)}
}

func (g *Graph) addPair(root, pair int) bool {
	pairs := g.nodes[root]
	if slices.Contains(pairs, pair) {
		return false
	}
	g.nodes[root] = append(g.nodes[root], pair)
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
		if len(pairs) > 0 {
			fmt.Printf("\t%d --> %v\n", node, pairs)
		}
	}
}

func (g *Graph) Size() int {
	length := 0
	for _, pairs := range g.nodes {
		if len(pairs) > 0 {
			length++
		}
	}
	return length
}
