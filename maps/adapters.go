package maps

import (
	"iter"

	"github.com/pgavlin/fx/v2"
)

// Pairs returns a sequence of (key, value) entries in m represented as fx.Pairs.
//
// Equivalent to fx.PackAll(maps.All(m)).
func Pairs[M ~map[K]V, K comparable, V any](m M) iter.Seq[fx.Pair[K, V]] {
	return func(yield func(kvp fx.Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(fx.Pack(k, v)) {
				return
			}
		}
	}
}

func pairPred[K comparable, V any](pred func(k K, v V) bool) func(fx.Pair[K, V]) bool {
	return func(kvp fx.Pair[K, V]) bool { return pred(kvp.Fst, kvp.Snd) }
}

// All returns true if pred returns true for every entry of the input map.
func All[M ~map[K]V, K comparable, V any](m M, pred func(k K, v V) bool) bool {
	return fx.All(Pairs(m), pairPred(pred))
}

// Any returns true if pred returns true for any entry of the input map.
func Any[M ~map[K]V, K comparable, V any](m M, pred func(k K, v V) bool) bool {
	return fx.Any(Pairs(m), pairPred(pred))
}

// FMap returns a sequence of values computed by invoking fn on each entry
// of the input map and returning only mapped values for with fn returns
// true.
func FMap[M ~map[K]V, K, L comparable, V, W any](m M, fn func(k K, v V) (L, W, bool)) iter.Seq2[L, W] {
	return fx.UnpackAll(fx.FMap(Pairs(m), func(kvp fx.Pair[K, V]) (fx.Pair[L, W], bool) {
		l, u, ok := fn(kvp.Fst, kvp.Snd)
		return fx.Pack(l, u), ok
	}))
}

// Filter returns a sequence of values computed by invoking fn on each entry
// of the input map and returning only those entries for with fn returns
// true.
func Filter[M ~map[K]V, K comparable, V any](m M, fn func(k K, v V) bool) iter.Seq2[K, V] {
	return fx.UnpackAll(fx.Filter(Pairs(m), pairPred(fn)))
}

// Map invokes fn on each value in the input map and returns the results.
func Map[M ~map[K]V, K, L comparable, V, W any](m M, fn func(k K, v V) (L, W)) iter.Seq2[L, W] {
	return fx.UnpackAll(fx.Map(Pairs(m), func(kvp fx.Pair[K, V]) fx.Pair[L, W] {
		return fx.Pack(fn(kvp.Fst, kvp.Snd))
	}))
}

// Reduce calls fn on each entry of the input map, passing in the
// current value of the accumulator with each invocation and updating the
// accumulator to the result of fn after each invocation.
func Reduce[M ~map[K]V, K comparable, V, U any](m M, init U, fn func(acc U, k K, v V) U) U {
	return fx.Reduce(Pairs(m), init, func(acc U, kvp fx.Pair[K, V]) U {
		return fn(acc, kvp.Fst, kvp.Snd)
	})
}

// TryCollect attempts to collect the key-value pairs in the input sequence into a map. If any pair in the input contains a
// non-nil error, TryCollect halts and returns the collected map up to that point (excluding the value returned with the error)
// and the error.
func TryCollect[K comparable, V any](it iter.Seq2[fx.Pair[K, V], error]) (map[K]V, error) {
	var m map[K]V
	for kvp, err := range it {
		if err != nil {
			return m, err
		}
		if m == nil {
			m = make(map[K]V)
		}
		m[kvp.Fst] = kvp.Snd
	}
	return m, nil
}
