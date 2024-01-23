package validator

type SuffixValidator struct {
	AbstractValidator
	suffix string
}

func (s *SuffixValidator) Validate() error {
	if !isString(s.ref) {
		return s.Error("Suffix: value must be a string")
	}

	if !hasSuffix(s.ref.String(), s.suffix) {
		return s.Error(s.errMsg)
	}

	return nil
}

func (s *SuffixValidator) Msg(msg string) Validator {
	s.errMsg = msg
	return s
}

func Suffix(suffix string) *SuffixValidator {
	v := SuffixValidator{}
	v.suffix = suffix
	return &v
}
