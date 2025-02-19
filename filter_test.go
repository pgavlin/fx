package fx

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	cases := []struct {
		name     string
		it       iter.Seq[int]
		pred     func(v int) bool
		expected []int
	}{
		{
			name:     "false",
			it:       Range(0, 10),
			pred:     func(_ int) bool { return false },
			expected: nil,
		},
		{
			name:     "true",
			it:       Range(0, 10),
			pred:     func(_ int) bool { return true },
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "evens",
			it:       Range(0, 10),
			pred:     func(v int) bool { return v%2 == 0 },
			expected: []int{0, 2, 4, 6, 8},
		},
		{
			name:     "odds",
			it:       Range(0, 10),
			pred:     func(v int) bool { return v%2 != 0 },
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "primes",
			it:       Range(0, 10),
			pred:     func(v int) bool { return (v == 2 || v%2 != 0) && (v == 3 || v%3 != 0) },
			expected: []int{1, 2, 3, 5, 7},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := slices.Collect(Filter(c.it, c.pred))
			assert.Equal(t, c.expected, actual)
		})
	}
}
