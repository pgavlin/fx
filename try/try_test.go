package try

import (
	"io"
	"iter"
	"testing"

	"github.com/pgavlin/fx/v2"
	"github.com/pgavlin/fx/v2/slices"
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
			it:       Map(OK(fx.Range(0, 10)), func(v int, err error) (int, error) { return v, err }),
			expected: fx.OK([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			name:     "error",
			it:       Concat(Map(OK(fx.Range(0, 10)), func(v int, err error) (int, error) { return v, err }), UnpackAll(fx.Only(fx.Err[int](io.EOF)))),
			expected: fx.Err[[]int](io.EOF),
		},
		{
			name:     "short-circuit",
			it:       Concat(UnpackAll(fx.Only(fx.Err[int](io.EOF))), Map(OK(fx.Range(0, 10)), func(v int, err error) (int, error) { panic("unexpected") })),
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
