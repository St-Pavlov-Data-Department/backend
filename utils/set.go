package utils

type EMPTY struct{}

var empty = EMPTY{}

type Set[T comparable] struct {
	m map[T]EMPTY
}

func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{}
	s.m = make(map[T]EMPTY)
	s.Add(items...)
	return s
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.m[item] = empty
	}
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Remove(item T) {
	delete(s.m, item)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}
