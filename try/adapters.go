package try

import (
	"iter"

	"github.com/pgavlin/fx"
)

func Results[T any](it iter.Seq2[T, error]) iter.Seq[fx.Result[T]] {
	return func(yield func(r fx.Result[T]) bool) {
		for v, err := range it {
			if !yield(fx.Try(v, err)) {
				break
			}
		}
	}
}

func Seq2[T any](it iter.Seq[fx.Result[T]]) iter.Seq2[T, error] {
	return func(yield func(t T, err error) bool) {
		for v := range it {
			if !yield(v.Unpack()) {
				break
			}
		}
	}
}
