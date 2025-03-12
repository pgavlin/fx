package fx

import "iter"

// Only returns a sequence that contains the single value v.
func Only[T any](v T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		yield(v)
	}
}

// Empty returns an empty sequence.
func Empty[T any]() iter.Seq[T] {
	return func(_ func(v T) bool) {}
}
