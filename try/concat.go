package try

import (
	"iter"
	"slices"
)

// Concat returns an iterator that returns values from each iterator in sequence.
func Concat[T any](iters ...iter.Seq2[T, error]) iter.Seq2[T, error] {
	return ConcatMany(slices.Values(iters))
}

// ConcatMany returns an iterator that returns values from each iterator in sequence.
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
