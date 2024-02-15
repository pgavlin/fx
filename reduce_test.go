package fx

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	cases := []struct {
		name     string
		it       iter.Seq[int]
		init     int
		fn       func(acc, v int) int
		expected int
	}{
		{
			name:     "sum",
			it:       Range(1, 11),
			init:     0,
			fn:       func(acc, v int) int { return acc + v },
			expected: 55,
		},
		{
			name:     "difference",
			it:       Range(1, 11),
			init:     55,
			fn:       func(acc, v int) int { return acc - v },
			expected: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := Reduce(c.it, c.init, c.fn)
			assert.Equal(t, c.expected, actual)
		})
	}
}
