package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOneOf(t *testing.T) {
	var str = "123"
	err := Validate(str, OneOf([]any{"123", "345"}))
	require.NoError(t, err)
	var num = 123
	err = Validate(num, OneOf([]any{123, 456, 789}))
	require.NoError(t, err)
	errMsg := "123 not in [456,789]"
	err = Validate(num, OneOf([]any{456, 789}, OneOfMsg(errMsg)))
	require.EqualError(t, err, errMsg)
}
