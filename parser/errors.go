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

type NoEndError component.Type

func (e NoEndError) Error() string {
	return fmt.Sprintf("finished without END:%s", string(e))
}

func NewParseError(cname component.Type, pname property.Name, e error) ParseError {
	return ParseError{
		componentName: cname,
		propertyName:  pname,
		err:           e,
	}
}

type ParseError struct {
	componentName component.Type
	propertyName  property.Name
	err           error
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse %s.%s: %s", e.componentName, e.propertyName, e.err)
}

type InvalidPropertyError property.Name

type UnknownComponentTypeError component.Type

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
