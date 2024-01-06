package validator

import "testing"

func TestIntRange(t *testing.T) {
	var i = 0
	err := Validate(i, Optional(IntRange(3, 0).Msg("最小值为3")))
	if err != nil {
		t.Fatal(err)
	}
}
