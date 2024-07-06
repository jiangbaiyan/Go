package class09

// SmallerEqualBigger 将单向链表按某值划分成左边小、中间相等、右边大的形式
func SmallerEqualBigger(head *Node, pivot int) *Node {
	var sH, sT, eH, eT, bH, bT *Node
	for head != nil {
		next := head.next
		head.next = nil
		if head.value < pivot {
			if sH == nil {
				sH = head
				sT = head
			} else {
				sT = &Node{}
				sT.next = head
				sT = head
			}
		}
		if head.value == pivot {
			if eH == nil {
				eH = head
				eT = head
			} else {
				eT = &Node{}
				eT.next = head
				eT = head
			}
		}
		if head.value > pivot {
			if bH == nil {
				bH = head
				bT = head
			} else {
				bT = &Node{}
				bT.next = head
				bT = head
			}
		}
		head = next
	}
	if sT != nil {
		sT.next = eH
		if eT == nil {
			eT = sT
		}
	}
	if eT != nil {
		eT.next = bH
	}
	if sH != nil {
		return sH
	} else {
		if eH != nil {
			return eH
		}
		return bH
	}
}
