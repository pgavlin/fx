package fx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	v, err := OK(42).Unpack()
	assert.Equal(t, 42, v)
	assert.NoError(t, err)
}

func TestErr(t *testing.T) {
	v, err := Err[int](errors.New("foo")).Unpack()
	assert.Equal(t, 0, v)
	assert.Error(t, err)
}

func TestTry(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		v, err := Try(func() (int, error) { return 42, nil }()).Unpack()
		assert.Equal(t, 42, v)
		assert.NoError(t, err)
	})
	t.Run("Err", func(t *testing.T) {
		v, err := Try(func() (int, error) { return 0, errors.New("foo") }()).Unpack()
		assert.Equal(t, 0, v)
		assert.Error(t, err)
	})
}
