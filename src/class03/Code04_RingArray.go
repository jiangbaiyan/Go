package class03

import "errors"

type MyQueue struct {
	pushi int
	popi  int
	size  int
	limit int
	arr   []interface{}
}

func (m *MyQueue) Push(value interface{}) error {
	if m.size == m.limit {
		return errors.New("队列已满")
	}
	m.arr[m.pushi] = value
	m.size++
	m.pushi = m.nextIndex(m.pushi)
	return nil
}

func (m *MyQueue) Pop() (interface{}, error) {
	if m.size == 0 {
		return nil, errors.New("队列为空")
	}
	ans := m.arr[m.popi]
	m.size--
	m.popi = m.nextIndex(m.popi)
	return ans, nil
}

func (m *MyQueue) nextIndex(i int) int {
	if i+1 >= m.limit {
		return 0
	}
	return i + 1
}
