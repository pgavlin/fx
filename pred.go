package fx

import "iter"

// Any returns true if pred returns true for any element of the input sequence.
func Any[T any](it iter.Seq[T], pred func(v T) bool) bool {
	for v := range it {
		if pred(v) {
			return true

		}
	}
	return false
}

// All returns true if pred returns true for every element of the input sequence.
func All[T any](it iter.Seq[T], pred func(v T) bool) bool {
	for v := range it {
		if !pred(v) {
			return false
		}
	}
	return true
}

// And combines a list of predicates into a predicate that returns true if every predicate in the list returns true.
func And[T any](preds ...func(v T) bool) func(T) bool {
	return func(v T) bool {
		for _, p := range preds {
			if !p(v) {
				return false
			}
		}
		return true
	}
}

// Or combines a list of predicates into a predicate that returns true if any predicate in the list returns true.
func Or[T any](preds ...func(v T) bool) func(T) bool {
	return func(v T) bool {
		for _, p := range preds {
			if p(v) {
				return true
			}
		}
		return false
	}
}

// Not inverts the result of a predicate.
func Not[T any](pred func(v T) bool) func(T) bool {
	return func(v T) bool { return !pred(v) }
}
