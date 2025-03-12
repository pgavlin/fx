package fx

import "iter"

// Range returns a sequence of each integer in the range [min, max).
func Range(min, max int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		for ; min < max; min++ {
			if !yield(min) {
				return
			}
		}
	}
}

// MinRange returns a sequence of each integer in the range [min, ∞)
func MinRange(min int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		for ; ; min++ {
			if !yield(min) {
				return
			}
		}
	}
}
