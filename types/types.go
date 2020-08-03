package types

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Binary is defined in https://tools.ietf.org/html/rfc5545#section-3.3.1
// BASE64 encoded string
type Binary struct {
	Value string
}

func NewBinary(v string) (Binary, error) {
	if _, err := base64.StdEncoding.DecodeString(v); err != nil {
		return Binary{}, fmt.Errorf("base64 decode: %w", err)
	}
	return Binary{Value: v}, nil
}

// Boolean is defined in https://tools.ietf.org/html/rfc5545#section-3.3.2
type Boolean bool

func NewBoolean(v string) (Boolean, error) {
	if strings.ToUpper(v) == "TRUE" {
		return Boolean(true), nil
	}
	if strings.ToUpper(v) == "FALSE" {
		return Boolean(false), nil
	}
	return Boolean(false), fmt.Errorf("input[%s] is not TRUE or FALSE", v)
}

// CalenderUserAddress is defined in https://tools.ietf.org/html/rfc5545#section-3.3.3
type CalenderUserAddress *url.URL

func NewCalenderUserAddress(v string) (CalenderUserAddress, error) {
	uri, err := url.ParseRequestURI(v)
	if err != nil {
		return nil, fmt.Errorf("invalid CalenderUserAddress type: %w", err)
	}
	return CalenderUserAddress(uri), nil
}

// Date is defined in https://tools.ietf.org/html/rfc5545#section-3.3.4
type Date time.Time

func (d Date) isTime() {}

func NewDate(v string) (Date, error) {
	t, err := time.Parse("20060102", v)
	if err != nil {
		return Date{}, fmt.Errorf("parse date: %w", err)
	}
	return Date(t), nil
}

// DateTime is defined in https://tools.ietf.org/html/rfc5545#section-3.3.5
type DateTime time.Time

func (dt DateTime) isTime() {}

func NewDateTime(v, tz string) (DateTime, error) {
	loc := time.Local
	if tz != "" {
		z, err := time.LoadLocation(tz)
		if err != nil {
			return DateTime{}, fmt.Errorf("get timezone: %w", err)
		}
		loc = z
	}
	t, err := time.ParseInLocation("20060102T150405", v, loc)
	if err == nil {
		return DateTime(t), nil
	}
	t, err = time.ParseInLocation("20060102T150405Z", v, time.UTC)
	if err == nil {
		return DateTime(t), nil
	}
	return DateTime{}, fmt.Errorf("input %s is invalid format for DateTime", v)
}

// Duration is defined in https://tools.ietf.org/html/rfc5545#section-3.3.6
type Duration struct {
	Direction    bool // true for plus, false for minus
	Week         int64
	Day          int64
	HourDuration time.Duration
}

var (
	durationWeekRe = regexp.MustCompile(`([+-]?)P(\d+W)`)
	durationDateRe = regexp.MustCompile(`([+-]?)P(\d+D)?(T(\d+H)?(\d+M)?(\d+S)?)`)
)

func NewDuration(v string) (Duration, error) {
	var d Duration
	d.Direction = true
	if res := durationWeekRe.FindAllStringSubmatch(v, -1); len(res) > 0 && len(res[0]) > 0 {
		matches := res[0]
		matches = matches[1:]
		switch matches[0] {
		case "-", "+":
			d.Direction = matches[0] == "+"
		}
		n, err := getDuration(matches[1], "W")
		if err != nil {
			return Duration{}, fmt.Errorf("parse week duration: %w", err)
		}
		d.Week = n
		return d, nil
	}
	if res := durationDateRe.FindAllStringSubmatch(v, -1); len(res) > 0 && len(res[0]) > 0 {
		matches := res[0]
		matches = matches[1:]
		switch matches[0] {
		case "-", "+":
			d.Direction = matches[0] == "+"
		}
		if matches[1] != "" {
			n, err := getDuration(matches[1], "D")
			if err != nil {
				return Duration{}, fmt.Errorf("parse day duration: %w", err)
			}
			d.Day = n
		}
		duration, err := time.ParseDuration(strings.ToLower(strings.TrimPrefix(matches[2], "T")))
		if err != nil {
			return Duration{}, fmt.Errorf("parse hour to second duration: %w", err)
		}
		d.HourDuration = duration
		return d, nil
	}
	return Duration{}, fmt.Errorf("invalid format for DURATION type, see https://tools.ietf.org/html/rfc5545#section-3.3.6")
}

func getDuration(v, unit string) (int64, error) {
	if strings.HasSuffix(v, unit) {
		n, err := strconv.ParseInt(strings.TrimSuffix(v, unit), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("parse %s into duration %s: %w", v, unit, err)
		}
		return n, nil
	}
	return 0, fmt.Errorf("input[%s] dont have required suffix %s", v, unit)
}

// Float is defined in https://tools.ietf.org/html/rfc5545#section-3.3.7
type Float float64

func NewFloat(v string) (Float, error) {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, fmt.Errorf("parse input[v] to float64: %w", err)
	}
	return Float(f), nil
}

// Integer is defined in https://tools.ietf.org/html/rfc5545#section-3.3.8
type Integer int64

func NewInteget(v string) (Integer, error) {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse input[v] to int64: %w", err)
	}
	return Integer(i), nil
}

// Period is defined in https://tools.ietf.org/html/rfc5545#section-3.3.9
// DateTime/DateTime or DateTime/Duration
type Period struct {
	Start DateTime
	Type  string // "Explicit" or "Start", "Explicit" has end time, "Start" has duration
	End   DateTime
	Range Duration
}

func NewPeriod(v string) (Period, error) {
	l := strings.Split(v, "/")
	if len(l) != 2 {
		return Period{}, fmt.Errorf("input[%s] must be divide by /", v)
	}
	var p Period
	s, err := NewDateTime(l[0], "")
	if err != nil {
		return Period{}, fmt.Errorf("parse start time: %w", err)
	}
	p.Start = s

	e, eErr := NewDateTime(l[1], "")
	if eErr == nil {
		p.Type = "Explicit"
		p.End = e
		return p, nil
	}
	d, dErr := NewDuration(l[1])
	if dErr == nil {
		p.Type = "Start"
		p.Range = d
		return p, nil
	}

	return Period{}, fmt.Errorf("%s need to match DURATION or DATE-TIME", l[1])
}

// RecurrenceRule is defined in https://tools.ietf.org/html/rfc5545#section-3.3.10
type RecurrenceRule struct {
}

// Text is defined in https://tools.ietf.org/html/rfc5545#section-3.3.11
type Text string

func NewText(v string) Text {
	return Text(v)
}

// Time is defined in https://tools.ietf.org/html/rfc5545#section-3.3.12
type Time time.Time

func NewTime(v, tz string) (Time, error) {
	loc := time.Local
	if tz != "" {
		z, err := time.LoadLocation(tz)
		if err != nil {
			return Time{}, fmt.Errorf("get timezone: %w", err)
		}
		loc = z
	}
	t, err := time.ParseInLocation("150405", v, loc)
	if err == nil {
		return Time(t), nil
	}
	t, err = time.ParseInLocation("20060102T150405Z", v, time.UTC)
	if err == nil {
		return Time(t), nil
	}
	return Time{}, fmt.Errorf("input %s is invalid format for Time", v)
}

// URI is defined in https://tools.ietf.org/html/rfc5545#section-3.3.13
type URI *url.URL

func NewURI(v string) (URI, error) {
	uri, err := url.ParseRequestURI(v)
	if err != nil {
		return nil, fmt.Errorf("invalid format for URI: %w", err)
	}
	return URI(uri), nil
}

// UTCOffset is defined in https://tools.ietf.org/html/rfc5545#section-3.3.14
type UTCOffset struct {
	Direction bool // true for "+"
	Hour      uint64
	Minute    uint64
	Second    uint64
}

func NewUTCOffset(v string) (UTCOffset, error) {
	var o UTCOffset
	switch v[0] {
	case '+':
		o.Direction = true
	case '-':
		o.Direction = false
	default:
		return UTCOffset{}, fmt.Errorf("UTCOffset must start from + or -")
	}
	if len(v) < 5 {
		return UTCOffset{}, fmt.Errorf("input[%s] is too short to parse", v)
	}
	h, err := strconv.ParseUint(v[1:2], 10, 64)
	if err != nil {
		return UTCOffset{}, fmt.Errorf("parse hour offset: %w", err)
	}
	o.Hour = h
	m, err := strconv.ParseUint(v[3:4], 10, 64)
	if err != nil {
		return UTCOffset{}, fmt.Errorf("parse minute offset: %w", err)
	}
	o.Minute = m

	if len(v) == 7 {
		s, err := strconv.ParseUint(v[5:6], 10, 64)
		if err != nil {
			return UTCOffset{}, fmt.Errorf("parse second offset: %w", err)
		}
		o.Second = s
	}
	return o, nil
}
