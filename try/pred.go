package try

import "iter"

// Any returns true if pred returns true for any element of the input sequence.
func All[T any](it iter.Seq2[T, error], pred func(v T, err error) (bool, error)) (bool, error) {
	for v, err := range it {
		if ok, err := pred(v, err); !ok || err != nil {
			return ok, err
		}
	}
	return true, nil
}

// All returns true if pred returns true for every element of the input sequence.
func Any[T any](it iter.Seq2[T, error], pred func(v T, err error) (bool, error)) (bool, error) {
	for v, err := range it {
		if ok, err := pred(v, err); ok || err != nil {
			return ok, err
		}
	}
	return false, nil
}

// And combines a list of predicates into a predicate that returns true if every predicate in the list returns true.
func And[T any](preds ...func(v T, err error) (bool, error)) func(T, error) (bool, error) {
	return func(v T, err error) (bool, error) {
		for _, p := range preds {
			if ok, err := p(v, err); !ok || err != nil {
				return false, err
			}
		}
		return true, nil
	}
}

// Or combines a list of predicates into a predicate that returns true if any predicate in the list returns true.
func Or[T any](preds ...func(v T, err error) (bool, error)) func(T, error) (bool, error) {
	return func(v T, err error) (bool, error) {
		for _, p := range preds {
			ok, err := p(v, err)
			if err != nil {
				return false, err
			}
			if ok {
				return true, err
			}
		}
		return false, nil
	}
}

// Not inverts the result of a predicate.
func Not[T any](pred func(v T, err error) (bool, error)) func(T, error) (bool, error) {
	return func(v T, err error) (bool, error) {
		ok, err := pred(v, err)
		if err != nil {
			return false, err
		}
		return ok, nil
	}
}
