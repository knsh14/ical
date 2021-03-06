package ical

import (
	"fmt"
	"io"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/types"
)

func NewTimezone() *Timezone {
	return &Timezone{TimezoneIdentifier: &property.TimezoneIdentifier{}}
}

// Timezone is VTIMEZONE
// https://tools.ietf.org/html/rfc5545#section-3.6.5
type Timezone struct {
	// required field
	TimezoneIdentifier *property.TimezoneIdentifier

	LastModified *property.LastModified
	TimezoneURL  *property.TimezoneURL

	Standards []*Standard
	Daylights []*Daylight

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (tz *Timezone) implementCalender() {}
func (tz *Timezone) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s:%s", property.NameBegin, component.TypeTimezone)
	fmt.Fprintf(w, "%s:%s", property.NameEnd, component.TypeTimezone)
	return nil
}

func (tz *Timezone) Validate() error {
	if tz.TimezoneIdentifier == nil {
		return NewValidationError(component.TypeTimezone, property.NameTimezoneIdentifier, "must not to be nil")
	}
	for _, s := range tz.Standards {
		if err := s.Validate(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	for _, d := range tz.Daylights {
		if err := d.Validate(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

func (tz *Timezone) SetTimezoneID(params parameter.Container, value types.Text) error {
	if tz.TimezoneIdentifier != nil {
		return tz.TimezoneIdentifier.SetTimezoneID(params, value)
	}
	tzid := &property.TimezoneIdentifier{}
	if err := tzid.SetTimezoneID(params, value); err != nil {
		return err
	}
	tz.TimezoneIdentifier = tzid
	return nil
}

func (tz *Timezone) SetLastModified(params parameter.Container, value types.DateTime) error {
	if tz.LastModified != nil {
		return tz.LastModified.SetLastModified(params, value)
	}
	lm := &property.LastModified{}
	if err := lm.SetLastModified(params, value); err != nil {
		return err
	}
	tz.LastModified = lm
	return nil
}

func (tz *Timezone) SetTimezoneURL(params parameter.Container, value types.URI) error {
	if tz.TimezoneURL != nil {
		return tz.TimezoneURL.SetTimezoneURL(params, value)
	}
	tzurl := &property.TimezoneURL{}
	if err := tzurl.SetTimezoneURL(params, value); err != nil {
		return err
	}
	tz.TimezoneURL = tzurl
	return nil
}

func NewStandard() *Standard {
	return &Standard{}
}

// Standard is sub component of timezone
type Standard struct {
	// required
	DateTimeStart      *property.DateTimeStart
	TimezoneOffsetFrom *property.TimezoneOffsetFrom
	TimezoneOffsetTo   *property.TimezoneOffsetTo

	// optional
	RecurrenceRule *property.RecurrenceRule

	Comment             *property.Comment
	RecurrenceDateTimes *property.RecurrenceDateTimes
	TimezoneName        *property.TimezoneName

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (s *Standard) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s:%s", property.NameBegin, component.TypeStandard)
	fmt.Fprintf(w, "%s:%s", property.NameEnd, component.TypeStandard)
	return nil
}

func (s *Standard) Validate() error {
	if s.DateTimeStart == nil {
		return NewValidationError(component.TypeStandard, property.NameDateTimeStart, "must not be nil")
	}
	if s.TimezoneOffsetFrom == nil {
		return NewValidationError(component.TypeStandard, property.NameTimezoneOffsetFrom, "must not be nil")
	}
	if s.TimezoneOffsetTo == nil {
		return NewValidationError(component.TypeStandard, property.NameTimezoneOffsetTo, "must not be nil")
	}
	return nil
}

func (s *Standard) SetStart(params parameter.Container, value types.DateTime) error {
	if s.DateTimeStart != nil {
		return s.DateTimeStart.SetStart(params, value)
	}
	dts := &property.DateTimeStart{}
	if err := dts.SetStart(params, value); err != nil {
		return err
	}
	s.DateTimeStart = dts
	return nil
}

func (s *Standard) SetTimezoneOffsetFrom(params parameter.Container, value types.UTCOffset) error {
	if s.TimezoneOffsetFrom != nil {
		return s.TimezoneOffsetFrom.SetTimezoneOffsetFrom(params, value)
	}
	tzof := &property.TimezoneOffsetFrom{}
	if err := tzof.SetTimezoneOffsetFrom(params, value); err != nil {
		return err
	}
	s.TimezoneOffsetFrom = tzof
	return nil
}

func (s *Standard) SetTimezoneOffsetTo(params parameter.Container, value types.UTCOffset) error {
	if s.TimezoneOffsetTo != nil {
		return s.TimezoneOffsetTo.SetTimezoneOffsetTo(params, value)
	}
	tzot := &property.TimezoneOffsetTo{}
	if err := tzot.SetTimezoneOffsetTo(params, value); err != nil {
		return err
	}
	s.TimezoneOffsetTo = tzot
	return nil
}

func (s *Standard) SetRecurrenceRule(params parameter.Container, value types.RecurrenceRule) error {
	if s.RecurrenceRule != nil {
		return s.RecurrenceRule.SetRecurrenceRule(params, value)
	}
	rr := &property.RecurrenceRule{}
	if err := rr.SetRecurrenceRule(params, value); err != nil {
		return err
	}
	s.RecurrenceRule = rr
	return nil
}

func (s *Standard) SetComment(params parameter.Container, value types.Text) error {
	if s.Comment != nil {
		return s.Comment.SetComment(params, value)
	}
	c := &property.Comment{}
	if err := c.SetComment(params, value); err != nil {
		return err
	}
	s.Comment = c
	return nil
}

func (s *Standard) SetRecurrenceDateTimes(params parameter.Container, values []types.RecurrenceDateTimeValue) error {
	if s.RecurrenceDateTimes != nil {
		return s.RecurrenceDateTimes.SetRecurrenceDateTimes(params, values)
	}
	rdt := &property.RecurrenceDateTimes{}
	if err := rdt.SetRecurrenceDateTimes(params, values); err != nil {
		return err
	}
	s.RecurrenceDateTimes = rdt
	return nil
}

func (s *Standard) SetTimezoneName(params parameter.Container, value types.Text) error {
	if s.TimezoneName != nil {
		return s.TimezoneName.SetTimezoneName(params, value)
	}
	tzn := &property.TimezoneName{}
	if err := tzn.SetTimezoneName(params, value); err != nil {
		return err
	}
	s.TimezoneName = tzn
	return nil
}

func NewDaylight() *Daylight {
	return &Daylight{}
}

// Daylight is sub component of timezone
type Daylight struct {
	// required
	DateTimeStart      *property.DateTimeStart
	TimezoneOffsetFrom *property.TimezoneOffsetFrom
	TimezoneOffsetTo   *property.TimezoneOffsetTo

	// optional
	RecurrenceRule *property.RecurrenceRule

	Comment             *property.Comment
	RecurrenceDateTimes *property.RecurrenceDateTimes
	TimezoneName        *property.TimezoneName

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (d *Daylight) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s:%s", property.NameBegin, component.TypeDaylight)
	fmt.Fprintf(w, "%s:%s", property.NameEnd, component.TypeDaylight)
	return nil
}

func (d *Daylight) Validate() error {
	if d.DateTimeStart == nil {
		return NewValidationError(component.TypeStandard, property.NameDateTimeStart, "must not be nil")
	}
	if d.TimezoneOffsetFrom == nil {
		return NewValidationError(component.TypeStandard, property.NameTimezoneOffsetFrom, "must not be nil")
	}
	if d.TimezoneOffsetTo == nil {
		return NewValidationError(component.TypeStandard, property.NameTimezoneOffsetTo, "must not be nil")
	}
	return nil
}

func (d *Daylight) SetStart(params parameter.Container, value types.DateTime) error {
	if d.DateTimeStart != nil {
		return d.DateTimeStart.SetStart(params, value)
	}
	dts := &property.DateTimeStart{}
	if err := dts.SetStart(params, value); err != nil {
		return err
	}
	d.DateTimeStart = dts
	return nil
}

func (d *Daylight) SetTimezoneOffsetFrom(params parameter.Container, value types.UTCOffset) error {
	if d.TimezoneOffsetFrom != nil {
		return d.TimezoneOffsetFrom.SetTimezoneOffsetFrom(params, value)
	}
	tzof := &property.TimezoneOffsetFrom{}
	if err := tzof.SetTimezoneOffsetFrom(params, value); err != nil {
		return err
	}
	d.TimezoneOffsetFrom = tzof
	return nil
}

func (d *Daylight) SetTimezoneOffsetTo(params parameter.Container, value types.UTCOffset) error {
	if d.TimezoneOffsetTo != nil {
		return d.TimezoneOffsetTo.SetTimezoneOffsetTo(params, value)
	}
	tzot := &property.TimezoneOffsetTo{}
	if err := tzot.SetTimezoneOffsetTo(params, value); err != nil {
		return err
	}
	d.TimezoneOffsetTo = tzot
	return nil
}

func (d *Daylight) SetRecurrenceRule(params parameter.Container, value types.RecurrenceRule) error {
	if d.RecurrenceRule != nil {
		return d.RecurrenceRule.SetRecurrenceRule(params, value)
	}
	rr := &property.RecurrenceRule{}
	if err := rr.SetRecurrenceRule(params, value); err != nil {
		return err
	}
	d.RecurrenceRule = rr
	return nil
}

func (d *Daylight) SetComment(params parameter.Container, value types.Text) error {
	if d.Comment != nil {
		return d.Comment.SetComment(params, value)
	}
	c := &property.Comment{}
	if err := c.SetComment(params, value); err != nil {
		return err
	}
	d.Comment = c
	return nil
}

func (d *Daylight) SetRecurrenceDateTimes(params parameter.Container, values []types.RecurrenceDateTimeValue) error {
	if d.RecurrenceDateTimes != nil {
		return d.RecurrenceDateTimes.SetRecurrenceDateTimes(params, values)
	}
	rdt := &property.RecurrenceDateTimes{}
	if err := rdt.SetRecurrenceDateTimes(params, values); err != nil {
		return err
	}
	d.RecurrenceDateTimes = rdt
	return nil
}

func (d *Daylight) SetTimezoneName(params parameter.Container, value types.Text) error {
	if d.TimezoneName != nil {
		return d.TimezoneName.SetTimezoneName(params, value)
	}
	tzn := &property.TimezoneName{}
	if err := tzn.SetTimezoneName(params, value); err != nil {
		return err
	}
	d.TimezoneName = tzn
	return nil
}
