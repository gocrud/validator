package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrefix(t *testing.T) {
	var str string = "hello world"
	var str2 string = "https://www.google.com"
	err := Validate(&str, Prefix("https://"))
	require.Error(t, err)
	err = Validate(&str2, Prefix("https://"))
	require.NoError(t, err)
}
