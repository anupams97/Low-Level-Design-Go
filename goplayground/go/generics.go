package _go

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elems) == 0 {
		var zero T
		return zero, false
	}
	last := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return last, true
}
