package sets

import (
	"iter"
	"maps"

	"github.com/pgavlin/fx/v2"
)

// All returns true if pred returns true for every element of the input sequence.
func All[S ~map[T]struct{}, T comparable](s S, pred func(v T) bool) bool {
	return fx.All(maps.Keys(s), pred)
}

// Any returns true if pred returns true for any element of the input sequence.
func Any[S ~map[T]struct{}, T comparable](s S, pred func(v T) bool) bool {
	return fx.Any(maps.Keys(s), pred)
}

// Contains returns true if the input sequence contains t.
func Contains[S ~map[T]struct{}, T comparable](s S, t T) bool {
	return fx.Contains(maps.Keys(s), t)
}

// FMap returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// mapped values for with fn returns true.
func FMap[S ~map[T]struct{}, T comparable, U any](s S, fn func(v T) (U, bool)) iter.Seq[U] {
	return fx.FMap(maps.Keys(s), fn)
}

// FMapUnpack returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// mapped values for with fn returns true.
func FMapUnpack[S ~map[T]struct{}, T comparable, U any, V any](s S, fn func(v T) (U, V, bool)) iter.Seq2[U, V] {
	return fx.FMapUnpack(maps.Keys(s), fn)
}

// Filter returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// those elements for with fn returns true.
func Filter[S ~map[T]struct{}, T comparable](s S, fn func(v T) bool) iter.Seq[T] {
	return fx.Filter(maps.Keys(s), fn)
}

// First returns the first element of it, if any elements exist.
func First[S ~map[T]struct{}, T comparable](s S) (T, bool) {
	return fx.First(maps.Keys(s))
}

// Last returns the last element of it, if any elements exist.
func Last[S ~map[T]struct{}, T comparable](s S) (T, bool) {
	return fx.Last(maps.Keys(s))
}

// Map invokes fn on each value in the input sequence and returns the results.
func Map[S ~map[T]struct{}, T comparable, U any](s S, fn func(v T) U) iter.Seq[U] {
	return fx.Map(maps.Keys(s), fn)
}

// MapUnpack invokes fn on each value in the input sequence and returns the results.
func MapUnpack[S ~map[T]struct{}, T comparable, U any, V any](s S, fn func(v T) (U, V)) iter.Seq2[U, V] {
	return fx.MapUnpack(maps.Keys(s), fn)
}

// OfType returns a sequence composed of all elements in the input sequence that are of type U.
func OfType[U any, S ~map[T]struct{}, T comparable](s S) iter.Seq[U] {
	return fx.OfType[U](maps.Keys(s))
}

// Reduce calls fn on each element of the input sequence, passing in the current value of the accumulator with
// each invocation and updating the accumulator to the result of fn after each invocation.
func Reduce[S ~map[T]struct{}, T comparable, U any](s S, init U, fn func(acc U, v T) U) U {
	return fx.Reduce(maps.Keys(s), init, fn)
}

// Skip returns an iterator that skips n values from its source.
func Skip[S ~map[T]struct{}, T comparable](s S, n int) iter.Seq[T] {
	return fx.Skip(maps.Keys(s), n)
}

// Take returns an iterator that takes at most n values from its source.
func Take[S ~map[T]struct{}, T comparable](s S, n int) iter.Seq[T] {
	return fx.Take(maps.Keys(s), n)
}
