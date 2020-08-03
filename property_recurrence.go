package ical

import (
	"fmt"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// Recurrence Component Properties
// https://tools.ietf.org/html/rfc5545#section-3.8.5

// ExceptionDateTimes is EXDATE
// https://tools.ietf.org/html/rfc5545#section-3.8.5.1
type ExceptionDateTimes struct {
	Parameter parameter.Container
	Values    []types.TimeType // default is DateTime
}

func (edt *ExceptionDateTimes) SetExceptionDateTimes(params parameter.Container, values []types.TimeType) error {

	var isDate bool
	valueParam, hasValueParam := params[parameter.TypeNameValueType]
	if hasValueParam && len(valueParam) == 1 {
		if v, ok := valueParam[0].(*parameter.ValueType); ok {
			if v.Value == "DATE-TIME" {
				isDate = false
			}
			if v.Value == "DATE" {
				isDate = true
			}
		}
	}

	for _, v := range values {
		_, ok := v.(types.Date)
		if ok != isDate {
			return fmt.Errorf("value type is different from parameter, %s", v)
		}
	}
	edt.Parameter = params
	edt.Values = values
	return nil
}

// RecurrenceDateTimes is RDATE
// TODO implement
// https://tools.ietf.org/html/rfc5545#section-3.8.5.2
type RecurrenceDateTimes struct {
	Parameter parameter.Container
	Values    []interface{} // default is DateTime, Date or Period are fine.
}

func (rdt *RecurrenceDateTimes) SetRecurrenceDateTimes(params parameter.Container, values []interface{}) error {
	rdt.Parameter = params
	rdt.Values = values
	return nil
}

// RecurrenceRule is RRULE
type RecurrenceRule struct {
	Parameter parameter.Container
	Values    interface{} // RECUR
}

// TODO implement
func (rrule *RecurrenceRule) SetRecurrenceRule(params parameter.Container, value interface{}) error {
	return nil
}
