package try

import "iter"

// Take returns an iterator that takes at most n values from the input sequence.
func Take[T any](it iter.Seq2[T, error], n int) iter.Seq2[T, error] {
	return func(yield func(v T, err error) bool) {
		for v, err := range it {
			if n <= 0 || !yield(v, err) {
				return
			}
			n--
		}
	}
}
