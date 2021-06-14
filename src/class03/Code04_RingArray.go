package class03

type MyQueue struct {
	pushi int
	popi  int
	size  int
	limit int
	arr   []interface{}
}

func (m *MyQueue) Push(value interface{}) {
	if m.size == m.limit {
		panic("队列已满")
	}
	m.arr[m.pushi] = value
	m.size++
	m.pushi = m.nextIndex(m.pushi)
	return
}

func (m *MyQueue) Pop() interface{} {
	if m.size == 0 {
		panic("队列为空")
	}
	ans := m.arr[m.popi]
	m.size--
	m.popi = m.nextIndex(m.popi)
	return ans
}

func (m *MyQueue) nextIndex(i int) int {
	if i+1 >= m.limit {
		return 0
	}
	return i + 1
}
