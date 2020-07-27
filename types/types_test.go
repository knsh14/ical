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
				Direction: true,
				Week:      7,
			},
			expectedErr: nil,
		},
		"date": {
			input: "P15DT5H0M20S",
			expected: Duration{
				Direction:    true,
				Week:         0,
				Day:          15,
				HourDuration: time.Duration(5*time.Hour + 20*time.Second),
			},
			expectedErr: nil,
		},
		"time": {
			input: "PT5H0M20S",
			expected: Duration{
				Direction:    true,
				Week:         0,
				HourDuration: time.Duration(5*time.Hour + 20*time.Second),
			},
			expectedErr: nil,
		},
		"hour": {
			input: "PT5H",
			expected: Duration{
				Direction:    true,
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
