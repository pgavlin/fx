package fx

import "iter"

// Reduce calls fn on each element of the input sequence, passing in the current value of the accumulator with
// each invocation and updating the accumulator to the result of fn after each invocation.
func Reduce[T, U any](it iter.Seq[T], init U, fn func(acc U, v T) U) U {
	for v := range it {
		init = fn(init, v)
	}
	return init
}
