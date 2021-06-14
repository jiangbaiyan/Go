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

func TestRingArray(t *testing.T) {
	q := &MyQueue{}
	q.limit = 3
	q.arr = make([]interface{}, q.limit, q.limit)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for _, item := range q.arr {
		t.Log(item)
	}
	// 1, 2, 3
	ans, _ := q.Pop()
	// 1
	t.Log(ans)
	// 1, 2, 3
	for _, item := range q.arr {
		t.Log(item)
	}
	q.Push(4)
	// 4, 2, 3
	for _, item := range q.arr {
		t.Log(item)
	}
}
