package fx

import "iter"

func OfType[T, U any](it iter.Seq[T]) iter.Seq[U] {
	return FMap[T, U](it, func(v T) (U, bool) {
		u, ok := ((interface{})(v)).(U)
		return u, ok
	})
}

func FMap[T, U any](it iter.Seq[T], fn func(v T) (U, bool)) iter.Seq[U] {
	return func(yield func(v U) bool) {
		for v := range it {
			if u, ok := fn(v); ok {
				if !yield(u) {
					return
				}
			}
		}
	}
}

func Map[T, U any](it iter.Seq[T], fn func(v T) U) iter.Seq[U] {
	return func(yield func(v U) bool) {
		for v := range it {
			if !yield(fn(v)) {
				return
			}
		}
	}
}
