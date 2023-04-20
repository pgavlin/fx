package fx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	cases := []struct {
		it       Iterator[int]
		n        int
		expected []int
	}{
		{
			it:       MinRange(0),
			n:        10,
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			it:       Range(0, 5),
			n:        10,
			expected: []int{0, 1, 2, 3, 4},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.expected), func(t *testing.T) {
			actual := ToSlice(Take(c.it, c.n))
			assert.Equal(t, c.expected, actual)
		})
	}
}
