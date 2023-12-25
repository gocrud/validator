package validator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIP_Validate(t *testing.T) {
	var ip = "127.0.0.1"
	var errIp = "127.0.0.256"
	err := Validate(ip, IP())
	require.NoError(t, err)
	err = Validate(errIp, IP())
	require.Error(t, err)
	err = Validate(errIp, IP(IpMsg("invalid ip")))
	require.EqualError(t, err, "invalid ip")
}
