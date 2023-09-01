package node

type Node[T any] struct {
	data T
	next *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data, nil}
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Append(data T) *Node[T] {
	node := NewNode[T](data)
	n.next = node
	return node
}

func (n *Node[T]) tail() *Node[T] {
	currentNode := n
	for {
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
	}
	return currentNode
}

func (n *Node[T]) Concat(next *Node[T]) *Node[T] {
	n.next = next
	return next.tail()
}

func (n *Node[T]) Data() T {
	return n.data
}
