package slices

import (
	"iter"
	"slices"

	"github.com/pgavlin/fx"
)

func All[T ~[]E, E any](s T, pred func(v E) bool) bool {
	return fx.All(slices.Values(s), pred)
}

func Any[T ~[]E, E any](s T, pred func(v E) bool) bool {
	return fx.Any(slices.Values(s), pred)
}

func FMap[T ~[]E, E, U any](s T, fn func(v E) (U, bool)) iter.Seq[U] {
	return fx.FMap(slices.Values(s), fn)
}

func Filter[T ~[]E, E any](s T, fn func(v E) bool) iter.Seq[E] {
	return fx.Filter(slices.Values(s), fn)
}

func Map[T ~[]E, E, U any](s T, fn func(v E) U) iter.Seq[U] {
	return fx.Map(slices.Values(s), fn)
}

func OfType[T ~[]E, E, U any](s T) iter.Seq[U] {
	return fx.OfType[E, U](slices.Values(s))
}

func Reduce[T ~[]E, E, U any](s T, init U, fn func(acc U, v E) U) U {
	return fx.Reduce(slices.Values(s), init, fn)
}

func Take[T ~[]E, E any](s T, n int) iter.Seq[E] {
	return fx.Take(slices.Values(s), n)
}

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
