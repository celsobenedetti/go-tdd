package generics

import "sync"

// thread safe genetic comparable stack
type Stack[T any] struct {
	data  []T
	mutex sync.Mutex
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.Len() == 0 {
		var zero T
		return zero, false
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	idx := s.Len() - 1
	result := s.data[idx]
	s.data = s.data[:idx]

	return result, true
}

func (s *Stack[T]) Push(value T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data = append(s.data, value)
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}
