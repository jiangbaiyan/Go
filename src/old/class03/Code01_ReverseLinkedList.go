package class03

type Node struct {
	value interface{}
	next  *Node
}

type DoubleNode struct {
	value interface{}
	next  *DoubleNode
	last  *DoubleNode
}

// ReverseLinkedList 单链表反转
// 采用pre head next三指针法
func ReverseLinkedList(head *Node) *Node {
	var pre *Node = nil
	var next *Node = nil
	for head != nil {
		next = head.next
		head.next = pre
		pre = head
		head = next
	}
	return pre
}

// ReverseDoubleLinkedList 双链表反转
func ReverseDoubleLinkedList(head *DoubleNode) *DoubleNode {
	var pre *DoubleNode = nil
	var next *DoubleNode = nil
	for head != nil {
		next = head.next
		head.next = pre
		head.last = next // 只比单链表反转多调整一个last指针
		pre = head
		head = next
	}
	return pre
}
