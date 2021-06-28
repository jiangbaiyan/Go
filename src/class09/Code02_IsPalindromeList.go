package class09

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
	for n2 != nil {
		n3 := n2.next
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
	return res
}
