package fx

import "iter"

// OfType returns a sequence composed of all elements in the input sequence that are of type U.
func OfType[T, U any](it iter.Seq[T]) iter.Seq[U] {
	return FMap(it, func(v T) (U, bool) {
		u, ok := any(v).(U)
		return u, ok
	})
}

// FMap returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// mapped values for with fn returns true.
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

// Map invokes fn on each value in the input sequence and returns the results.
func Map[T, U any](it iter.Seq[T], fn func(v T) U) iter.Seq[U] {
	return func(yield func(v U) bool) {
		for v := range it {
			if !yield(fn(v)) {
				return
			}
		}
	}
}
