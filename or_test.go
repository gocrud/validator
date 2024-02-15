package validator

import (
    "github.com/stretchr/testify/require"
    "testing"
)

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
