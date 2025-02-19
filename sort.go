package fx

import (
	"cmp"
	"iter"
	"slices"
)

func Sorted[T cmp.Ordered](it iter.Seq[T]) []T {
	return slices.Sorted(it)
}

func SortedFunc[T any](it iter.Seq[T], cmp func(a, b T) int) []T {
	return slices.SortedFunc(it, cmp)
}
