package sets

import (
	"iter"
	"maps"

	"github.com/pgavlin/fx"
)

func All[T ~map[E]struct{}, E comparable](s T, pred func(v E) bool) bool {
	return fx.All(maps.Keys(s), pred)
}

func Any[T ~map[E]struct{}, E comparable](s T, pred func(v E) bool) bool {
	return fx.Any(maps.Keys(s), pred)
}

func FMap[T ~map[E]struct{}, E comparable, U any](s T, fn func(v E) (U, bool)) iter.Seq[U] {
	return fx.FMap(maps.Keys(s), fn)
}

func Filter[T ~map[E]struct{}, E comparable](s T, fn func(v E) bool) iter.Seq[E] {
	return fx.Filter(maps.Keys(s), fn)
}

func Map[T ~map[E]struct{}, E comparable, U any](s T, fn func(v E) U) iter.Seq[U] {
	return fx.Map(maps.Keys(s), fn)
}

func OfType[T ~map[E]struct{}, E comparable, U any](s T) iter.Seq[U] {
	return fx.OfType[E, U](maps.Keys(s))
}

func Reduce[T ~map[E]struct{}, E comparable, U any](s T, init U, fn func(acc U, v E) U) U {
	return fx.Reduce(maps.Keys(s), init, fn)
}

func Take[T ~map[E]struct{}, E comparable](s T, n int) iter.Seq[E] {
	return fx.Take(maps.Keys(s), n)
}

func Values[T ~map[E]struct{}, E comparable](s T) iter.Seq[E] {
	return maps.Keys(s)
}

func Intersection[T ~map[E]struct{}, U ~map[E]struct{}, E comparable](s1 T, s2 U) fx.Set[E] {
	return Collect(Filter(s1, func(e E) bool { return fx.Set[E](s2).Has(e) }))
}

func Union[T ~map[E]struct{}, U ~map[E]struct{}, E comparable](s1 T, s2 U) fx.Set[E] {
	return Collect(fx.Concat(Values(s1), Values(s2)))
}

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
