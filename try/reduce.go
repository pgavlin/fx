package try

import "iter"

// Reduce calls fn on each element of the input sequence, passing in the
// current value of the accumulator with each invocation and updating the
// accumulator to the result of fn after each invocation.
func Reduce[T, U any](it iter.Seq2[T, error], init U, fn func(acc U, v T, err error) (U, error)) (_ U, err error) {
	for v, err := range it {
		init, err = fn(init, v, err)
		if err != nil {
			return init, err
		}
	}
	return init, nil
}
