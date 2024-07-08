package class10

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	var head *ListNode
	// cur1指向初始为head的next，cur2指向另一个链表头
	var cur1, cur2 *ListNode
	if h1.Val < h2.Val {
		head = h1
		cur1 = head.Next
		cur2 = h2
	} else {
		head = h2
		cur1 = head.Next
		cur2 = h1
	}
	// 这样做会少了一个追踪当前合并链表的变量
	//for cur1 != nil && cur2 != nil {
	//	if cur1.Val <= cur2.Val {
	//		next := cur1.Next
	//		cur1.Next = cur2
	//		cur1 = next
	//	} else {
	//		next := cur2.Next
	//		cur2.Next = cur1
	//		cur2 = next
	//	}
	//}
	var mergedLinkedNode = head
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			mergedLinkedNode.Next = cur1
			cur1 = cur1.Next
			mergedLinkedNode = mergedLinkedNode.Next
		} else {
			mergedLinkedNode.Next = cur2
			cur2 = cur2.Next
			mergedLinkedNode = mergedLinkedNode.Next
		}
	}
	if cur1 == nil {
		mergedLinkedNode.Next = cur2
	} else {
		mergedLinkedNode.Next = cur1
	}
	return head
}
