package main

import (
	"fmt"
	adjacentlist "graph/adjacent-list"
	adjacentmatrix "graph/adjacent-matrix"
	edgelist "graph/edge-list"
)

func main() {
	fmt.Println("---------- Edge List ----------")
	graph1 := edgelist.NewGraph(5)
	fmt.Printf("Initial size: %d\n", graph1.Size())
	graph1.Print()

	graph1.PairNodes(0, 2)
	graph1.PairNodes(0, 3)
	graph1.PairNodes(1, 3)
	graph1.PairNodes(1, 4)
	graph1.PairNodes(2, 4)
	graph1.PairNodes(3, 4)
	fmt.Printf("Updated size: %d\n", graph1.Size())
	graph1.Print()
	fmt.Println()

	fmt.Println("---------- Adjacent List ----------")
	graph2 := adjacentlist.NewGraph(5)
	fmt.Printf("Initial size: %d\n", graph2.Size())
	graph2.Print()

	graph2.PairNodes(0, 2)
	graph2.PairNodes(0, 3)
	graph2.PairNodes(1, 3)
	graph2.PairNodes(1, 4)
	graph2.PairNodes(2, 4)
	graph2.PairNodes(3, 4)
	fmt.Printf("Updated size: %d\n", graph2.Size())
	graph2.Print()
	fmt.Println()

	fmt.Println("---------- Adjacent Matrix ----------")
	graph3 := adjacentmatrix.NewGraph(5)
	fmt.Printf("Initial size: %d\n", graph3.Size())
	graph3.Print()
	fmt.Println()
	graph3.PrintMatrix()

	graph3.PairNodes(0, 2)
	graph3.PairNodes(0, 3)
	graph3.PairNodes(1, 3)
	graph3.PairNodes(1, 4)
	graph3.PairNodes(2, 4)
	graph3.PairNodes(3, 4)
	fmt.Printf("Updated size: %d\n", graph3.Size())
	graph3.Print()
	fmt.Println()
	graph3.PrintMatrix()
}
