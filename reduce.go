package fx

import "iter"

func Reduce[T, U any](it iter.Seq[T], init U, fn func(acc U, v T) U) U {
	for v := range it {
		init = fn(init, v)
	}
	return init
}
