package sets

import (
	"iter"
	"maps"
	"slices"

	"github.com/pgavlin/fx/v2"
)

// Values returns a sequence of each value in the input set. The ordering of elements is undefined.
func Values[T ~map[E]struct{}, E comparable](s T) iter.Seq[E] {
	return maps.Keys(s)
}

// Intersection returns a new set that is the intersection of the input sets.
func Intersection[T ~map[E]struct{}, U ~map[E]struct{}, E comparable](s1 T, s2 U) fx.Set[E] {
	return Collect(Filter(s1, func(e E) bool { return fx.Set[E](s2).Has(e) }))
}

// IntersectMany returns a new set that is the intersection of the input sets.
func IntersectMany[T ~map[E]struct{}, E comparable](sets iter.Seq[T]) fx.Set[E] {
	first, ok := fx.First(sets)
	if !ok {
		return nil
	}
	return fx.Reduce(sets, Collect(Values(first)), func(acc fx.Set[E], s T) fx.Set[E] {
		acc.Intersect(fx.Set[E](s))
		return acc
	})
}

// IntersectAll returns a new set that is the intersection of the input sets.
func IntersectAll[T ~map[E]struct{}, E comparable](sets ...T) fx.Set[E] {
	return IntersectMany(slices.Values(sets))
}

// Union returns a new set that is the union of the input sets.
func Union[T ~map[E]struct{}, U ~map[E]struct{}, E comparable](s1 T, s2 U) fx.Set[E] {
	return Collect(fx.Concat(Values(s1), Values(s2)))
}

// UnionMany returns a new set that is the union of the input sets.
func UnionMany[T ~map[E]struct{}, E comparable](sets iter.Seq[T]) fx.Set[E] {
	return Collect(fx.ConcatMany(fx.Map(sets, func(s T) iter.Seq[E] { return Values(s) })))
}

// UnionAll returns a new set that is the union of the input sets.
func UnionAll[T ~map[E]struct{}, E comparable](sets ...T) fx.Set[E] {
	return UnionMany(slices.Values(sets))
}

// Collect collects the values in the input sequence into a set.
func Collect[T comparable](it iter.Seq[T]) fx.Set[T] {
	var s fx.Set[T]
	for t := range it {
		if s == nil {
			s = make(fx.Set[T])
		}
		s.Add(t)
	}
	return s
}

// TryCollect attempts to collect the values in the input sequence into a set. If any pair in the input contains a
// non-nil error, TryCollect halts and returns the values collected up to that point (excluding the value returned
// with the error) and the error.
func TryCollect[T comparable](it iter.Seq2[T, error]) (fx.Set[T], error) {
	var s fx.Set[T]
	for t, err := range it {
		if err != nil {
			return s, err
		}
		if s == nil {
			s = make(fx.Set[T])
		}
		s.Add(t)
	}
	return s, nil
}
