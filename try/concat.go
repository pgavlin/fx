package try

import (
	"iter"
	"slices"
)

func Concat[T any](iters ...iter.Seq2[T, error]) iter.Seq2[T, error] {
	return ConcatMany(slices.Values(iters))
}

func ConcatMany[T any](iters iter.Seq[iter.Seq2[T, error]]) iter.Seq2[T, error] {
	return func(yield func(v T, err error) bool) {
		for it := range iters {
			for v, err := range it {
				if !yield(v, err) {
					return
				}
			}
		}
	}
}
