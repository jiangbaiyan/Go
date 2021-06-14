package class03

import (
	"log"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	node1 := &Node{value: 1, next: nil}
	node2 := &Node{next: node1, value: 2}
	node3 := &Node{next: node2, value: 3}
	head := node3
	resHead := ReverseLinkedList(head)
	for resHead != nil {
		log.Print(resHead.value)
		resHead = resHead.next
	}
}

func TestReverseDoubleLinkedList(t *testing.T) {
	node1 := &DoubleNode{value: 1, next: nil}
	node2 := &DoubleNode{next: node1, value: 2}
	node3 := &DoubleNode{next: node2, value: 3}
	node1.last = node2
	node2.last = node3
	node3.last = nil
	head := node3
	resHead := ReverseDoubleLinkedList(head)
	for resHead != nil {
		log.Print(resHead.value)
		resHead = resHead.next
	}
}

func TestDeleteGivenValue(t *testing.T) {
	node1 := &Node{value: 1, next: nil}
	node2 := &Node{next: node1, value: 2}
	node3 := &Node{next: node2, value: 3}
	head := node3
	resHead := DeleteGivenValue(head, 3)
	for resHead != nil {
		log.Print(resHead.value)
		resHead = resHead.next
	}
}
