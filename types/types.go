package types

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ErrEmpty = fmt.Errorf("empty")

// Binary is defined in https://tools.ietf.org/html/rfc5545#section-3.3.1
// BASE64 encoded string
type Binary struct {
	Value string
}

func (b Binary) attachmentValue() {}
func (b Binary) String() string {
	return b.Value
}

func NewBinary(v string) (Binary, error) {
	if _, err := base64.StdEncoding.DecodeString(v); err != nil {
		return Binary{}, fmt.Errorf("base64 decode: %w", err)
	}
	return Binary{Value: v}, nil
}

// Boolean is defined in https://tools.ietf.org/html/rfc5545#section-3.3.2
type Boolean bool

func (b Boolean) String() string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

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
type CalenderUserAddress struct {
	URI *url.URL
}

func (cua CalenderUserAddress) String() string {
	if cua.URI == nil {
		return ""
	}
	return cua.URI.String()
}

func NewCalenderUserAddress(v string) (CalenderUserAddress, error) {
	uri, err := url.ParseRequestURI(v)
	if err != nil {
		return CalenderUserAddress{}, fmt.Errorf("invalid CalenderUserAddress type: %w", err)
	}
	return CalenderUserAddress{URI: uri}, nil
}

// Date is defined in https://tools.ietf.org/html/rfc5545#section-3.3.4
type Date time.Time

func (d Date) timeValue()               {}
func (d Date) recurrenceDateTimeValue() {}

func (d Date) String() string {
	return time.Time(d).Format("20060102")
}

func NewDate(v string) (Date, error) {
	t, err := time.Parse("20060102", v)
	if err != nil {
		return Date{}, fmt.Errorf("parse date: %w", err)
	}
	return Date(t), nil
}

// DateTime is defined in https://tools.ietf.org/html/rfc5545#section-3.3.5
type DateTime time.Time

func (dt DateTime) timeValue()               {}
func (dt DateTime) recurrenceDateTimeValue() {}
func (dt DateTime) triggerValue()            {}

func (dt DateTime) String() string {
	v := time.Time(dt)
	if v.Location() == time.UTC {
		return v.Format("20060102T150405Z")
	}
	return v.Format("20060102T150405")
}

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
	Direction    string
	Week         int64
	Day          int64
	HourDuration time.Duration
}

func (d Duration) triggerValue() {}

func (d Duration) String() string {
	var buf bytes.Buffer
	buf.WriteString(d.Direction)
	buf.WriteString("P")
	if d.Week > 0 {
		fmt.Fprintf(&buf, "%dW", d.Week)
		return buf.String()
	}
	if d.Day > 0 {
		fmt.Fprintf(&buf, "%dD", d.Day)
	}
	if d.HourDuration == 0 {
		return buf.String()
	}
	buf.WriteString("T")
	if d.HourDuration.Hours() > 0 {
		fmt.Fprintf(&buf, "%.fH", math.Floor(d.HourDuration.Hours()))
	}
	if int64(d.HourDuration.Minutes())%60 > 0 {
		fmt.Fprintf(&buf, "%dM", int64(d.HourDuration.Minutes())%60)
	}
	if int64(d.HourDuration.Seconds())%60 > 0 {
		if d.HourDuration.Hours() == 0 {
			fmt.Fprint(&buf, "0H")
		}
		if int64(d.HourDuration.Minutes())%60 == 0 {
			fmt.Fprint(&buf, "0M")
		}
		fmt.Fprintf(&buf, "%dS", int64(d.HourDuration.Seconds())%3600)
	}
	return buf.String()
}

var (
	durationWeekRe = regexp.MustCompile(`([+-]?)P(\d+W)`)
	durationDateRe = regexp.MustCompile(`([+-]?)P(\d+D)?(T(\d+H)?(\d+M)?(\d+S)?)`)
)

func NewDuration(v string) (Duration, error) {
	var d Duration
	if res := durationWeekRe.FindAllStringSubmatch(v, -1); len(res) > 0 && len(res[0]) > 0 {
		matches := res[0]
		matches = matches[1:]
		switch matches[0] {
		case "-", "+":
			d.Direction = matches[0]
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
			d.Direction = matches[0]
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

func NewInteger(v string) (Integer, error) {
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
	Type  PeriodType // "Explicit" or "Start", "Explicit" has end time, "Start" has duration
	End   DateTime
	Range Duration
}

func (p Period) recurrenceDateTimeValue() {}
func (p Period) String() string {
	switch p.Type {
	case PeriodTypeExplicit:
		return fmt.Sprintf("%s/%s", p.Start.String(), p.End.String())
	case PeriodTypeStart:
		return fmt.Sprintf("%s/%s", p.Start.String(), p.Range.String())
	default:
		return ""
	}
}

func NewPeriod(v string) (Period, error) {
	l := strings.Split(v, "/")
	if len(l) != 2 {
		return Period{}, fmt.Errorf("input[%s] must be divided by /", v)
	}
	var p Period
	s, err := NewDateTime(l[0], "")
	if err != nil {
		return Period{}, fmt.Errorf("parse start time: %w", err)
	}
	p.Start = s

	e, eErr := NewDateTime(l[1], "")
	if eErr == nil {
		p.Type = PeriodTypeExplicit
		p.End = e
		return p, nil
	}
	d, dErr := NewDuration(l[1])
	if dErr == nil {
		p.Type = PeriodTypeStart
		p.Range = d
		return p, nil
	}

	return Period{}, fmt.Errorf("%s need to match DURATION or DATE-TIME", l[1])
}

// RecurrenceRule is defined in https://tools.ietf.org/html/rfc5545#section-3.3.10
type RecurrenceRule struct {
	Frequency  FrequencyPattern
	EndDate    TimeValue // UNTIL
	Count      int64
	Interval   int64
	BySecond   []int64
	ByMinute   []int64
	ByHour     []int64
	ByDay      []WeekDay
	ByMonthDay []int64
	ByYearDay  []int64
	ByWeekNo   []int64
	ByMonth    []int64
	BySetPos   []int64
	WeekDay    WeekDayPattern
}

func (rr RecurrenceRule) String() string {
	var res []string
	if rr.Frequency != FrequencyPatternInvalid {
		res = append(res, fmt.Sprintf("FREQ=%s", rr.Frequency))
	}
	if rr.EndDate != nil {
		res = append(res, fmt.Sprintf("UNTIL=%v", rr.EndDate))
	}
	if rr.Count > 0 {
		res = append(res, fmt.Sprintf("COUNT=%d", rr.Count))
	}
	if rr.Interval > 0 {
		res = append(res, fmt.Sprintf("INTERVAL=%d", rr.Interval))
	}
	if len(rr.BySecond) > 0 {
		s := make([]string, 0, len(rr.BySecond))
		for _, v := range rr.BySecond {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYSECOND=%s", strings.Join(s, ",")))
	}
	if len(rr.ByMinute) > 0 {
		s := make([]string, 0, len(rr.ByMinute))
		for _, v := range rr.ByMinute {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYMINUTE=%s", strings.Join(s, ",")))
	}
	if len(rr.ByHour) > 0 {
		s := make([]string, 0, len(rr.ByHour))
		for _, v := range rr.ByHour {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYHOUR=%s", strings.Join(s, ",")))
	}
	if len(rr.ByDay) > 0 {
		s := make([]string, 0, len(rr.ByDay))
		for _, v := range rr.ByDay {
			s = append(s, v.String())
		}
		res = append(res, fmt.Sprintf("BYDAY=%s", strings.Join(s, ",")))
	}
	if len(rr.ByMonthDay) > 0 {
		s := make([]string, 0, len(rr.ByMonthDay))
		for _, v := range rr.ByMonthDay {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYMONTHDAY=%s", strings.Join(s, ",")))
	}
	if len(rr.ByYearDay) > 0 {
		s := make([]string, 0, len(rr.ByYearDay))
		for _, v := range rr.ByYearDay {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYYEARDAY=%s", strings.Join(s, ",")))
	}
	if len(rr.ByWeekNo) > 0 {
		s := make([]string, 0, len(rr.ByWeekNo))
		for _, v := range rr.ByWeekNo {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYWEEKNO=%s", strings.Join(s, ",")))
	}
	if len(rr.ByMonth) > 0 {
		s := make([]string, 0, len(rr.ByMonth))
		for _, v := range rr.ByMonth {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYMONTH=%s", strings.Join(s, ",")))
	}
	if len(rr.BySetPos) > 0 {
		s := make([]string, 0, len(rr.BySetPos))
		for _, v := range rr.BySetPos {
			s = append(s, strconv.Itoa(int(v)))
		}
		res = append(res, fmt.Sprintf("BYSETPOS=%s", strings.Join(s, ",")))
	}
	if rr.WeekDay != WeekDayPatternInvalid {
		res = append(res, fmt.Sprintf("WKST=%s", rr.WeekDay))
	}

	return strings.Join(res, ";")
}

type WeekDay struct {
	Week int64
	Day  WeekDayPattern
}

func (wd WeekDay) String() string {
	if wd.Week == 0 {
		return fmt.Sprintf("%s", wd.Day)
	}
	return fmt.Sprintf("%d%s", wd.Week, wd.Day)
}

func NewRecurrenceRule(v string) (RecurrenceRule, error) {
	if v == "" {
		return RecurrenceRule{}, ErrEmpty
	}
	values := strings.Split(v, ";")
	if len(values) == 0 {
		return RecurrenceRule{}, ErrEmpty
	}
	var res RecurrenceRule
	for _, value := range values {
		kv := strings.Split(value, "=")
		if len(kv) != 2 {
			return RecurrenceRule{}, fmt.Errorf("")
		}
		switch kv[0] {
		case "FREQ":
			res.Frequency = recurrenceRuleFrequencyPattern(kv[1])
			if res.Frequency == FrequencyPatternInvalid {
				return RecurrenceRule{}, fmt.Errorf("%s is invalid Frequency pattern", kv[1])
			}
		case "WKST":
			res.WeekDay = recurrenceRuleWeekdayPattern(kv[1])
			if res.WeekDay == WeekDayPatternInvalid {
				return RecurrenceRule{}, fmt.Errorf("%s is invalid WeekDay pattern", kv[1])
			}
		case "UNTIL":
			dt, err := NewDateTime(kv[1], "")
			if err == nil {
				res.EndDate = dt
				break
			}

			d, err := NewDate(kv[1])
			if err == nil {
				res.EndDate = d
				break
			}
			return RecurrenceRule{}, fmt.Errorf("%s cant convert DATE or DATE-TIME", kv[1])
		case "COUNT":
			c, err := strconv.Atoi(kv[1])
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to Int: %w", kv[1], err)
			}
			res.Count = int64(c)
		case "INTERVAL":
			c, err := strconv.Atoi(kv[1])
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to Int: %w", kv[1], err)
			}
			res.Interval = int64(c)
		case "BYSECOND":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return 0 <= n && n <= 60
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to second list: %w", kv[1], err)
			}
			res.BySecond = nums
		case "BYMINUTE":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return 0 <= n && n <= 59
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to minute list: %w", kv[1], err)
			}
			res.ByMinute = nums
		case "BYHOUR":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return 0 <= n && n <= 23
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to hour list: %w", kv[1], err)
			}
			res.ByHour = nums
		case "BYDAY":
			w, err := getWeekDayList(kv[1])
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to by day list: %w", kv[1], err)
			}
			res.ByDay = w
		case "BYMONTHDAY":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return -31 <= n && n <= 31
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to month day list: %w", kv[1], err)
			}
			res.ByMonthDay = nums
		case "BYYEARDAY":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return -366 <= n && n <= 366 && n != 0
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to year day list: %w", kv[1], err)
			}
			res.ByYearDay = nums
		case "BYWEEKNO":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return -53 <= n && n <= 53 && n != 0
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to by week no list: %w", kv[1], err)
			}
			res.ByWeekNo = nums
		case "BYMONTH":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return 1 <= n && n <= 12
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to month list: %w", kv[1], err)
			}
			res.ByMonth = nums
		case "BYSETPOS":
			nums, err := getNumberList(kv[1], func(n int64) bool {
				return -366 <= n && n <= 366 && n != 0
			})
			if err != nil {
				return RecurrenceRule{}, fmt.Errorf("convert %s to year day list: %w", kv[1], err)
			}
			res.BySetPos = nums
		default:
		}
	}
	return res, nil
}

func recurrenceRuleFrequencyPattern(v string) FrequencyPattern {
	switch p := FrequencyPattern(v); p {
	case FrequencyPatternSecondly, FrequencyPatternMinutely, FrequencyPatternHourly, FrequencyPatternDaily, FrequencyPatternWeekly, FrequencyPatternMonthly, FrequencyPatternYearly:
		return p
	default:
		return FrequencyPatternInvalid
	}
}
func recurrenceRuleWeekdayPattern(v string) WeekDayPattern {
	switch w := WeekDayPattern(v); w {
	case WeekDayPatternSunday, WeekDayPatternMonday, WeekDayPatternTuesday, WeekDayPatternWednesday, WeekDayPatternThursday, WeekDayPatternFriday, WeekDayPatternSaturday:
		return w
	default:
		return WeekDayPatternInvalid
	}
}

func getNumberList(v string, check func(int64) bool) ([]int64, error) {
	var res []int64
	values := strings.Split(v, ",")
	if len(values) == 0 {
		return nil, fmt.Errorf("get number list: %w", ErrEmpty)
	}
	for _, v := range values {
		a, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("convert %s into int: %w", v, err)
		}
		n := int64(a)
		if !check(n) {
			return nil, fmt.Errorf("%d is out of range", n)
		}
		res = append(res, n)
	}
	return res, nil
}

var weekDayNumRe = regexp.MustCompile(`([+-]?\d{1,2})?([[:upper:]]{2})`)

func getWeekDayList(v string) ([]WeekDay, error) {
	var days []WeekDay
	values := strings.Split(v, ",")
	if len(values) == 0 {
		return nil, fmt.Errorf("get number list: %w", ErrEmpty)
	}
	for _, v := range values {
		var w WeekDay
		res := weekDayNumRe.FindAllStringSubmatch(v, -1)
		if len(res) == 0 || len(res[0]) != 3 {
			return nil, fmt.Errorf("%s is invalid pattern", v)
		}
		matches := res[0]
		if matches[1] != "" {
			a, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, fmt.Errorf("convert %s into int: %w", v, err)
			}
			w.Week = int64(a)
		}
		w.Day = recurrenceRuleWeekdayPattern(matches[2])
		if w.Day == WeekDayPatternInvalid {
			return nil, fmt.Errorf("convert %s into week day", matches[2])
		}
		days = append(days, w)
	}
	return days, nil
}

// Text is defined in https://tools.ietf.org/html/rfc5545#section-3.3.11
type Text string

func NewText(v string) Text {
	return Text(v)
}

// Time is defined in https://tools.ietf.org/html/rfc5545#section-3.3.12
type Time time.Time

func (t Time) String() string {
	v := time.Time(t)
	if v.Location() == time.UTC {
		return v.Format("150405Z")
	}
	return v.Format("150405")
}

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
	t, err = time.ParseInLocation("150405Z", v, time.UTC)
	if err == nil {
		return Time(t), nil
	}
	return Time{}, fmt.Errorf("input %s is invalid format for Time", v)
}

// URI is defined in https://tools.ietf.org/html/rfc5545#section-3.3.13
type URI struct {
	URI *url.URL
}

func (u URI) attachmentValue() {}
func (u URI) String() string {
	if u.URI == nil {
		return ""
	}
	return u.URI.String()
}

func NewURI(v string) (URI, error) {
	uri, err := url.ParseRequestURI(v)
	if err != nil {
		return URI{}, fmt.Errorf("invalid format for URI: %w", err)
	}
	return URI{URI: uri}, nil
}

// UTCOffset is defined in https://tools.ietf.org/html/rfc5545#section-3.3.14
type UTCOffset struct {
	Direction bool // true for "+"
	Hour      uint64
	Minute    uint64
	Second    uint64
}

func (utco UTCOffset) String() string {
	var buf bytes.Buffer
	if utco.Direction {
		fmt.Fprint(&buf, "+")
	} else {
		fmt.Fprint(&buf, "-")
	}
	fmt.Fprintf(&buf, "%02d%02d", utco.Hour, utco.Minute)
	if utco.Second > 0 {
		fmt.Fprintf(&buf, "%02d", utco.Second)
	}
	return buf.String()
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
	h, err := strconv.ParseUint(v[1:3], 10, 64)
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
