package node

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{data: data}
}

func (n *Node[T]) Append(data T) *Node[T] {
	if data == n.data {
		return n
	}

	if data < n.data {
		if n.left == nil {
			newNode := NewNode[T](data)
			n.left = newNode
			return newNode
		}

		return n.left.Append(data)
	}

	if n.right == nil {
		newNode := NewNode[T](data)
		n.right = newNode
		return newNode
	}

	return n.right.Append(data)
}

func (n *Node[T]) Data() T {
	return n.data
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) Print() {
	if n == nil {
		return
	}
	fmt.Printf("  %v\n", n.data)
	if n.left == nil && n.right == nil {
		return
	} else if n.left != nil && n.right != nil {
		fmt.Printf(" / \\\n")
		fmt.Printf("%v   %v\n", n.left.data, n.right.data)
	} else if n.left != nil {
		fmt.Printf(" /\n")
		fmt.Printf("%v\n", n.left.data)
	} else {
		fmt.Printf("   \\\n")
		fmt.Printf("   %v\n", n.right.data)
	}
}
