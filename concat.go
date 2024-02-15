package fx

import "iter"

func Concat[T any](iters ...iter.Seq[T]) iter.Seq[T] {
	return ConcatMany(IterSlice(iters))
}

func ConcatMany[T any, I iter.Seq[iter.Seq[T]]](iters I) iter.Seq[T] {
	return func(yield func(v T) bool) {
		for it := range iters {
			for v := range it {
				if !yield(v) {
					return
				}
			}
		}
	}
}
