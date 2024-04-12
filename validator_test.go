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

func TestIP_Validate(t *testing.T) {
	err := Validate("127.0.0.1", IP())
	require.NoError(t, err)

	err = Validate("127.0.0.1:8080", IP())
	require.Error(t, err)

	err = Validate("127.0.0.256", IP())
	require.Error(t, err)
}

func TestLength(t *testing.T) {
	var s = ""
	err := Validate(s, Optional(Range(3, 5)))
	require.NoError(t, err)

	s = "12"
	err = Validate(s, Optional(Range(3, 5)))
	require.Error(t, err)
}

func TestOneOf(t *testing.T) {
	err := Validate(3, In(1, 2).Msg("must be one of 1, 2"))
	require.EqualError(t, err, "must be one of 1, 2")
	err = Validate(3, In(1, 2, 3))
	require.NoError(t, err)
}

func TestOr(t *testing.T) {
	var str = "https://www.google.com"
	err := Validate(str, Or(Prefix("https://"), Prefix("http://")))
	require.NoError(t, err)
	err = Validate(str, Or(Prefix("https://"), Prefix("ftp://")))
	require.NoError(t, err)
	err = Validate(str, Or(Prefix("ftp://"), Prefix("https://")))
	require.NoError(t, err)
	err = Validate(str, Or(Prefix("ftp://"), Prefix("ftp://")))
	require.Error(t, err)
	err = Validate(str, Or(Prefix("ftp://"), Prefix("ftp://")).Msg("invalid url"))
	require.EqualError(t, err, "invalid url")
	err = Validate(str, Or(Prefix("ftp://"), Suffix(".com")).Msg("invalid url"))
	require.NoError(t, err)
}

func TestPrefix(t *testing.T) {
	var str string = "hello world"
	var str2 string = "https://www.google.com"
	err := Validate(&str, Prefix("https://"))
	require.Error(t, err)
	err = Validate(&str2, Prefix("https://"))
	require.NoError(t, err)
}

func TestRange(t *testing.T) {
	var i = 0
	err := Validate(i, Range(3, 0))
	require.NoError(t, err)

	i = 2
	err = Validate(i, Range(3, 6).Msg("取值范围为3-6"))
	require.Error(t, err)
}

func TestMax(t *testing.T) {
	var i = 0
	err := Validate(i, Max(3))
	require.NoError(t, err)

	i = 4
	err = Validate(i, Max(3).Msg("不能大于3"))
	require.Error(t, err)
}

func TestMin(t *testing.T) {
	var i = 4
	err := Validate(i, Min(3))
	require.NoError(t, err)

	i = 2
	err = Validate(i, Min(3).Msg("不能小于3"))
	require.Error(t, err)
}

func TestSuffix(t *testing.T) {
	var str string = "hello world"
	var str2 string = "https://www.google.com"
	err := Validate(&str, Suffix(".com"))
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	err = Validate(&str2, Suffix(".com"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegex(t *testing.T) {
	var str = "hello world"
	err := Validate(&str, Regexp(`^hello`))
	require.NoError(t, err)
	err = Validate(&str, Regexp(`^world`))
	require.Error(t, err)
	err = Validate(&str, Regexp(`^world`).Msg("must start with world"))
	require.EqualError(t, err, "must start with world")
}
