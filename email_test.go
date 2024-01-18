package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmailValidator_Validate(t *testing.T) {
	err := Validate("xyz@abc.com", Email())
	require.NoError(t, err)

	err = Validate("xyz", Email())
	require.Error(t, err)

	err = Validate("xyz@", Email().Msg("invalid email"))
	require.Error(t, err)

	err = Validate("xyz@abc", Email().Regex(`^[a-zA-Z0-9._%+-]+@qq\.com$`))
	require.Error(t, err)
}
