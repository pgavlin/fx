package sets

import (
	"iter"
	"maps"

	"github.com/pgavlin/fx/v2"
)

// All returns true if pred returns true for every element of the input set.
func All[T ~map[E]struct{}, E comparable](s T, pred func(v E) bool) bool {
	return fx.All(maps.Keys(s), pred)
}

// Any returns true if pred returns true for any element of the input set.
func Any[T ~map[E]struct{}, E comparable](s T, pred func(v E) bool) bool {
	return fx.Any(maps.Keys(s), pred)
}

// FMap returns a sequence of values computed by invoking fn on each element
// of the input set and returning only mapped values for with fn returns
// true.
func FMap[T ~map[E]struct{}, E comparable, U any](s T, fn func(v E) (U, bool)) iter.Seq[U] {
	return fx.FMap(maps.Keys(s), fn)
}

// Filter returns a sequence of values computed by invoking fn on each element
// of the input set and returning only those elements for with fn returns
// true.
func Filter[T ~map[E]struct{}, E comparable](s T, fn func(v E) bool) iter.Seq[E] {
	return fx.Filter(maps.Keys(s), fn)
}

// Map invokes fn on each value in the input set and returns the results.
func Map[T ~map[E]struct{}, E comparable, U any](s T, fn func(v E) U) iter.Seq[U] {
	return fx.Map(maps.Keys(s), fn)
}

// OfType returns a sequence composed of all elements in the input set
// that are of type U.
func OfType[U any, T ~map[E]struct{}, E comparable](s T) iter.Seq[U] {
	return fx.OfType[U](maps.Keys(s))
}

// Reduce calls fn on each element of the input set, passing in the
// current value of the accumulator with each invocation and updating the
// accumulator to the result of fn after each invocation.
func Reduce[T ~map[E]struct{}, E comparable, U any](s T, init U, fn func(acc U, v E) U) U {
	return fx.Reduce(maps.Keys(s), init, fn)
}

// Take returns an iterator that takes at most n values from the input set.
func Take[T ~map[E]struct{}, E comparable](s T, n int) iter.Seq[E] {
	return fx.Take(maps.Keys(s), n)
}

// Values returns a sequence of each value in the input set. The ordering of elements is undefined.
func Values[T ~map[E]struct{}, E comparable](s T) iter.Seq[E] {
	return maps.Keys(s)
}

// Intersection returns a new set that is the intersection of the input sets.
func Intersection[T ~map[E]struct{}, U ~map[E]struct{}, E comparable](s1 T, s2 U) fx.Set[E] {
	return Collect(Filter(s1, func(e E) bool { return fx.Set[E](s2).Has(e) }))
}

// Union returns a new set that is the union of the input sets.
func Union[T ~map[E]struct{}, U ~map[E]struct{}, E comparable](s1 T, s2 U) fx.Set[E] {
	return Collect(fx.Concat(Values(s1), Values(s2)))
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
