package fx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			i := c.min
			for v := range Range(c.min, c.max) {
				assert.Equal(t, i, v)
				i++
			}
			assert.Equal(t, c.max, i)
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
			i, j := c.min, 0
			for v := range MinRange(c.min) {
				if j == 10 {
					return
				}
				assert.Equal(t, i, v)
				i, j = i+1, j+1
			}
		})
	}
}
