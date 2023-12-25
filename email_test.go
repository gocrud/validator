package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmailValidator_Validate(t *testing.T) {
	var email = "123@qq.com"
	var errEmail = "123"
	err := Validate(email, Email())
	require.NoError(t, err)
	err = Validate(errEmail, Email())
	require.Error(t, err)
	err = Validate(errEmail, Email(EmailMsg("invalid email")))
	require.EqualError(t, err, "invalid email")
	err = Validate(email, Email(EmailRex(`^[a-zA-Z0-9._%+-]+@qq.com$`)))
	require.NoError(t, err)
	err = Validate(email, Email(EmailRex(`^[a-zA-Z0-9._%+-]+@163.com$`)))
	require.Error(t, err)
	err = Validate(email, Email(EmailRex(`^[a-zA-Z0-9._%+-]+@163.com$`), EmailMsg("invalid email")))
	require.EqualError(t, err, "invalid email")
}
