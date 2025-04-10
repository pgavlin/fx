package fx

import "iter"

// Take returns an iterator that takes at most n values from its source.
func Take[T any](it iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if n <= 0 || !yield(v) {
				return
			}
			n--
		}
	}
}

// Take2 returns an iterator that takes at most n values from the input slice.
func Take2[T, U any](it iter.Seq2[T, U], n int) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for t, u := range it {
			if n <= 0 || !yield(t, u) {
				return
			}
			n--
		}
	}
}
