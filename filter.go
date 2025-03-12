package fx

import "iter"

// Filter returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// those elements for with fn returns true.
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
