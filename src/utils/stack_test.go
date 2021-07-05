package utils

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	t.Log(s.Pop()) // 1
	s.Push(2)
	s.Push(3)
	t.Log(s.IsEmpty()) // false
	t.Log(s.Pop())     // 3
	t.Log(s.Pop())     // 2
	t.Log(s.IsEmpty()) // true
}
