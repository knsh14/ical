package ical

import (
	"fmt"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/property"
)

func NewValidationError(c component.Type, p property.Name, msg string) error {
	return ValidationError{
		componentType: c,
		propertyName:  p,
		msg:           msg,
	}
}

type ValidationError struct {
	componentType component.Type
	propertyName  property.Name
	msg           string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Component[%s], Property[%s] %s", v.componentType, v.propertyName, v.msg)
}
