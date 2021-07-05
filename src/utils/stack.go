package utils

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
