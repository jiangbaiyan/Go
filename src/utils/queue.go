package utils

type Queue struct {
	data []interface{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() interface{} {
	ans := q.data[0]
	q.data = q.data[1:]
	return ans
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
