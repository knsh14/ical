package property

import "fmt"

var (
	ErrInputIsEmpty = fmt.Errorf("Input is empty")
)

func NewValidationError(msg string) error {
	return &ValidationError{
		msg: msg,
	}
}

type ValidationError struct {
	msg string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s", v.msg)
}
