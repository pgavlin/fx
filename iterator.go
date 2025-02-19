package fx

import "iter"

func Only[T any](v T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		yield(v)
	}
}

func Empty[T any]() iter.Seq[T] {
	return func(_ func(v T) bool) {}
}
