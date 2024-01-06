package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOneOf(t *testing.T) {
	err := Validate(3, OneOf(1, 2).Msg("must be one of 1, 2"))
	require.EqualError(t, err, "must be one of 1, 2")
	err = Validate(3, OneOf(1, 2, 3))
	require.NoError(t, err)
}
