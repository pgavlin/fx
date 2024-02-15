package fx

import "iter"

func Range(min, max int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		for ; min < max; min++ {
			if !yield(min) {
				return
			}
		}
	}
}

func MinRange(min int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		for ; ; min++ {
			if !yield(min) {
				return
			}
		}
	}
}
