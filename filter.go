package fx

import "iter"

func Filter[T any](it iter.Seq[T], fn func(v T) bool) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for v := range it {
			if fn(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
