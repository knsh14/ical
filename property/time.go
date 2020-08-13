package property

import (
	"fmt"
	"time"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// https://tools.ietf.org/html/rfc5545#section-3.8.2

// DateTimeCompleted is COMPLETED
// https://tools.ietf.org/html/rfc5545#section-3.8.2.1
type DateTimeCompleted struct {
	Parameter parameter.Container
	Value     types.DateTime
}

func (dtc *DateTimeCompleted) SetCompleted(params parameter.Container, value types.DateTime) error {
	if value != types.DateTime(time.Time{}) {
		return ErrInputIsEmpty
	}
	if loc := time.Time(value).Location(); loc != time.UTC {
		return fmt.Errorf("Completed timezone must be UTC, but %s", loc)
	}
	dtc.Parameter = params
	dtc.Value = value
	return nil
}

// DateTimeEnd is DTEND
// https://tools.ietf.org/html/rfc5545#section-3.8.2.2
type DateTimeEnd struct {
	Parameter parameter.Container
	Value     types.TimeType // DateTime or Date
}

func (dte *DateTimeEnd) SetEnd(params parameter.Container, value types.TimeType) error {
	dte.Parameter = params
	dte.Value = value
	return nil
}

// DateTimeDue is DUE
// https://tools.ietf.org/html/rfc5545#section-3.8.2.3
type DateTimeDue struct {
	Parameter parameter.Container
	Value     types.TimeType // DateTime or Date
}

func (dtd *DateTimeDue) SetDue(params parameter.Container, value types.TimeType) error {
	dtd.Parameter = params
	dtd.Value = value
	return nil
}

// DateTimeStart is DTSTART
// https://tools.ietf.org/html/rfc5545#section-3.8.2.4
type DateTimeStart struct {
	Parameter parameter.Container
	Value     types.TimeType // DateTime or Date
}

func (dts *DateTimeStart) SetStart(params parameter.Container, value types.TimeType) error {
	dts.Parameter = params
	dts.Value = value
	return nil
}

// Duration is DURATION
// https://tools.ietf.org/html/rfc5545#section-3.8.2.5
type Duration struct {
	Parameter parameter.Container
	Value     types.Duration
}

func (d *Duration) SetDuration(params parameter.Container, value types.Duration) error {
	d.Parameter = params
	d.Value = value
	return nil
}

// FreeBusyTime is FREEBUSY
// https://tools.ietf.org/html/rfc5545#section-3.8.2.6
type FreeBusyTime struct {
	Parameter parameter.Container
	Values    []types.Period
}

func (fbt *FreeBusyTime) SetFreeBusyTime(params parameter.Container, values []types.Period) error {
	if len(values) == 0 {
		return fmt.Errorf("")
	}
	if len(params[parameter.TypeNameFreeBusyTimeType]) > 1 {
		return fmt.Errorf("")
	}
	fbt.Parameter = params
	fbt.Values = values
	return nil
}

// TimeTransparency is TRANSP
// https://tools.ietf.org/html/rfc5545#section-3.8.2.7
type TimeTransparency struct {
	Parameter parameter.Container
	Value     TransparencyValueType
}

// TODO: maybe fix input value type to TransparencyValueType to force value type
func (tt *TimeTransparency) SetTransparency(params parameter.Container, value types.Text) error {
	switch v := TransparencyValueType(value); v {
	case TransparencyValueTypeOpaque, TransparencyValueTypeTransparent:
		tt.Parameter = params
		tt.Value = v
		return nil
	default:
		return fmt.Errorf("")
	}
}
