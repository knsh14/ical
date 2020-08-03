package ical

import (
	"fmt"
	"time"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// https://tools.ietf.org/html/rfc5545#section-3.8.2

// DateTimeCompleted is ...
// https://tools.ietf.org/html/rfc5545#section-3.8.2.1
type DateTimeCompleted struct {
	Parameter parameter.Container
	Value     types.DateTime
}

func (dtc *DateTimeCompleted) SetCompleted(params parameter.Container, value types.DateTime) error {
	if value != types.DateTime(time.Time{}) {
		return fmt.Errorf("")
	}
	if loc := time.Time(value).Location(); loc != time.UTC {
		return fmt.Errorf("Completed timezone must be UTC, but %s", loc)
	}
	dtc.Parameter = params
	dtc.Value = value
	return nil
}

// DateTimeEnd is ...
// https://tools.ietf.org/html/rfc5545#section-3.8.2.2
type DateTimeEnd struct {
	Parameter parameter.Container
	Value     time.Time // DateTime or Date
}

func (dte *DateTimeEnd) SetEnd(params parameter.Container, value interface{}) error {
	dateTime, isDateTime := value.(types.DateTime)
	date, isDate := value.(types.Date)
	if !(isDate || isDateTime) {
		return fmt.Errorf("value must be DateTime or Date, but %T", value)
	}
	valueParam, hasValueParam := params[parameter.TypeNameValueType]
	if hasValueParam && len(valueParam) == 1 {
		if v, ok := valueParam[0].(*parameter.ValueType); ok {
			if v.Value == "DATE-TIME" && isDate {
				return fmt.Errorf("type defined by parameter is DATE-TIME, but value type is DATE")
			}
			if v.Value == "DATE" && isDateTime {
				return fmt.Errorf("type defined by parameter is DATE, but value type is DATE-TIME")
			}
		}
	}

	refTimezone, hasRefTimezone := params[parameter.TypeNameReferenceTimezone]
	if hasRefTimezone && len(refTimezone) > 1 {
		return fmt.Errorf("too many value for TZID parameter, has %d", len(params[parameter.TypeNameReferenceTimezone]))
	}
	tz := time.UTC
	if len(refTimezone) == 1 {
		tzval := string(refTimezone[0].(*parameter.ReferenceTimezone).Value)
		t, err := time.LoadLocation(tzval)
		if err != nil {
			return fmt.Errorf("load timezone %s: %w", tzval, err)
		}
		tz = t
	}
	var t time.Time
	if isDateTime {
		t = time.Time(dateTime)
	} else if isDate {
		t = time.Time(date)
	}

	if t.Location() != tz {
		return fmt.Errorf("value[%s] and parameter[%s] different timzone", t.Location(), tz)
	}

	dte.Parameter = params
	dte.Value = t

	return nil
}

// DateTimeDue is ...
// https://tools.ietf.org/html/rfc5545#section-3.8.2.3
type DateTimeDue struct {
	Parameter parameter.Container
	Value     time.Time // DateTime or Date
}

func (dtd *DateTimeDue) SetDue(params parameter.Container, value interface{}) error {
	dateTime, isDateTime := value.(types.DateTime)
	date, isDate := value.(types.Date)
	if !(isDate || isDateTime) {
		return fmt.Errorf("value must be DateTime or Date, but %T", value)
	}
	valueParam, hasValueParam := params[parameter.TypeNameValueType]
	if hasValueParam && len(valueParam) == 1 {
		if v, ok := valueParam[0].(*parameter.ValueType); ok {
			if v.Value == "DATE-TIME" && isDate {
				return fmt.Errorf("type defined by parameter is DATE-TIME, but value type is DATE")
			}
			if v.Value == "DATE" && isDateTime {
				return fmt.Errorf("type defined by parameter is DATE, but value type is DATE-TIME")
			}
		}
	}

	refTimezone, hasRefTimezone := params[parameter.TypeNameReferenceTimezone]
	if hasRefTimezone && len(refTimezone) > 1 {
		return fmt.Errorf("too many value for TZID parameter, has %d", len(params[parameter.TypeNameReferenceTimezone]))
	}
	tz := time.UTC
	if len(refTimezone) == 1 {
		tzval := string(refTimezone[0].(*parameter.ReferenceTimezone).Value)
		t, err := time.LoadLocation(tzval)
		if err != nil {
			return fmt.Errorf("load timezone %s: %w", tzval, err)
		}
		tz = t
	}
	var t time.Time
	if isDateTime {
		t = time.Time(dateTime)
	} else if isDate {
		t = time.Time(date)
	}

	if t.Location() != tz {
		return fmt.Errorf("value[%s] and parameter[%s] different timzone", t.Location(), tz)
	}
	dtd.Parameter = params
	dtd.Value = t
	return nil
}

// DateTimeStart is ...
// https://tools.ietf.org/html/rfc5545#section-3.8.2.4
type DateTimeStart struct {
	Parameter parameter.Container
	Value     time.Time // DateTime or Date
}

func (dts *DateTimeStart) SetStart(params parameter.Container, value interface{}) error {
	dateTime, isDateTime := value.(types.DateTime)
	date, isDate := value.(types.Date)
	if !(isDate || isDateTime) {
		return fmt.Errorf("value must be DateTime or Date, but %T", value)
	}
	valueParam, hasValueParam := params[parameter.TypeNameValueType]
	if hasValueParam && len(valueParam) == 1 {
		if v, ok := valueParam[0].(*parameter.ValueType); ok {
			if v.Value == "DATE-TIME" && isDate {
				return fmt.Errorf("type defined by parameter is DATE-TIME, but value type is DATE")
			}
			if v.Value == "DATE" && isDateTime {
				return fmt.Errorf("type defined by parameter is DATE, but value type is DATE-TIME")
			}
		}
	}

	refTimezone, hasRefTimezone := params[parameter.TypeNameReferenceTimezone]
	if hasRefTimezone && len(refTimezone) > 1 {
		return fmt.Errorf("too many value for TZID parameter, has %d", len(params[parameter.TypeNameReferenceTimezone]))
	}
	tz := time.UTC
	if len(refTimezone) == 1 {
		tzval := string(refTimezone[0].(*parameter.ReferenceTimezone).Value)
		t, err := time.LoadLocation(tzval)
		if err != nil {
			return fmt.Errorf("load timezone %s: %w", tzval, err)
		}
		tz = t
	}
	var t time.Time
	if isDateTime {
		t = time.Time(dateTime)
	} else if isDate {
		t = time.Time(date)
	}

	if t.Location() != tz {
		return fmt.Errorf("value[%s] and parameter[%s] different timzone", t.Location(), tz)
	}
	dts.Parameter = params
	dts.Value = t
	return nil
}

// Duration is ...
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

// FreeBusyTime is ...
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

// TimeTransparency is ...
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
