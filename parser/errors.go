package parser

import (
	"fmt"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/property"
	"github.com/morikuni/failure"
)

const (
	Invalid failure.StringCode = "Invalid"
)

type NoEndError component.ComponentType

func (e NoEndError) Error() string {
	return fmt.Sprintf("finished without END:%s", string(e))
}

func NewParseError(cname component.ComponentType, pname property.PropertyName, e error) ParseError {
	return ParseError{
		componentName: cname,
		propertyName:  pname,
		err:           e,
	}
}

type ParseError struct {
	componentName component.ComponentType
	propertyName  property.PropertyName
	err           error
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse %s.%s: %s", e.componentName, e.propertyName, e.err)
}

type InvalidPropertyError property.PropertyName

type UnknownComponentTypeError component.ComponentType

func (e UnknownComponentTypeError) Error() string {
	return fmt.Sprintf("unknown type %s", string(e))
}

func NewInvalidValueLengthError(r, a int) InvalidValueLengthError {
	return InvalidValueLengthError{
		require: r,
		actual:  a,
	}
}

type InvalidValueLengthError struct {
	require, actual int
}

func (ivle InvalidValueLengthError) Error() string {
	return fmt.Sprintf("value length must be %d, but %d", ivle.require, ivle.actual)
}
