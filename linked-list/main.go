package main

import (
	"fmt"
	linkedlist "linked-list/linked-list"
)

func main() {
	linkedList := linkedlist.NewLinkedList[int]()
	for i := 0; i < 20; i += 2 {
		linkedList.Append(i)
	}
	linkedList.Print()
	fmt.Printf("contains %d? %t\n", 12, linkedList.Contains(12))
	fmt.Printf("contains %d? %t\n", 7, linkedList.Contains(7))
}
