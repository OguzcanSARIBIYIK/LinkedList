package main

import "fmt"

type LinkedList struct {
	Node   *Node
	Length uint
}

type Node struct {
	Data int
	Next *Node
}

func (l *LinkedList) add(n *Node) {

	if l.Node == nil {
		l.Node = n
		l.Node.Next = nil
	} else {
		n.Next = l.Node
		l.Node = n
	}
	l.Length++
}

func (l *LinkedList) removeWithValue(value int) {

	if l.Length == 0 {
		return
	}

	if l.Node.Data == value {
		l.Node = l.Node.Next
		l.Length--
		return
	}

	currentNode := l.Node
	len := l.Length

	for len > 0 {
		if currentNode.Next == nil {
			len--
			return
		}
		if currentNode.Next.Data == value {
			currentNode.Next = currentNode.Next.Next
			l.Length--
			break
		} else {
			currentNode = currentNode.Next
		}
		len--
	}

}

func (l LinkedList) printAllList() {

	if l.Length == 0 {
		fmt.Println("Empty List!")
	}

	len := l.Length
	toPrint := l.Node

	for len > 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.Next
		len--
	}

	fmt.Printf("\n")
}

func main() {
	linkedList := LinkedList{}
	node1 := &Node{Data: 12}
	node2 := &Node{Data: 23}
	node3 := &Node{Data: 24}
	node4 := &Node{Data: 26}
	linkedList.add(node1)
	linkedList.add(node2)
	linkedList.add(node3)
	linkedList.add(node4)

	linkedList.printAllList()
	linkedList.removeWithValue(23)
	linkedList.printAllList()
}
