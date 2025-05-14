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

// MapUntil invokes fn with a view of the input sequence as if it had no errors. This view terminates on the
// first error (if any) in the input. The output is an iter.Seq2[U, error]; if an error is encountered in the
// input, the output will terminate with that error.
func MapUntil[T, U any](it iter.Seq2[T, error], fn func(it iter.Seq[T]) iter.Seq[U]) iter.Seq2[U, error] {
	var err error
	ts := func(yield func(T) bool) {
		var t T
		for t, err = range it {
			if err != nil || !yield(t) {
				break
			}
		}
	}
	return func(yield func(U, error) bool) {
		for u := range fn(ts) {
			if !yield(u, nil) {
				return
			}
		}
		if err != nil {
			var u U
			yield(u, err)
		}
	}
}
