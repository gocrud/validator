package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntRange(t *testing.T) {
	var i = 0
	err := Validate(i, Optional(IntRange(3, 0).Msg("最小值为3")))
	require.NoError(t, err)

	i = 2
	err = Validate(i, Optional(IntRange(3, 6).Msg("取值范围为3-6")))
	require.Error(t, err)
}
