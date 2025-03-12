package slices

import (
	"iter"
	"slices"

	"github.com/pgavlin/fx/v2"
)

// All returns true if pred returns true for every element of the input slice.
func All[T ~[]E, E any](s T, pred func(v E) bool) bool {
	return fx.All(slices.Values(s), pred)
}

// Any returns true if pred returns true for any element of the input slice.
func Any[T ~[]E, E any](s T, pred func(v E) bool) bool {
	return fx.Any(slices.Values(s), pred)
}

// FMap returns a sequence of values computed by invoking fn on each element
// of the input slice and returning only mapped values for with fn returns
// true.
func FMap[T ~[]E, E, U any](s T, fn func(v E) (U, bool)) iter.Seq[U] {
	return fx.FMap(slices.Values(s), fn)
}

// Filter returns a sequence of values computed by invoking fn on each element
// of the input slice and returning only those elements for with fn returns
// true.
func Filter[T ~[]E, E any](s T, fn func(v E) bool) iter.Seq[E] {
	return fx.Filter(slices.Values(s), fn)
}

// Map invokes fn on each value in the input slice and returns the results.
func Map[T ~[]E, E, U any](s T, fn func(v E) U) iter.Seq[U] {
	return fx.Map(slices.Values(s), fn)
}

// OfType returns a sequence composed of all elements in the input slice
// that are of type U.
func OfType[T ~[]E, E, U any](s T) iter.Seq[U] {
	return fx.OfType[E, U](slices.Values(s))
}

// Reduce calls fn on each element of the input slice, passing in the
// current value of the accumulator with each invocation and updating the
// accumulator to the result of fn after each invocation.
func Reduce[T ~[]E, E, U any](s T, init U, fn func(acc U, v E) U) U {
	return fx.Reduce(slices.Values(s), init, fn)
}

// Take returns an iterator that takes at most n values from the input slice.
func Take[T ~[]E, E any](s T, n int) iter.Seq[E] {
	return fx.Take(slices.Values(s), n)
}

// TryCollect attempts to collect the values in the input sequence into a slice. If any pair in the input contains a
// non-nil error, TryCollect halts and returns the values collected up to that point (excluding the value returned
// with the error) and the error.
func TryCollect[T any](it iter.Seq2[T, error]) ([]T, error) {
	var s []T
	for t, err := range it {
		if err != nil {
			return s, err
		}
		s = append(s, t)
	}
	return s, nil
}
