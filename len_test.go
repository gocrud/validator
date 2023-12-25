package validator

import "testing"

type User struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func TestStrLenValidator_Validate(t *testing.T) {
	var name = "12345678901"
	err := Validate(name, &StrLen{Min: 1, Max: 10, Message: "名字长度必须在1到10之间"})
	if err != nil {
		t.Error(err)
	}

}
