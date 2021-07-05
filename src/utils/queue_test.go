package utils

import "testing"

func TestQueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	t.Log(q.Dequeue()) // 1
	q.Enqueue(2)
	q.Enqueue(3)
	t.Log(q.IsEmpty()) // false
	t.Log(q.Dequeue()) // 2
	t.Log(q.Dequeue()) // 3
	t.Log(q.IsEmpty()) // true
}
