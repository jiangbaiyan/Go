package class09

type Node struct {
	value interface{}
	next  *Node
}

// MidOrUpMidNode 输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func MidOrUpMidNode(head *Node) *Node {
	if head == nil || head.next == nil || head.next.next == nil {
		return head
	}
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
func MidOrUpMidPreNode(head *Node) {

}

// MidOrDownMidPreNode 输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func MidOrDownMidPreNode(head *Node) {

}
