package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLength(t *testing.T) {
	var s = 3
	err := Validate(s, Optional(Length(3, 0)))
	require.NoError(t, err)
}
