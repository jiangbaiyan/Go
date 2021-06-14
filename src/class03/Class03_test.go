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

func TestMyQueue(t *testing.T) {
	q := &MyQueue{q: &DoubleEndsQueue{}}
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
	t.Log(q.DeQueue())
}

func TestMyStack(t *testing.T) {
	q := &MyStack{q: &DoubleEndsQueue{}}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
}

func TestRingArray(t *testing.T) {
	q := &MyRingQueue{}
	q.limit = 3
	q.arr = make([]interface{}, q.limit, q.limit)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for _, item := range q.arr {
		t.Log(item)
	}
	// 1, 2, 3
	ans := q.Pop()
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

func TestGetMin(t *testing.T) {
	ms := &MinStack{minStack: &MyStack{q: &DoubleEndsQueue{}}, dataStack: &MyStack{q: &DoubleEndsQueue{}}}
	ms.Push(2)
	t.Log(ms.GetMin()) // 2
	ms.Push(2)
	t.Log(ms.GetMin()) // 2
	ms.Push(3)
	t.Log(ms.GetMin()) // 2
	ms.Push(1)
	t.Log(ms.GetMin()) // 1
}

func TestTwoStackQueue(t *testing.T) {
	q := &TwoStackQueue{stackPush: &MyStack{q: &DoubleEndsQueue{}}, stackPop: &MyStack{q: &DoubleEndsQueue{}}}
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
	t.Log(q.Dequeue())
}

func TestTwoQueueStack(t *testing.T) {
	s := &TwoQueueStack{queue: &MyQueue{q: &DoubleEndsQueue{}}, help: &MyQueue{q: &DoubleEndsQueue{}}}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}

func TestGetMax(t *testing.T) {
	arr := []int{1, 2, 3, 4, 2, 3, 5, 6, 7, 2, 1}
	t.Log(GetMax(arr))
}
