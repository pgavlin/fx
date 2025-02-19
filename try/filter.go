package try

import "iter"

func Filter[T any](it iter.Seq[T], fn func(v T) (bool, error)) iter.Seq2[T, error] {
	return func(yield func(v T, err error) bool) {
		for v := range it {
			if ok, err := fn(v); ok || err != nil {
				if !yield(v, err) {
					return
				}
			}
		}
	}
}
