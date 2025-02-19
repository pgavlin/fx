package try

import "iter"

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
