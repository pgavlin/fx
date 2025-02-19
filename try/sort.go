package try

import (
	"cmp"
	"iter"
	"slices"

	fxs "github.com/pgavlin/fx/slices"
)

func Sorted[T cmp.Ordered](it iter.Seq2[T, error]) ([]T, error) {
	s, err := fxs.TryCollect(it)
	if err != nil {
		return nil, err
	}
	slices.Sort(s)
	return s, nil
}

func SortedFunc[T any](it iter.Seq2[T, error], cmp func(a, b T) int) ([]T, error) {
	s, err := fxs.TryCollect(it)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(s, cmp)
	return s, nil
}
