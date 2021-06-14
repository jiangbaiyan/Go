package class03

// MinStack 实现具有getMin功能的栈
type MinStack struct {
	dataStack *MyStack
	minStack  *MyStack
}

// Push 如果当前数比最小栈小，向最小栈压入此数，否则重复压入最小栈栈顶的数，做到记录此时此刻的最小值
func (m *MinStack) Push(v int) {
	m.dataStack.Push(v)
	if m.minStack.IsEmpty() {
		m.minStack.Push(v)
	} else {
		newMin := m.GetMin().(int)
		if v > newMin {
			// 如果当前数较大，不更新最小值，重复压入最小栈顶元素
			m.minStack.Push(newMin)
		} else {
			// 当前数更小，直接向最小栈压入当前数
			m.minStack.Push(v)
		}
	}
}

func (m *MinStack) Pop() interface{} {
	if m.dataStack.IsEmpty() {
		panic("栈已空")
	}
	// 两个栈同时弹出
	m.minStack.Pop()
	return m.dataStack.Pop()
}

func (m *MinStack) GetMin() interface{} {
	if m.minStack.IsEmpty() {
		panic("栈已空")
	}
	return m.minStack.Peek()
}
