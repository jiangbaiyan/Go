package class10

type Node struct {
	value interface{}
	next  *Node
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
