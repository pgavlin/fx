package fx

import (
	"iter"
	"sort"
)

func Sorted[T any](it iter.Seq[T], less func(a, b T) bool) iter.Seq[T] {
	s := ToSlice(it)
	sort.Slice(s, func(i, j int) bool { return less(s[i], s[j]) })
	return IterSlice(s)
}
