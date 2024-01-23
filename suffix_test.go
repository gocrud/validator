package validator

import "testing"

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
