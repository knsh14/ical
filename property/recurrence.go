package property

import (
	"fmt"
	"io"
	"strings"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// Recurrence Component Properties
// https://tools.ietf.org/html/rfc5545#section-3.8.5

// ExceptionDateTimes is EXDATE
// https://tools.ietf.org/html/rfc5545#section-3.8.5.1
type ExceptionDateTimes struct {
	Parameter parameter.Container
	Values    []types.TimeValue // default is DateTime
}

func (edt *ExceptionDateTimes) Decode(w io.Writer) error {
	var s []string
	for _, v := range edt.Values {
		s = append(s, v.String())
	}
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameExceptionDateTimes, edt.Parameter.String(), strings.Join(s, ",")); err != nil {
		return err
	}
	return nil
}

func (edt *ExceptionDateTimes) Validate() error {
	// TODO
	return nil
}

func (edt *ExceptionDateTimes) SetExceptionDateTimes(params parameter.Container, values []types.TimeValue) error {

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

func NewRecurrenceDateTime(params parameter.Container, s string) (types.RecurrenceDateTimeValue, error) {
	value, ok := params[parameter.TypeNameValueType]
	if !ok {
		return nil, fmt.Errorf("no value type")
	}
	if len(value) == 1 {
		return nil, fmt.Errorf("no value type")
	}
	vt, ok := value[0].(*parameter.ValueType)
	if !ok {
		return nil, fmt.Errorf("not VALUE, but %T", value[0])
	}
	var tz string
	if tzid, ok := params[parameter.TypeNameReferenceTimezone]; ok {
		t, ok := tzid[0].(*parameter.ReferenceTimezone)
		if ok {
			tz = t.Value
		}
	}

	switch vt.Value {
	case "DATE-TIME":
		dt, err := types.NewDateTime(s, tz)
		if err != nil {
			return nil, fmt.Errorf("convert %s to DATE-TIME: %w", s, err)
		}
		return dt, nil
	case "DATE":
		d, err := types.NewDate(s)
		if err != nil {
			return nil, fmt.Errorf("convert %s to DATE: %w", s, err)
		}
		return d, nil
	case "RERIOD":
		p, err := types.NewPeriod(s)
		if err != nil {
			return nil, fmt.Errorf("convert %s to PERIOD: %w", s, err)
		}
		return p, nil
	default:
		return nil, fmt.Errorf("%s is invalid name for VALUE", vt.Value)
	}
}

// RecurrenceDateTimes is RDATE
// https://tools.ietf.org/html/rfc5545#section-3.8.5.2
type RecurrenceDateTimes struct {
	Parameter parameter.Container
	Values    []types.RecurrenceDateTimeValue // default is DateTime, Date or Period are fine.
}

func (rdt *RecurrenceDateTimes) Decode(w io.Writer) error {
	var values []string
	for _, v := range rdt.Values {
		values = append(values, v.String())
	}
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameRecurrenceDateTimes, rdt.Parameter.String(), strings.Join(values, ",")); err != nil {
		return err
	}
	return nil
}

func (rdt *RecurrenceDateTimes) Validate() error {
	// TODO
	return nil
}

func (rdt *RecurrenceDateTimes) SetRecurrenceDateTimes(params parameter.Container, values []types.RecurrenceDateTimeValue) error {
	rdt.Parameter = params
	rdt.Values = values
	return nil
}

// RecurrenceRule is RRULE
// https://tools.ietf.org/html/rfc5545#section-3.8.5.3
type RecurrenceRule struct {
	Parameter parameter.Container
	Value     types.RecurrenceRule
}

func (rr *RecurrenceRule) Decode(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameRecurrenceRule, rr.Parameter.String(), rr.Value); err != nil {
		return err
	}
	return nil
}

func (rr *RecurrenceRule) Validate() error {
	// TODO
	return nil
}

func (rrule *RecurrenceRule) SetRecurrenceRule(params parameter.Container, value types.RecurrenceRule) error {
	rrule.Parameter = params
	rrule.Value = value
	return nil
}
