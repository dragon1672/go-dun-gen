package memai

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) > 0 {
		return s.data[len(s.data)-1], true
	}
	var empty T
	return empty, false
}

func (s *Stack[T]) Pop() (T, bool) {
	var empty T
	if len(s.data) == 0 {
		return empty, false
	}
	i := len(s.data) - 1
	ret := s.data[i]
	s.data[i] = empty
	s.data = s.data[:i]
	return ret, true
}
