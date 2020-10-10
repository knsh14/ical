package property

import (
	"fmt"
	"io"
	"strings"
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

func (dtc *DateTimeCompleted) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameDateTimeCompleted, dtc.Parameter.String(), dtc.Value); err != nil {
		return err
	}
	return nil
}

func (dtc *DateTimeCompleted) Validate() error {
	// TODO
	return nil
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
	Value     types.TimeValue // DateTime or Date
}

func (dte *DateTimeEnd) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameDateTimeEnd, dte.Parameter.String(), dte.Value); err != nil {
		return err
	}
	return nil
}

func (dte *DateTimeEnd) Validate() error {
	// TODO
	return nil
}

func (dte *DateTimeEnd) SetEnd(params parameter.Container, value types.TimeValue) error {
	dte.Parameter = params
	dte.Value = value
	return nil
}

// DateTimeDue is DUE
// https://tools.ietf.org/html/rfc5545#section-3.8.2.3
type DateTimeDue struct {
	Parameter parameter.Container
	Value     types.TimeValue // DateTime or Date
}

func (dtd *DateTimeDue) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameDateTimeDue, dtd.Parameter.String(), dtd.Value); err != nil {
		return err
	}
	return nil
}

func (dtd *DateTimeDue) Validate() error {
	// TODO
	return nil
}

func (dtd *DateTimeDue) SetDue(params parameter.Container, value types.TimeValue) error {
	dtd.Parameter = params
	dtd.Value = value
	return nil
}

// DateTimeStart is DTSTART
// https://tools.ietf.org/html/rfc5545#section-3.8.2.4
type DateTimeStart struct {
	Parameter parameter.Container
	Value     types.TimeValue // DateTime or Date
}

func (dts *DateTimeStart) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameDateTimeStart, dts.Parameter.String(), dts.Value); err != nil {
		return err
	}
	return nil
}

func (dts *DateTimeStart) Validate() error {
	// TODO
	return nil
}

func (dts *DateTimeStart) SetStart(params parameter.Container, value types.TimeValue) error {
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

func (d *Duration) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameDuration, d.Parameter.String(), d.Value); err != nil {
		return err
	}
	return nil
}

func (d *Duration) Validate() error {
	// TODO
	return nil
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

func (fbt *FreeBusyTime) Decoce(w io.Writer) error {
	var s []string
	for _, v := range fbt.Values {
		s = append(s, v.String())
	}
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameFreeBusyTime, fbt.Parameter.String(), strings.Join(s, ",")); err != nil {
		return err
	}
	return nil
}

func (fbt *FreeBusyTime) Validate() error {
	// TODO
	return nil
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

func (tt *TimeTransparency) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameTimeTransparency, tt.Parameter.String(), tt.Value); err != nil {
		return err
	}
	return nil
}

func (tt *TimeTransparency) Validate() error {
	// TODO
	return nil
}

func (tt *TimeTransparency) SetTransparency(params parameter.Container, value TransparencyValueType) error {
	switch value {
	case TransparencyValueTypeOpaque, TransparencyValueTypeTransparent:
		tt.Parameter = params
		tt.Value = value
		return nil
	default:
		return fmt.Errorf("unknown TransparencyValueType %s", value)
	}
}
