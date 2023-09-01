package main

import (
	binarytree "binary-tree/binary-tree"
	"fmt"
)

func main() {
	bTree := binarytree.NewBinaryTree[int]()
	bTree.Add(5)
	bTree.Add(2)
	bTree.Add(3)
	bTree.Add(9)
	bTree.Add(4)
	bTree.Add(1)
	bTree.Add(6)
	bTree.Add(7)
	bTree.Add(8)
	bTree.Add(10)
	bTree.Print()

	fmt.Printf("contains %v? %t\n", 11, bTree.Contains(11))
	fmt.Printf("contains %v? %t\n", 3, bTree.Contains(3))
}
