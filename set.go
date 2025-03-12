package fx

import (
	"iter"
	"maps"
	"slices"
)

// A Set represents a set of comparable values.
type Set[T comparable] map[T]struct{}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Add adds a value to the set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Remove removes a value from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Has returns true if the set contains the given value.
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// Copy returns a shallow copy of S.
func (s Set[T]) Copy() Set[T] {
	other := make(Set[T], s.Len())
	for k := range s {
		other.Add(k)
	}
	return other
}

// Intersect sets the contents of s to the intersection of s and other.
func (s Set[T]) Intersect(other Set[T]) {
	for k := range s {
		if _, ok := other[k]; !ok {
			delete(s, k)
		}
	}
}

// Union sets the contents of s to the union of s and other.
func (s Set[T]) Union(other Set[T]) {
	for k := range other {
		s.Add(k)
	}
}

// Values returns a sequence of each value in the set. The ordering of elements is undefined.
func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

// ToSlice returns a slice that contains the values in the set. The ordering of elements is undefined.
func (s Set[T]) ToSlice() []T {
	return slices.Collect(s.Values())
}
