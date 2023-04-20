package fx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrySlice(t *testing.T) {
	cases := []struct {
		name     string
		it       Iterator[Result[int]]
		expected []int
		err      bool
	}{
		{
			name:     "ok",
			it:       Map(Range(0, 10), func(v int) Result[int] { return OK(v) }),
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "error",
			it:   Concat(Map(Range(0, 10), func(v int) Result[int] { return OK(v) }), Only(Err[int](errors.New("foo")))),
			err:  true,
		},
		{
			name: "short-circuit",
			it:   Concat(Only(Err[int](errors.New("foo"))), Map(Range(0, 10), func(v int) Result[int] { panic("unexpected") })),
			err:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := TrySlice(c.it)
			assert.Equal(t, c.expected, res.Value())
			if c.err {
				assert.Error(t, res.Err())
			}
		})
	}
}
