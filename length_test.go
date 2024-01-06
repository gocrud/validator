package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLength(t *testing.T) {
	var s = ""
	err := Validate(s, Optional(Length(3, 0)))
	require.NoError(t, err)

	s = "12"
	err = Validate(s, Optional(Length(3, 0)))
	require.Error(t, err)
}
