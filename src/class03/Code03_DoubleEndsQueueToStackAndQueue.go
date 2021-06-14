package class03

// DoubleEndsQueue 链表实现双端队列
type DoubleEndsQueue struct {
	head *Node
	tail *Node
}

func (d *DoubleEndsQueue) EnqueueFromHead(v interface{}) {

}

func (d *DoubleEndsQueue) EnqueueFromTail(v interface{}) {

}

func (d *DoubleEndsQueue) DequeueFromHead() interface{} {
	return nil
}

func (d *DoubleEndsQueue) DequeueFromTail() interface{} {
	return nil
}

type MyStack struct {
	q *DoubleEndsQueue
}

type MyQueue struct {
	q *DoubleEndsQueue
}
