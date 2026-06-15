package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head * Node
}

func (list *SinglyLinkedList) addAtEnd(data int) {
	node := Node{data: data, next: nil}
	if list.head == nil {
		list.head = &node
		return
	}
	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &node
}

func (list *SinglyLinkedList) addAtBeginning(data int) {
	node := Node {data: data, next: nil}
	if list.head == nil {
		list.head = &node
		return
	}
	node.next = list.head
	list.head = &node
}

func (list *SinglyLinkedList) printList() {
	temp := list.head
	for temp != nil {
		fmt.Printf(" %d =>", temp.data)
		temp = temp.next
	}
	fmt.Println("")
}

func main() {
	list := &SinglyLinkedList{}
	for i := range 10 {
		list.addAtEnd((i+1) * 10)
	}
	list.printList()
	fmt.Println("--------------After Adding node at beginning----------------")
	list.addAtBeginning(444)
	list.printList()
}
