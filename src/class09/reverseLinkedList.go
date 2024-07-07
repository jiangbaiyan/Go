package class09

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	for head != nil {
		// 先记录下next，防止指针指向修改后丢失next
		next = head.Next
		// 主要操作，让当前节点的下一个节点指向pre
		head.Next = pre
		// 移动pre、head指针，让循环继续进行
		pre = head
		head = next
	}
	return pre
}

type DoubleListNode struct {
	Val  int
	Next *DoubleListNode
	Last *DoubleListNode
}

// 双链表翻转
func reverseDoubleList(head *DoubleListNode) *DoubleListNode {
	var pre *DoubleListNode
	var next *DoubleListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		head.Last = next
		pre = head
		head = next
	}
	return pre
}
