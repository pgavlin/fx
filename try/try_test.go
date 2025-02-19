package try

import (
	"io"
	"iter"
	"testing"

	"github.com/pgavlin/fx"
	"github.com/pgavlin/fx/slices"
	"github.com/stretchr/testify/assert"
)

func TestTrySlice(t *testing.T) {
	cases := []struct {
		name     string
		it       iter.Seq2[int, error]
		expected fx.Result[[]int]
	}{
		{
			name:     "ok",
			it:       Map(fx.Range(0, 10), func(v int) (int, error) { return v, nil }),
			expected: fx.OK([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			name:     "error",
			it:       Concat(Map(fx.Range(0, 10), func(v int) (int, error) { return v, nil }), Seq2(fx.Only(fx.Err[int](io.EOF)))),
			expected: fx.Err[[]int](io.EOF),
		},
		{
			name:     "short-circuit",
			it:       Concat(Seq2(fx.Only(fx.Err[int](io.EOF))), Map(fx.Range(0, 10), func(v int) (int, error) { panic("unexpected") })),
			expected: fx.Err[[]int](io.EOF),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res, err := slices.TryCollect(c.it)
			assert.Equal(t, c.expected, fx.Try(res, err))
		})
	}
}
