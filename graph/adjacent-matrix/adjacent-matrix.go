package adjacentmatrix

import (
	"fmt"
)

type Graph struct {
	nodes [][]int
}

func (g *Graph) initialize() {
	bufSize := len(g.nodes)
	for i := 0; i < bufSize; i++ {
		g.nodes[i] = make([]int, bufSize)
	}
}

func NewGraph(bufSize int) Graph {
	newGraph := Graph{make([][]int, bufSize)}
	newGraph.initialize()
	return newGraph
}

func (g *Graph) addPair(root, pair int) bool {
	pairs := g.nodes[root]
	if paired := pairs[pair]; paired > 0 {
		return false
	}
	g.nodes[root][pair] = 1
	return true
}

func (g *Graph) PairNodes(node1, node2 int) bool {
	if ok := g.addPair(node1, node2); !ok {
		return false
	}
	g.addPair(node2, node1)
	return true
}

func repeatString(str string, count int) string {
	s := ""
	for i := 0; i < count; i++ {
		s += str
	}
	return s
}

func getPairs(entries []int) []int {
	pairs := []int{}

	for i, paired := range entries {
		if paired > 0 {
			pairs = append(pairs, i)
		}
	}

	return pairs
}

func (g *Graph) Print() {
	for node, entries := range g.nodes {
		pairs := getPairs(entries)
		if len(pairs) == 0 {
			continue
		}

		fmt.Printf("\t%d --> %v\n", node, pairs)
	}
}

func (g *Graph) PrintMatrix() {
	rangeSlice := make([]int, len(g.nodes))
	for i, _ := range rangeSlice {
		rangeSlice[i] = i
	}
	fmt.Printf("\t  %v\n", rangeSlice)
	fmt.Printf("\t  %s\n", repeatString("-", len(g.nodes)*2+1))

	for node, entries := range g.nodes {
		if len(entries) > 0 {
			fmt.Printf("\t%d %v\n", node, entries)
		}
	}
}

func (g *Graph) Size() int {
	length := 0
	for _, entries := range g.nodes {
		pairs := getPairs(entries)
		if len(pairs) > 0 {
			length++
		}
	}
	return length
}
