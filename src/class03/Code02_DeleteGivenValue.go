package class03

// DeleteGivenValue 删除链表中指定值的元素，返回头结点
func DeleteGivenValue(head *Node, num int) *Node {
	for head.value == num {
		head = head.next
	}
	// head来到第一个不是num的位置，head不再往后移动，需要cur来进行接下来的遍历
	var pre *Node = head
	var cur *Node = head
	for cur != nil {
		if cur.value == num {
			pre.next = cur.next
		} else {
			pre = cur
		}
		cur = cur.next
	}
	return head
}
