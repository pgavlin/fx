package fx

import "iter"

// Only returns a sequence that contains the single value v.
func Only[T any](v T) iter.Seq[T] {
	return func(yield func(T) bool) {
		yield(v)
	}
}

// Only2 returns a sequence that contains the single value v.
func Only2[T, U any](t T, u U) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		yield(t, u)
	}
}

// Empty returns an empty sequence.
func Empty[T any]() iter.Seq[T] {
	return func(_ func(T) bool) {}
}

// Empty2 returns an empty sequence.
func Empty2[T, U any]() iter.Seq2[T, U] {
	return func(_ func(T, U) bool) {}
}
