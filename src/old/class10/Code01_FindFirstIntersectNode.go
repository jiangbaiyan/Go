package class10

type Node struct {
	value interface{}
	next  *Node
}

// FindFirstIntersectNode
// 给定两个可能有环也可能无环的单链表，头节点head1和head2。请实现一个函数，如果两个链表相交，请返回相交的 第一个节点。如果不相交，返回null
// 【要求】
// 如果两个链表长度之和为N，时间复杂度请达到O(N)，额外空间复杂度 请达到O(1)。
func FindFirstIntersectNode(head1 *Node, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	loop1 := GetLoopNode(head1)
	loop2 := GetLoopNode(head2)
	// 两个都无环
	if loop1 == nil && loop2 == nil {
		return NoLoop(head1, head2)
	}
	// 两个都有环
	if loop1 != nil && loop2 != nil {
		return BothLoop(head1, loop1, head2, loop2)
	}
	// 一个有环一个无环, 不可能
	return nil
}

// GetLoopNode 找到链表第一个入环节点, 如果无环, 返回null
func GetLoopNode(head *Node) *Node {
	// 1个或2个节点, 不可能形成环
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head
	fast := head
	for slow != fast {
		// 如果fast来到了空, 那么肯定无环
		if fast.next == nil || fast.next.next == nil {
			return nil
		}
		fast = fast.next.next
		slow = slow.next
	}
	// 代码走到这, 两个指针相遇, 让fast指针回到原点, 步长改为1
	fast = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}

// NoLoop 找到链表第一个相交节点(无环)
func NoLoop(head1 *Node, head2 *Node) *Node {
	if head1 == nil || head2 == nil {
		return nil
	}
	cur1 := head1
	cur2 := head2
	len1 := 0
	len2 := 0
	for cur1.next != nil {
		len1++
		cur1 = cur1.next
	}
	for cur2.next != nil {
		len2++
		cur2 = cur2.next
	}
	// 长的那个链表先走差值步
	cur1 = head1
	cur2 = head2
	if len1-len2 > 0 {
		n := len1 - len2
		for n != 0 {
			cur1 = cur1.next
			n--
		}
	} else {
		n := len2 - len1
		for n != 0 {
			cur2 = cur2.next
			n--
		}
	}
	for cur1 != cur2 {
		cur1 = cur1.next
		cur2 = cur2.next
	}
	return cur1
}

// BothLoop 两个链表都有环
func BothLoop(head1 *Node, head2 *Node, loop1 *Node, loop2 *Node) *Node {
	// 1. 入环节点在环外
	if loop1 == loop2 {
		cur1 := head1
		cur2 := head2
		len1 := 0
		len2 := 0
		for cur1 != loop1 {
			len1++
			cur1 = cur1.next
		}
		for cur2 != loop1 {
			len2++
			cur2 = cur2.next
		}
		cur1 = head1
		cur2 = head2
		if len1-len2 > 0 {
			n := len1 - len2
			for n != 0 {
				cur1 = cur1.next
				n--
			}
		} else {
			n := len2 - len1
			for n != 0 {
				cur2 = cur2.next
				n--
			}
		}
		return cur1
	} else {
		// 2. 入环节点在环内
		cur1 := loop1
		// 转一圈, 看回来之前是否嫩关于到loop2, 如果能, 则返回loop1/loop2
		for cur1 != loop1 {
			if cur1 == loop2 {
				return loop1
			}
			cur1 = cur1.next
		}
		// 若转了一圈都没遇到, 则不相交
		return nil
	}
}
