package validator

import "testing"

func TestValidateAll(t *testing.T) {
	var name = "12345678901"
	var email = "123@qq.com"

	err := ValidateAll(
		Validators{
			Value: name,
			Validators: []Validator{
				&StrLen{Min: 1, Max: 10, Message: "名字长度必须在1到10之间"},
			},
		},
		Validators{
			Value: email,
			Validators: []Validator{
				&Email{Message: "邮箱格式错误"},
			},
		},
	)
	if err != nil {
		t.Error(err)
	}
}
