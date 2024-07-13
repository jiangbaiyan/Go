package class11

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给你两个 非空 的链表，表示两个非负的整数
// 它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头
// 测试链接：https://leetcode.cn/problems/add-two-numbers/
// l1 =
// [9,9,9,9,9,9,9]
// l2 =
// [9,9,9,9]
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 思路是搞一个新链表, 定义一个变量ans作为结果的链表头, 一个变量用于追踪答案链表
	var ans, cur *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		sum := l1Val + l2Val + carry
		curNodeVal := sum % 10
		carry = sum / 10
		if ans == nil {
			ans = &ListNode{Val: curNodeVal}
			cur = ans
		} else {
			cur.Next = &ListNode{Val: curNodeVal}
			cur = cur.Next
		}
		// 注意这里要判断,否则可能空指针panic
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// 最后的进位要挂到最后
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return ans
}
