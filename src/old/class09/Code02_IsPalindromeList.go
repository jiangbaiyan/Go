package class09

// IsPalindromeList 回文链表
func IsPalindromeList(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}
	// 1. 找到链表中点, 最终slow来到mid位置, fast来到end位置
	fast := head
	slow := head
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	// 2. 反转后部分的链表
	// n2 来到右边第一个节点
	n1 := slow
	n2 := slow.next
	slow.next = nil
	var n3 *Node
	for n2 != nil {
		n3 = n2.next
		n2.next = n1
		n1 = n2
		n2 = n3
	}
	// 此时n1来到结尾位置, n2 n3都为nil
	// 3. 判断回文
	t := n1
	h := head
	res := true
	for h != nil && t != nil {
		if h.value != t.value {
			res = false
			break
		}
		h = h.next
		t = t.next
	}
	// 4. 这里不能直接return, 需要将原链表调整回去
	// 翻转后面这部分链表, n3 = pre, n1 = cur , n2 = next
	for n1 != nil {
		n2 = n1.next
		n1.next = n3
		n3 = n1
		n1 = n2
	}
	return res
}
