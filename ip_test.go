package validator

import "testing"

func TestIP_Validate(t *testing.T) {
	var ip = "192.168.0.12"
	err := Validate(ip, &IP{Message: "ip格式错误"})
	if err != nil {
		t.Error(err)
	}
}
