package linked_list

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(data int) {
	node := &Node{Data: data}

	if l.Head == nil {
		l.Head = node
		return
	}

	currentNode := l.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = node
}

func (l *LinkedList) InsertFirst(data int) {
	node := &Node{Data: data}

	if l.Head == nil {
		l.Head = node
		return
	}

	node.Next = l.Head
	l.Head = node
}

func (l *LinkedList) InsertAt(index, data int) {
	newNode := &Node{Data: data}

	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	currentNode := l.Head
	counter := 0

	for currentNode != nil && counter < index-1 {
		currentNode = currentNode.Next
		counter++
	}

	if currentNode != nil {
		newNode.Next = currentNode.Next
		currentNode.Next = newNode
	}

}

func (l *LinkedList) DeleteFirst() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

func (l *LinkedList) DeleteLast() {
	if l.Head != nil {

		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode = nil
	}
}

func (l *LinkedList) DeleteAt(index int) {
	if l.Head == nil {
		return
	}

	if index == 0 {
		l.Head = l.Head.Next
		return
	}

	counter := 0
	currentNode := l.Head

	for counter < index-1 && currentNode.Next != nil {
		currentNode = currentNode.Next
		counter++
	}

	if currentNode.Next != nil {
		currentNode.Next = currentNode.Next.Next
	}
}

// Reverse will reverse the element from the list
// Key point of the algorithm is we need to Reverse
// the pointer of each node from next node to prev node
// ex: a -> b -> c will be reversed to a <- b <- c
func (l *LinkedList) Reverse() {
	var prevNode *Node
	currentNode := l.Head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	l.Head = prevNode
}

func (l *LinkedList) Print() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Print(currentNode.Data, " ")
		currentNode = currentNode.Next
	}
}
