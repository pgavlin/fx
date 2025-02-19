package fx

import (
	"iter"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

func (s Set[T]) ToSlice() []T {
	return slices.Collect(s.Values())
}
