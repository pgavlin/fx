package try

import (
	"errors"
	"iter"
	"testing"

	"github.com/pgavlin/fx/v2"
	"github.com/pgavlin/fx/v2/slices"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapUntil(t *testing.T) {
	result, err := slices.TryCollect(MapUntil(OK(fx.Range(0, 10)), func(i iter.Seq[int]) iter.Seq[int] { return i }))
	require.NoError(t, err)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)

	result, err = slices.TryCollect(MapUntil(fx.Concat2(OK(fx.Range(0, 10)), fx.Only2(0, errors.New("oh no"))), func(i iter.Seq[int]) iter.Seq[int] { return i }))
	assert.Error(t, err)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
}
