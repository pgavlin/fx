package fx

import (
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected []int
	}{
		{[][]int{{}, {}}, nil},
		{[][]int{{0, 1, 2, 3}, {4, 5, 6, 7}}, []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{[][]int{{0, 1}, {2, 3}, {4, 5}}, []int{0, 1, 2, 3, 4, 5}},
		{[][]int{{0, 2}, {4, 6, 8, 10}, {12}}, []int{0, 2, 4, 6, 8, 10, 12}},
		{[][]int{{1}, {3}, {5, 7}, {11, 13, 17}, {19, 23, 29, 31, 37}}, []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37}},
	}
	for _, c := range cases {
		iters := make([]iter.Seq[int], len(c.input))
		for i, in := range c.input {
			iters[i] = IterSlice(in)
		}
		actual := ToSlice(Concat(iters...))
		assert.Equal(t, c.expected, actual)
	}
}
