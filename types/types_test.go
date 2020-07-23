package types

import (
	"encoding/base64"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestBinary(t *testing.T) {
	encode := func(s string) string { return base64.StdEncoding.EncodeToString([]byte(s)) }
	testcases := map[string]struct {
		input       string
		expected    Binary
		expectedErr error
	}{
		"success": {
			input:       encode("hello world"),
			expected:    Binary{Value: encode("hello world")},
			expectedErr: nil,
		},
		"failure": {
			input:       encode("hello world") + "hello",
			expected:    Binary{},
			expectedErr: fmt.Errorf("base64 decode: %w", base64.CorruptInputError(16)),
		},
	}

	for title, tt := range testcases {
		tt := tt
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			b, err := NewBinary(tt.input)
			if diff := cmp.Diff(tt.expectedErr, err, cmp.AllowUnexported()); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
			if diff := cmp.Diff(tt.expected, b); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}

func TestNewDuration(t *testing.T) {
	// t.Parallel()
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
