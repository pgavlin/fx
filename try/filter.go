package try

import "iter"

// Filter returns a sequence of values computed by invoking fn on each element of the input sequence and returning only
// those elements for with fn returns true.
func Filter[T any](it iter.Seq2[T, error], fn func(v T, err error) (bool, error)) iter.Seq2[T, error] {
	return func(yield func(v T, err error) bool) {
		for v, err := range it {
			if ok, err := fn(v, err); ok || err != nil {
				if !yield(v, err) {
					return
				}
			}
		}
	}
}
