package maps

import (
	"iter"

	"github.com/pgavlin/fx"
)

func Pairs[M ~map[K]V, K comparable, V any](m M) iter.Seq[fx.Pair[K, V]] {
	return func(yield func(kvp fx.Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(fx.NewPair(k, v)) {
				return
			}
		}
	}
}

func Seq2[K comparable, V any](it iter.Seq[fx.Pair[K, V]]) iter.Seq2[K, V] {
	return func(yield func(k K, v V) bool) {
		for kvp := range it {
			if !yield(kvp.Fst, kvp.Snd) {
				return
			}
		}
	}
}

func pairPred[K comparable, V any](pred func(k K, v V) bool) func(fx.Pair[K, V]) bool {
	return func(kvp fx.Pair[K, V]) bool { return pred(kvp.Fst, kvp.Snd) }
}

func All[M ~map[K]V, K comparable, V any](m M, pred func(k K, v V) bool) bool {
	return fx.All(Pairs(m), pairPred(pred))
}

func Any[M ~map[K]V, K comparable, V any](m M, pred func(k K, v V) bool) bool {
	return fx.Any(Pairs(m), pairPred(pred))
}

func FMap[M ~map[K]V, K, L comparable, V, W any](m M, fn func(k K, v V) (L, W, bool)) iter.Seq2[L, W] {
	return fx.Seq2(fx.FMap(Pairs(m), func(kvp fx.Pair[K, V]) (fx.Pair[L, W], bool) {
		l, u, ok := fn(kvp.Fst, kvp.Snd)
		return fx.NewPair(l, u), ok
	}))
}

func Filter[M ~map[K]V, K comparable, V any](m M, fn func(k K, v V) bool) iter.Seq2[K, V] {
	return fx.Seq2(fx.Filter(Pairs(m), pairPred(fn)))
}

func Map[M ~map[K]V, K, L comparable, V, W any](m M, fn func(k K, v V) (L, W)) iter.Seq2[L, W] {
	return fx.Seq2(fx.Map(Pairs(m), func(kvp fx.Pair[K, V]) fx.Pair[L, W] {
		return fx.NewPair(fn(kvp.Fst, kvp.Snd))
	}))
}

func Reduce[M ~map[K]V, K comparable, V, U any](m M, init U, fn func(acc U, k K, v V) U) U {
	return fx.Reduce(Pairs(m), init, func(acc U, kvp fx.Pair[K, V]) U {
		return fn(acc, kvp.Fst, kvp.Snd)
	})
}

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
