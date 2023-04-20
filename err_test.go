package fx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	v := OK(42)
	assert.Equal(t, 42, v.Value())
	assert.NoError(t, v.Err())
}

func TestErr(t *testing.T) {
	v := Err[int](errors.New("foo"))
	assert.Equal(t, 0, v.Value())
	assert.Error(t, v.Err())
}

func TestTry(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		v := Try(func() (int, error) { return 42, nil }())
		assert.Equal(t, 42, v.Value())
		assert.NoError(t, v.Err())
	})
	t.Run("Err", func(t *testing.T) {
		v := Try(func() (int, error) { return 0, errors.New("foo") }())
		assert.Equal(t, 0, v.Value())
		assert.Error(t, v.Err())
	})
}
