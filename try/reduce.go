package try

import "iter"

func Reduce[T, U any](it iter.Seq[T], init U, fn func(acc U, v T) (U, error)) (_ U, err error) {
	for v := range it {
		init, err = fn(init, v)
		if err != nil {
			return init, err
		}
	}
	return init, nil
}
