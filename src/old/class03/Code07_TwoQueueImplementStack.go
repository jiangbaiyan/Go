package class03

// TwoQueueStack 两个队列实现栈
type TwoQueueStack struct {
	queue *MyQueue
	help  *MyQueue
}

func (t *TwoQueueStack) Push(v interface{}) {
	t.queue.EnQueue(v)
}

func (t *TwoQueueStack) Pop() interface{} {
	// 将主队列依次出队，放到help队列里去，直到剩下最后一个元素，那么最后一个元素就是要弹出的值
	for t.queue.GetSize() > 1 {
		t.help.EnQueue(t.queue.DeQueue())
	}
	ans := t.queue.DeQueue()
	// 现在元素都在help队列里，所以交换指针引用
	t.queue, t.help = t.help, t.queue
	return ans
}
