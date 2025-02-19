package try

import "iter"

func FMap[T, U any](it iter.Seq[T], fn func(v T) (U, bool, error)) iter.Seq2[U, error] {
	return func(yield func(u U, err error) bool) {
		for v := range it {
			if u, ok, err := fn(v); ok || err != nil {
				if !yield(u, err) {
					return
				}
			}
		}
	}
}

func Map[T, U any](it iter.Seq[T], fn func(v T) (U, error)) iter.Seq2[U, error] {
	return func(yield func(v U, err error) bool) {
		for v := range it {
			if !yield(fn(v)) {
				return
			}
		}
	}
}
