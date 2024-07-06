package class09

type Node struct {
	value int
	next  *Node
}

// MidOrUpMidNode 输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func MidOrUpMidNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
	// 核心: 通过改变slow和fast的初始位置来实现中点的指定返回
	slow := head
	fast := head
	// 奇数个, head.next == nil 停
	// 偶数个, head.next.next == nil 停
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

// MidOrDownMidNode 输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func MidOrDownMidNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	if head.next.next == nil {
		return head.next
	}
	slow := head.next
	fast := head.next
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

// MidOrUpMidPreNode 输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func MidOrUpMidPreNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	fast := head.next.next
	slow := head
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}

// MidOrDownMidPreNode 输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func MidOrDownMidPreNode(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}
	if head.next.next == nil {
		return head
	}
	fast := head.next
	slow := head
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow
}
