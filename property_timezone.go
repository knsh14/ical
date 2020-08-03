package ical

import (
	"fmt"
	"time"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// https://tools.ietf.org/html/rfc5545#section-3.8.3

// TimezoneIdentifier is TZID
// https://tools.ietf.org/html/rfc5545#section-3.8.3.1
type TimezoneIdentifier struct {
	Parameter parameter.Container
	Value     types.Text
}

func (tzid *TimezoneIdentifier) SetTimezoneID(params parameter.Container, value types.Text) error {

	_, err := time.LoadLocation(string(value))
	if err != nil {
		return fmt.Errorf("value %s is not TimezoneID: %w", value, err)
	}
	tzid.Parameter = params
	tzid.Value = value
	return nil
}

// TimezoneName is TZNAME
// https://tools.ietf.org/html/rfc5545#section-3.8.3.2
type TimezoneName struct {
	Parameter parameter.Container
	Value     types.Text
}

func (tzn *TimezoneName) SetTimezoneName(params parameter.Container, value types.Text) error {
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("")
	}
	tzn.Parameter = params
	tzn.Value = value
	return nil
}

// TimezoneOffsetFrom is TZOFFSETFROM
// https://tools.ietf.org/html/rfc5545#section-3.8.3.3
type TimezoneOffsetFrom struct {
	Parameter parameter.Container
	Value     types.UTCOffset
}

func (tzofrom *TimezoneOffsetFrom) SetTimezoneOffsetFrom(params parameter.Container, value types.UTCOffset) error {
	tzofrom.Parameter = params
	tzofrom.Value = value
	return nil
}

// TimezoneOffsetFrom is TZOFFSETTO
// https://tools.ietf.org/html/rfc5545#section-3.8.3.4
type TimezoneOffsetTo struct {
	Parameter parameter.Container
	Value     types.UTCOffset
}

func (tzoto *TimezoneOffsetTo) SetTimezoneOffsetTo(params parameter.Container, value types.UTCOffset) error {
	tzoto.Parameter = params
	tzoto.Value = value
	return nil
}

// TimezoneURL is TZURL
// https://tools.ietf.org/html/rfc5545#section-3.8.3.5
type TimezoneURL struct {
	Parameter parameter.Container
	Value     types.URI
}

func (tzurl *TimezoneURL) SetTimezoneURL(params parameter.Container, value types.URI) error {
	tzurl.Parameter = params
	tzurl.Value = value
	return nil
}
