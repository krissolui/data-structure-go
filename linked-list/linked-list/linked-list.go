package linkedlist

import (
	"fmt"
	"linked-list/node"
)

type LinkedList[T comparable] struct {
	head *node.Node[T]
	tail *node.Node[T]
}

func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{}
}

func (l *LinkedList[T]) Append(data T) {
	var newNode *node.Node[T]
	if l.head == nil {
		newNode = node.NewNode[T](data)
		l.head = newNode
	} else {
		newNode = l.tail.Append(data)
	}
	l.tail = newNode
}

func searchAll[T comparable](n *node.Node[T], data T) bool {
	if n == nil {
		return false
	}
	if n.Data() == data {
		return true
	}
	return searchAll[T](n.Next(), data)
}

func (l *LinkedList[T]) Contains(data T) bool {
	return searchAll(l.head, data)
}

func (l *LinkedList[T]) Print() {
	currentNode := l.head
	for {
		if currentNode == nil {
			break
		}
		fmt.Println(currentNode.Data())
		currentNode = currentNode.Next()
	}
}
