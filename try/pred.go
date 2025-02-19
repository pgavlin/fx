package try

import "iter"

func All[T any](it iter.Seq[T], pred func(v T) (bool, error)) (bool, error) {
	for v := range it {
		if ok, err := pred(v); !ok || err != nil {
			return ok, err
		}
	}
	return true, nil
}

func Any[T any](it iter.Seq[T], pred func(v T) (bool, error)) (bool, error) {
	for v := range it {
		if ok, err := pred(v); ok || err != nil {
			return ok, err
		}
	}
	return false, nil
}
