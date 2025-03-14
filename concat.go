package fx

import (
	"iter"
	"slices"
)

// Concat returns an iterator that returns values from each iterator in sequence.
func Concat[T any](iters ...iter.Seq[T]) iter.Seq[T] {
	return ConcatMany(slices.Values(iters))
}

// ConcatMany returns an iterator that returns values from each iterator in sequence.
func ConcatMany[T any](iters iter.Seq[iter.Seq[T]]) iter.Seq[T] {
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
