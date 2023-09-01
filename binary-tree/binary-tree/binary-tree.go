package binarytree

import (
	"binary-tree/node"

	"golang.org/x/exp/constraints"
)

type BinaryTree[T constraints.Ordered] struct {
	root *node.Node[T]
}

func NewBinaryTree[T constraints.Ordered]() BinaryTree[T] {
	return BinaryTree[T]{}
}

func (bt *BinaryTree[T]) Add(data T) {
	if bt.root == nil {
		bt.root = node.NewNode[T](data)
		return
	}

	if data == bt.root.Data() {
		return
	}

	bt.root.Append(data)
}

func printAll[T constraints.Ordered](root *node.Node[T]) {
	if root == nil {
		return
	}
	root.Print()
	printAll(root.Left())
	printAll(root.Right())
}

func (bt *BinaryTree[T]) Print() {
	// bt.root.PrintAll()
	printAll[T](bt.root)
}

func searchAll[T constraints.Ordered](root *node.Node[T], data T) bool {
	if root.Data() == data {
		return true
	}
	if root.Left() != nil && root.Right() != nil {
		return searchAll[T](root.Left(), data) || searchAll[T](root.Right(), data)
	}
	if root.Left() != nil {
		return searchAll[T](root.Left(), data)
	}
	if root.Right() != nil {
		return searchAll[T](root.Right(), data)
	}
	return false
}

func (bt *BinaryTree[T]) Contains(data T) bool {
	return searchAll[T](bt.root, data)
}
