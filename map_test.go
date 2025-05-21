package fx

import (
	"iter"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFMap(t *testing.T) {
	cases := []struct {
		name     string
		it       iter.Seq[int]
		pred     func(v int) (string, bool)
		expected []string
	}{
		{
			name:     "false",
			it:       Range(0, 10),
			pred:     func(_ int) (string, bool) { return "", false },
			expected: nil,
		},
		{
			name:     "true",
			it:       Range(0, 10),
			pred:     func(v int) (string, bool) { return strconv.FormatInt(int64(v), 10), true },
			expected: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			name:     "evens",
			it:       Range(0, 10),
			pred:     func(v int) (string, bool) { return strconv.FormatInt(int64(v), 10), v%2 == 0 },
			expected: []string{"0", "2", "4", "6", "8"},
		},
		{
			name:     "odds",
			it:       Range(0, 10),
			pred:     func(v int) (string, bool) { return strconv.FormatInt(int64(v), 10), v%2 != 0 },
			expected: []string{"1", "3", "5", "7", "9"},
		},
		{
			name: "primes",
			it:   Range(0, 10),
			pred: func(v int) (string, bool) {
				return strconv.FormatInt(int64(v), 10), (v == 2 || v%2 != 0) && (v == 3 || v%3 != 0)
			},
			expected: []string{"1", "2", "3", "5", "7"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := slices.Collect(FMap(c.it, c.pred))
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestMap(t *testing.T) {
	cases := []struct {
		name     string
		it       iter.Seq[int]
		mapf     func(v int) string
		expected []string
	}{
		{
			name:     "empty",
			it:       Empty[int](),
			mapf:     func(_ int) string { return "" },
			expected: nil,
		},
		{
			name:     "decimal",
			it:       Range(0, 10),
			mapf:     func(v int) string { return strconv.FormatInt(int64(v), 10) },
			expected: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		},
		{
			name:     "hex",
			it:       Range(10, 16),
			mapf:     func(v int) string { return strconv.FormatInt(int64(v), 16) },
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := slices.Collect(Map(c.it, c.mapf))
			assert.Equal(t, c.expected, actual)
		})
	}
}

func TestOfType(t *testing.T) {
	seq1 := slices.Values([]any{1, 2, "three"})
	seq2 := slices.Collect(OfType[int](seq1))
	assert.Equal(t, []int{1, 2}, seq2)

}
