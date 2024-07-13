package class12

import (
	"fmt"
	"testing"
)

// 辅助函数：打印链表
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func TestPartitionList(t *testing.T) {
	// 创建链表 1 -> 4 -> 3 -> 2 -> 5 -> 2 -> nil
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}

	fmt.Println("Original list:")
	printList(head)

	// 分隔链表
	x := 3
	newHead := partition(head, x)

	fmt.Println("Partitioned list:")
	printList(newHead)
}
