package class09

type NodeRandom struct {
	next   *NodeRandom
	random *NodeRandom
	value  interface{}
}

// CopyListWithRand rand指针是单链表节点结构中新增的指针，rand可能指向链表中的任意一个节点，也可能指向null。
// 给定一个由Node节点类型组成的无环单链表的头节点 head，请实现一个函数完成这个链表的复制，并返回复制的新链表的头节点。
//【要求】
// 时间复杂度O(N)，额外空间复杂度O(1)
// 容器方法
func CopyListWithRand(head *NodeRandom) *NodeRandom {
	m := make(map[*NodeRandom]*NodeRandom)
	cur := head
	for cur != nil {
		// map存储新结点
		m[cur] = &NodeRandom{value: cur.value}
		cur = cur.next
	}
	cur = head
	for cur != nil {
		m[cur].next = m[cur.next]
		m[cur].random = m[cur.random]
		cur = cur.next
	}
	return m[head]
}
