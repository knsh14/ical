package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/types"
)

func TestParseTimezone(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       []*contentline.ContentLine
		expected    *ical.Timezone
		assertError func(*testing.T, error)
	}{
		"empty": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeTimezone)},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeTimezone)},
				},
			},
			expected: ical.NewTimezone(),
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatal(err)
				}
			},
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			parser := NewParser(tc.input)
			tz, err := parser.parseTimezone()
			tc.assertError(t, err)
			if diff := cmp.Diff(tc.expected, tz, cmp.AllowUnexported(types.DateTime{}, types.Date{})); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}

func TestParseStandard(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       []*contentline.ContentLine
		expected    *ical.Standard
		assertError func(*testing.T, error)
	}{
		"empty": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeStandard)},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeStandard)},
				},
			},
			expected: &ical.Standard{},
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatal(err)
				}
			},
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			parser := NewParser(tc.input)
			std, err := parser.parseStandard()
			tc.assertError(t, err)
			if diff := cmp.Diff(tc.expected, std, cmp.AllowUnexported(types.DateTime{}, types.Date{})); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}
