package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOptional(t *testing.T) {
	var ls = []int{1, 2, 3, 4, 5}
	var empty []int
	err := Validate(ls, Optional(SliceLen(1, 5)))
	require.NoError(t, err)
	err = Validate(ls, Optional(SliceLen(1, 4)))
	require.Error(t, err)
	errMsg := "slice length must be between 1 and 4"
	err = Validate(ls, Optional(SliceLen(1, 4, SliceLenMsg(errMsg))))
	require.EqualError(t, err, errMsg)
	err = Validate(empty, Optional(SliceLen(1, 4)))
	require.NoError(t, err)
}
