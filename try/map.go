package try

import "iter"

// FMap returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// mapped values for with fn returns true.
func FMap[T, U any](it iter.Seq2[T, error], fn func(v T, err error) (U, bool, error)) iter.Seq2[U, error] {
	return func(yield func(u U, err error) bool) {
		for v, err := range it {
			if u, ok, err := fn(v, err); ok || err != nil {
				if !yield(u, err) {
					return
				}
			}
		}
	}
}

// Map invokes fn on each value in the input sequence and returns the results.
func Map[T, U any](it iter.Seq2[T, error], fn func(v T, err error) (U, error)) iter.Seq2[U, error] {
	return func(yield func(v U, err error) bool) {
		for v, err := range it {
			if !yield(fn(v, err)) {
				return
			}
		}
	}
}
