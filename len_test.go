package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayLen(t *testing.T) {
	var ls = [5]int{1, 2, 3, 4, 5}
	err := Validate(ls, ArrayLen(1, 5))
	require.NoError(t, err)
	err = Validate(ls, ArrayLen(1, 4))
	require.Error(t, err)
	err = Validate(ls, ArrayLen(1, 4, ArrayLenMsg("array length must be between 1 and 4")))
	require.EqualError(t, err, "array length must be between 1 and 4")
}

func TestSliceLen(t *testing.T) {
	var ls = []int{1, 2, 3, 4, 5}
	err := Validate(ls, SliceLen(1, 5))
	require.NoError(t, err)
	err = Validate(ls, SliceLen(1, 4))
	require.Error(t, err)
	err = Validate(ls, SliceLen(1, 4, SliceLenMsg("slice length must be between 1 and 4")))
	require.EqualError(t, err, "slice length must be between 1 and 4")
}

func TestStrLen(t *testing.T) {
	var str = "123456"
	err := Validate(str, StrLen(1, 6))
	require.NoError(t, err)
	err = Validate(str, StrLen(1, 5))
	require.Error(t, err)
	err = Validate(str, StrLen(1, 5, StrLenMsg("string length must be between 1 and 5")))
	require.EqualError(t, err, "string length must be between 1 and 5")
}
