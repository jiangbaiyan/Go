package class09

import "testing"

func TestMidOrUpMidNode(t *testing.T) {
	n1 := &Node{value: 1}
	t.Log(MidOrUpMidNode(n1).value) // 1
	n2 := &Node{value: 2}
	n1.next = n2
	t.Log(MidOrUpMidNode(n1).value) // 1
	n3 := &Node{value: 3}
	n2.next = n3
	n4 := &Node{value: 4}
	n3.next = n4
	n5 := &Node{value: 5}
	n4.next = n5
	t.Log(MidOrUpMidNode(n1).value) // 3
	n6 := &Node{value: 6}
	n5.next = n6
	t.Log(MidOrUpMidNode(n1).value) // 3
}

func TestMidOrDownMidNode(t *testing.T) {
	n1 := &Node{value: 1}
	t.Log(MidOrDownMidNode(n1).value) // 1
	n2 := &Node{value: 2}
	n1.next = n2
	t.Log(MidOrDownMidNode(n1).value) // 2
	n3 := &Node{value: 3}
	n2.next = n3
	n4 := &Node{value: 4}
	n3.next = n4
	n5 := &Node{value: 5}
	n4.next = n5
	t.Log(MidOrDownMidNode(n1).value) // 3
	n6 := &Node{value: 6}
	n5.next = n6
	t.Log(MidOrDownMidNode(n1).value) // 4
}

func TestMidOrUpMidPreNode(t *testing.T) {
	n1 := &Node{value: 1}
	t.Log(MidOrUpMidPreNode(n1)) // nil
	n2 := &Node{value: 2}
	n1.next = n2
	t.Log(MidOrUpMidPreNode(n1)) // nil
	n3 := &Node{value: 3}
	n2.next = n3
	n4 := &Node{value: 4}
	n3.next = n4
	n5 := &Node{value: 5}
	n4.next = n5
	t.Log(MidOrUpMidPreNode(n1).value) // 2
	n6 := &Node{value: 6}
	n5.next = n6
	t.Log(MidOrUpMidPreNode(n1).value) // 2
}

func TestMidOrDownMidPreNode(t *testing.T) {
	n1 := &Node{value: 1}
	t.Log(MidOrDownMidPreNode(n1)) // nil
	n2 := &Node{value: 2}
	n1.next = n2
	t.Log(MidOrDownMidPreNode(n1).value) // 1
	n3 := &Node{value: 3}
	n2.next = n3
	n4 := &Node{value: 4}
	n3.next = n4
	n5 := &Node{value: 5}
	n4.next = n5
	t.Log(MidOrDownMidPreNode(n1).value) // 2
	n6 := &Node{value: 6}
	n5.next = n6
	t.Log(MidOrDownMidPreNode(n1).value) // 3
}

func TestIsPalindromeList(t *testing.T) {
	n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	n1.next = n2
	n3 := &Node{value: 3}
	n2.next = n3
	n4 := &Node{value: 3}
	n3.next = n4
	n5 := &Node{value: 2}
	n4.next = n5
	n6 := &Node{value: 1}
	n5.next = n6
	t.Log(IsPalindromeList(n1))
}
