package fx

import "iter"

func Any[T any](it iter.Seq[T], pred func(v T) bool) bool {
	for v := range it {
		if pred(v) {
			return true

		}
	}
	return false
}

func All[T any](it iter.Seq[T], pred func(v T) bool) bool {
	for v := range it {
		if !pred(v) {
			return false
		}
	}
	return true
}
