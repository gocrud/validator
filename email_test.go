package validator

import "testing"

func TestEmailValidator_Validate(t *testing.T) {
	var email = "12345678901@qq.com"
	err := Validate(email, &Email{Message: "邮箱格式错误"})
	if err != nil {
		t.Error(err)
	}
}
