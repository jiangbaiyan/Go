package class03

// TwoStackQueue 两个栈实现队列
type TwoStackQueue struct {
	stackPush *MyStack
	stackPop  *MyStack
}

func (t *TwoStackQueue) EnQueue(v interface{}) {
	t.stackPush.Push(v)
	t.PushToPop()
}

func (t *TwoStackQueue) Dequeue() interface{} {
	if t.stackPush.IsEmpty() && t.stackPop.IsEmpty() {
		panic("队列为空")
	}
	t.PushToPop()
	return t.stackPop.Pop()
}

// PushToPop 两个栈互相倒，需要满足两个条件：
// 1. push栈要一次性倒完
// 2. pop栈不能有值
func (t *TwoStackQueue) PushToPop() {
	if t.stackPop.IsEmpty() {
		for !t.stackPush.IsEmpty() {
			t.stackPop.Push(t.stackPush.Pop())
		}
	}
}
