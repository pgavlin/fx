package fx

import (
	"iter"
	"slices"
)

func Concat[T any](iters ...iter.Seq[T]) iter.Seq[T] {
	return ConcatMany(slices.Values(iters))
}

func ConcatMany[T any, I iter.Seq[iter.Seq[T]]](iters I) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for it := range iters {
			for v := range it {
				if !yield(v) {
					return
				}
			}
		}
	}
}
