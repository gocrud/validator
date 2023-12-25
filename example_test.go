package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type User struct {
	Name  string
	Age   int
	Email string
	Tags  []string
	State string // on or off
}

func TestValidate(t *testing.T) {
	var user = User{
		Name:  "Peter",
		Age:   18,
		Email: "peter@qq.com",
		Tags:  []string{"Student", "Cool"},
		State: "one",
	}
	err := ValidateMultiple(
		Group{
			Field:      user.Name,
			Validators: []Validator{StrLen(1, 10, StrLenMsg("name length must be between 1 and 10"))},
		},
		Group{
			Field:      user.Age,
			Validators: []Validator{IntRange(1, 100, IntRangeMsg("age must be between 1 and 100"))},
		},
		Group{
			Field:      user.Email,
			Validators: []Validator{Email(EmailRex("^[a-zA-Z0-9._%+-]+@qq.com$"))},
		},
		Group{
			Field:      user.Tags,
			Validators: []Validator{SliceLen(1, 2)},
		},
		Group{
			Field:      user.State,
			Validators: []Validator{OneOf([]any{"one", "off"})},
		},
	)
	require.NoError(t, err)
}
