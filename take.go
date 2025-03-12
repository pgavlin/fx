package fx

import "iter"

// Take returns an iterator that takes at most n values from its source.
func Take[T any](it iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for v := range it {
			if n <= 0 || !yield(v) {
				return
			}
			n--
		}
	}
}
