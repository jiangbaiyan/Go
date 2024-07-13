package class12

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你一个链表的头节点 head 和一个特定值 x
// 请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置
// 测试链接 : https://leetcode.cn/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	var smallHead, smallTail, bigHead, bigTail *ListNode
	for head != nil {
		if head.Val < x {
			if smallHead == nil {
				smallHead = head
				smallTail = smallHead
			} else {
				smallTail.Next = head
				smallTail = smallTail.Next
			}
		} else {
			if bigHead == nil {
				bigHead = head
				bigTail = bigHead
			} else {
				bigTail.Next = head
				bigTail = bigTail.Next
			}
		}
		head = head.Next
	}
	if bigTail != nil {
		bigTail.Next = nil // 确保大于或等于 x 的部分的尾节点指向 nil
	}
	if bigHead == nil {
		return smallHead
	}
	if smallHead == nil {
		return bigHead
	}
	smallTail.Next = bigHead
	return smallHead
}
