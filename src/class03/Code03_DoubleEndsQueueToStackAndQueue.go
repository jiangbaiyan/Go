package class03

// DoubleEndsQueue 链表实现双端队列
type DoubleEndsQueue struct {
	head *DoubleNode
	tail *DoubleNode
}

// EnqueueFromHead 入队，头插法
func (d *DoubleEndsQueue) EnqueueFromHead(v interface{}) {
	node := &DoubleNode{value: v}
	// 头结点为空
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		// 头结点不空，头部插入
		node.next = d.head
		d.head.last = node
		d.head = node
	}
}

// EnqueueFromHead 入队，尾插法
func (d *DoubleEndsQueue) EnqueueFromTail(v interface{}) {
	node := &DoubleNode{value: v}
	// 头结点为空
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		// 头结点不空，尾部插入
		d.tail.next = node
		node.last = d.tail
		d.tail = node
	}
}

// DequeueFromHead 从头部弹出
func (d *DoubleEndsQueue) DequeueFromHead() interface{} {
	if d.head == nil {
		return nil
	}
	cur := d.head
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
		d.head.last = nil
	}
	return cur.value
}

// DequeueFromTail 从尾部弹出
func (d *DoubleEndsQueue) DequeueFromTail() interface{} {
	if d.head == nil {
		return nil
	}
	cur := d.tail
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.last
		d.tail.next = nil
	}
	return cur.value
}

func (d *DoubleEndsQueue) IsEmpty() bool {
	return d.head == nil
}

type MyStack struct {
	q *DoubleEndsQueue
}

func (m *MyStack) Push(v interface{}) {
	m.q.EnqueueFromHead(v)
}

func (m *MyStack) Pop() interface{} {
	return m.q.DequeueFromHead()
}

func (m *MyStack) Peek() interface{} {
	return m.q.head.value
}

func (m *MyStack) IsEmpty() bool {
	return m.q.IsEmpty()
}

type MyQueue struct {
	q *DoubleEndsQueue
}

func (m *MyQueue) EnQueue(v interface{}) {
	m.q.EnqueueFromHead(v)
}

func (m *MyQueue) DeQueue() interface{} {
	return m.q.DequeueFromTail()
}
