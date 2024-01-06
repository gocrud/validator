package validator

type errorMessage struct {
	value string
}

func (e *errorMessage) errMsg(defaultMsg string) error {
	msg := defaultMsg
	if e.value != "" {
		msg = e.value
	}
	return &Error{Message: msg}
}

func (e *errorMessage) setMsg(msg string) {
	e.value = msg
}
