package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIP_Validate(t *testing.T) {
	err := Validate("127.0.0.1", IP())
	require.NoError(t, err)

	err = Validate("127.0.0.1:8080", IP())
	require.Error(t, err)

	err = Validate("127.0.0.256", IP())
	require.Error(t, err)
}
