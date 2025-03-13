package try

import (
	"iter"

	"github.com/pgavlin/fx/v2"
)

// PackAll transforms a sequence of (T, error) pairs into a sequence of Result[T] values.
func PackAll[T any](it iter.Seq2[T, error]) iter.Seq[fx.Result[T]] {
	return func(yield func(r fx.Result[T]) bool) {
		for v, err := range it {
			if !yield(fx.Try(v, err)) {
				break
			}
		}
	}
}

// UnpackAll transforms a sequence of Result[T] values into a sequence of (T, error) pairs.
func UnpackAll[T any](it iter.Seq[fx.Result[T]]) iter.Seq2[T, error] {
	return func(yield func(t T, err error) bool) {
		for v := range it {
			if !yield(v.Unpack()) {
				break
			}
		}
	}
}

// Must transforms a sequence of (T, error) pairs into a sequence of T values. If any element of the input sequence
// contains a non-nil error, the returned iterator will panic when that error is observed.
func Must[T any](it iter.Seq2[T, error]) iter.Seq[T] {
	return func(yield func(t T) bool) {
		for v, err := range it {
			if err != nil {
				panic(err)
			}
			if !yield(v) {
				break
			}
		}
	}
}

// OK transforms a sequence of values into a sequence of (T, nil) pairs.
func OK[T any](it iter.Seq[T]) iter.Seq2[T, error] {
	return func(yield func(t T, err error) bool) {
		for v := range it {
			if !yield(v, nil) {
				break
			}
		}
	}
}

// NoError returns true if err is nil.
func NoError[T any](_ T, err error) bool {
	return err == nil
}
