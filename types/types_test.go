package types

import (
	"encoding/base64"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestBinary(t *testing.T) {
	t.Parallel()
	encode := func(s string) string { return base64.StdEncoding.EncodeToString([]byte(s)) }
	testcases := map[string]struct {
		input       string
		expected    Binary
		assertError func(*testing.T, error)
	}{
		"success": {
			input:    encode("hello world"),
			expected: Binary{Value: encode("hello world")},
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			},
		},
		"failure": {
			input:    encode("hello world") + "hello",
			expected: Binary{},
			assertError: func(t *testing.T, err error) {
				if !errors.Is(err, base64.CorruptInputError(16)) {
					t.Fatalf("unexpected error: %v", err)
				}
			},
		},
	}

	for title, tt := range testcases {
		tt := tt
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			b, err := NewBinary(tt.input)
			tt.assertError(t, err)
			if diff := cmp.Diff(tt.expected, b); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}

func TestNewBoolean(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       string
		expected    Boolean
		assertError func(*testing.T, error)
	}{
		"TRUE_SUCCESS": {
			input:    "TRUE",
			expected: Boolean(true),
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			},
		},
		"FALSE_SUCCESS": {
			input:    "FALSE",
			expected: Boolean(false),
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			},
		},
		"TRUE_not_all_upper_SUCCESS": {
			input:    "TrUe",
			expected: Boolean(true),
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			},
		},
		"failure": {
			input:    "FAIL",
			expected: Boolean(false),
			assertError: func(t *testing.T, err error) {
				e := fmt.Errorf("input[%s] is not TRUE or FALSE", "FAIL")
				if errors.Is(err, e) {
					t.Fatalf("unexpected error: %v", err)
				}
			},
		},
	}

	for title, tt := range testcases {
		tt := tt
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			b, err := NewBoolean(tt.input)
			tt.assertError(t, err)
			if diff := cmp.Diff(tt.expected, b); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}

func TestNewDateTime(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input         string
		inputTimezone string
		assert        func(*testing.T, DateTime, error)
	}{
		"failure": {
			input: "",
			assert: func(t *testing.T, d DateTime, err error) {
			},
		},
	}

	for title, tt := range testcases {
		tt := tt
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			dt, err := NewDateTime(tt.input, tt.inputTimezone)
			tt.assert(t, dt, err)
		})
	}
}

func TestNewDuration(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       string
		expected    Duration
		expectedErr error
	}{
		"week": {
			input: "P7W",
			expected: Duration{
				Direction: "",
				Week:      7,
			},
			expectedErr: nil,
		},
		"date": {
			input: "P15DT5H0M20S",
			expected: Duration{
				Direction:    "",
				Week:         0,
				Day:          15,
				HourDuration: time.Duration(5*time.Hour + 20*time.Second),
			},
			expectedErr: nil,
		},
		"time": {
			input: "PT5H0M20S",
			expected: Duration{
				Direction:    "",
				Week:         0,
				HourDuration: time.Duration(5*time.Hour + 20*time.Second),
			},
			expectedErr: nil,
		},
		"hour": {
			input: "PT5H",
			expected: Duration{
				Direction:    "",
				Week:         0,
				Day:          0,
				HourDuration: time.Duration(5 * time.Hour),
			},
			expectedErr: nil,
		},
	}

	for title, tt := range testcases {
		tt := tt
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			d, err := NewDuration(tt.input)
			if diff := cmp.Diff(tt.expectedErr, err, cmp.AllowUnexported()); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
			if diff := cmp.Diff(tt.expected, d); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}

func TestRecurrenceRule(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       string
		expected    RecurrenceRule
		expectError func(*testing.T, error)
	}{
		"empty": {
			input:    "",
			expected: RecurrenceRule{},
			expectError: func(t *testing.T, err error) {
				if !errors.Is(err, ErrEmpty) {
					t.Fatalf("expect:%v\nactual:%v", ErrEmpty, err)
				}
			},
		},
		"FREQ": {
			input: "FREQ=DAILY",
			expected: RecurrenceRule{
				Frequency: FrequencyPatternDaily,
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"FREQ_fail": {
			input:    "FREQ=INVALID_PATTERN",
			expected: RecurrenceRule{},
			expectError: func(t *testing.T, err error) {
				expect := fmt.Errorf("%s is invalid Frequency pattern", "INVALID_PATTERN")
				if err == nil {
					t.Fatal("err expected but nil")
				}
				if err.Error() != expect.Error() {
					t.Fatalf("expect:%v\nactual:%v", expect, err)
				}
			},
		},
		"WKST": {
			input: "WKST=SU",
			expected: RecurrenceRule{
				WeekDay: WeekDayPatternSunday,
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"WKST_fail": {
			input:    "WKST=INVALID_PATTERN",
			expected: RecurrenceRule{},
			expectError: func(t *testing.T, err error) {
				expect := fmt.Errorf("%s is invalid WeekDay pattern", "INVALID_PATTERN")
				if err == nil {
					t.Fatal("err expected but nil")
				}
				if err.Error() != expect.Error() {
					t.Fatalf("expect:%v\nactual:%v", expect, err)
				}
			},
		},
		"BYWEEKNO": {
			input: "BYWEEKNO=32",
			expected: RecurrenceRule{
				ByWeekNo: []int64{32},
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"BYWEEKNO_negative": {
			input: "BYWEEKNO=-32",
			expected: RecurrenceRule{
				ByWeekNo: []int64{-32},
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"BYWEEKNO_out_of_range": {
			input:    "BYWEEKNO=0",
			expected: RecurrenceRule{},
			expectError: func(t *testing.T, err error) {
				expect := fmt.Errorf("convert %s to by week no list: %w", "0", fmt.Errorf("%d is out of range", 0))
				if err == nil {
					t.Fatal("err expected but nil")
				}
				if err.Error() != expect.Error() {
					t.Fatalf("expect:%v\nactual:%v", expect, err)
				}
			},
		},
		"BYDAY": {
			input: "BYDAY=14SU",
			expected: RecurrenceRule{
				ByDay: []WeekDay{
					{
						Week: 14,
						Day:  WeekDayPatternSunday,
					},
				},
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"BYDAY_negative": {
			input: "BYDAY=-14SU",
			expected: RecurrenceRule{
				ByDay: []WeekDay{
					{
						Week: -14,
						Day:  WeekDayPatternSunday,
					},
				},
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"BYDAY_only_weekday": {
			input: "BYDAY=SU,MO",
			expected: RecurrenceRule{
				ByDay: []WeekDay{
					{
						Day: WeekDayPatternSunday,
					},
					{
						Day: WeekDayPatternMonday,
					},
				},
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
		"BYMONTH_multi": {
			input: "BYMONTH=1,3,5,7",
			expected: RecurrenceRule{
				ByMonth: []int64{1, 3, 5, 7},
			},
			expectError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatalf("expect:%v\nactual:%v", nil, err)
				}
			},
		},
	}

	for title, tt := range testcases {
		t.Run(title, func(t *testing.T) {
			tt := tt
			t.Parallel()
			v, err := NewRecurrenceRule(tt.input)
			tt.expectError(t, err)
			if diff := cmp.Diff(tt.expected, v); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		title   string
		input   string
		convert func(string) (string, error)
	}{
		{
			title: "Binary",
			input: "AAABAAEAEBAQAAEABAAoAQAAFgAAACgAAAAQAAAAIAAAAAEABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAAAAgIAAAICAgADAwMAA////AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAMwAAAAAAABNEMQAAAAAAAkQgAAAAAAJEREQgAAACECQ0QgEgAAQxQzM0E0AABERCRCREQAADRDJEJEQwAAAhA0QwEQAAAAAEREAAAAAAAAREQAAAAAAAAkQgAAAAAAAAMgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			convert: func(s string) (string, error) {
				v, err := NewBinary(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Boolean",
			input: "TRUE",
			convert: func(s string) (string, error) {
				v, err := NewBoolean(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "CalAddress",
			input: "mailto:jane_doe@example.com",
			convert: func(s string) (string, error) {
				v, err := NewCalenderUserAddress(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Date",
			input: "19970714",
			convert: func(s string) (string, error) {
				v, err := NewDate(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "DateTime",
			input: "19980118T230000",
			convert: func(s string) (string, error) {
				v, err := NewDateTime(s, "")
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "DateTimeUTC",
			input: "19980119T070000Z",
			convert: func(s string) (string, error) {
				v, err := NewDateTime(s, "")
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "DateTimeLocalTZ",
			input: "19980119T020000",
			convert: func(s string) (string, error) {
				v, err := NewDateTime(s, "America/New_York")
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Duration_1",
			input: "P15DT5H0M20S",
			convert: func(s string) (string, error) {
				v, err := NewDuration(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Duration_2",
			input: "P7W",
			convert: func(s string) (string, error) {
				v, err := NewDuration(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Duration_3",
			input: "PT5H30M",
			convert: func(s string) (string, error) {
				v, err := NewDuration(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Period_1",
			input: "19970101T180000Z/19970102T070000Z",
			convert: func(s string) (string, error) {
				v, err := NewPeriod(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Period_2",
			input: "19970101T180000Z/PT5H30M",
			convert: func(s string) (string, error) {
				v, err := NewPeriod(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "RecurrenceRule",
			input: "FREQ=DAILY;COUNT=10;INTERVAL=2",
			convert: func(s string) (string, error) {
				v, err := NewRecurrenceRule(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "RecurrenceRule",
			input: "FREQ=MONTHLY;BYDAY=MO,TU,WE,TH,FR;BYSETPOS=-1",
			convert: func(s string) (string, error) {
				v, err := NewRecurrenceRule(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "RecurrenceRule",
			input: "FREQ=YEARLY;INTERVAL=2;BYMINUTE=30;BYHOUR=8,9;BYDAY=SU;BYMONTH=1",
			convert: func(s string) (string, error) {
				v, err := NewRecurrenceRule(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Time",
			input: "230000",
			convert: func(s string) (string, error) {
				v, err := NewTime(s, "")
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "Time",
			input: "070000Z",
			convert: func(s string) (string, error) {
				v, err := NewTime(s, "")
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "URI",
			input: "http://example.com/my-report.txt",
			convert: func(s string) (string, error) {
				v, err := NewURI(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "UTCOffset",
			input: "-0500",
			convert: func(s string) (string, error) {
				v, err := NewUTCOffset(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
		{
			title: "UTCOffset",
			input: "+0100",
			convert: func(s string) (string, error) {
				v, err := NewUTCOffset(s)
				if err != nil {
					return "", err
				}
				return v.String(), nil
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.title, func(t *testing.T) {
			res, err := tc.convert(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			if res != tc.input {
				t.Fatalf("unexpected result\ninput:\t%s\noutput:\t%s", tc.input, res)
			}
		})
	}
}
