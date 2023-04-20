package fx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {
	cases := []struct {
		min, max int
	}{
		{0, 10},
		{-5, 5},
		{3, 42},
		{1, 100},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[%v, %v)", c.min, c.max), func(t *testing.T) {
			r := Range(c.min, c.max)
			for i := c.min; i < c.max; i++ {
				ok := r.Next()
				require.True(t, ok)
				assert.Equal(t, i, r.Value())
			}
			assert.False(t, r.Next())
		})
	}
}

func TestMinRange(t *testing.T) {
	cases := []struct {
		min int
	}{
		{0},
		{-5},
		{3},
		{1},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("[%v, ...)", c.min), func(t *testing.T) {
			r := MinRange(c.min)
			for i, j := c.min, 0; j < 10; i, j = i+1, j+1 {
				ok := r.Next()
				require.True(t, ok)
				assert.Equal(t, i, r.Value())
			}
		})
	}
}
