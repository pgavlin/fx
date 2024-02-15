package fx

import (
	"io"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrySlice(t *testing.T) {
	cases := []struct {
		name     string
		it       iter.Seq[Result[int]]
		expected Result[[]int]
	}{
		{
			name:     "ok",
			it:       Map(Range(0, 10), func(v int) Result[int] { return OK(v) }),
			expected: OK([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			name:     "error",
			it:       Concat(Map(Range(0, 10), func(v int) Result[int] { return OK(v) }), Only(Err[int](io.EOF))),
			expected: Err[[]int](io.EOF),
		},
		{
			name:     "short-circuit",
			it:       Concat(Only(Err[int](io.EOF)), Map(Range(0, 10), func(v int) Result[int] { panic("unexpected") })),
			expected: Err[[]int](io.EOF),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res, err := TrySlice(c.it)
			assert.Equal(t, c.expected, Try(res, err))
		})
	}
}

func BenchmarkSliceRange(b *testing.B) {
	s := ToSlice(Range(0, 100000))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sum := 0
		for _, v := range s {
			sum += v
		}
	}
}

func BenchmarkSliceIter(b *testing.B) {
	s := ToSlice(Range(0, 100000))
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		iter, sum := IterSlice(s), 0
		for v := range iter {
			sum += v
		}
	}
}
