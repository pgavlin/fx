package fx

import "iter"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) In(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Iter() iter.Seq[T] {
	return IterSet(s)
}

func (s Set[T]) ToSlice() []T {
	return ToSlice(IterSet(s))
}
