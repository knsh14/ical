package parser

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/types"
)

func TestParseAlarm(t *testing.T) {
	t.Parallel()
	testcases := map[string]struct {
		input       []*contentline.ContentLine
		expected    ical.Alarm
		assertError func(*testing.T, error)
	}{
		"empty": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeAlarm)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeAlarm)},
				},
			},
			expected: nil,
			assertError: func(t *testing.T, err error) {
				if !errors.As(err, &ParseError{}) {
					t.Fatal(err)
				}
			},
		},
		"audio": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeAlarm)},
				},
				{
					Name:   string(property.NameAction),
					Values: []string{string(property.ActionTypeAudio)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeAlarm)},
				},
			},
			expected: &ical.AlarmAudio{
				Action: &property.Action{
					Parameter: parameter.Container{},
					Value:     types.Text(property.ActionTypeAudio),
				},
			},
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		"display": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeAlarm)},
				},
				{
					Name:   string(property.NameAction),
					Values: []string{string(property.ActionTypeDisplay)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeAlarm)},
				},
			},
			expected: &ical.AlarmDisplay{
				Action: &property.Action{
					Parameter: parameter.Container{},
					Value:     types.Text(property.ActionTypeDisplay),
				},
			},
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		"email": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeAlarm)},
				},
				{
					Name:   string(property.NameAction),
					Values: []string{string(property.ActionTypeEMail)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeAlarm)},
				},
			},
			expected: &ical.AlarmEmail{
				Action: &property.Action{
					Parameter: parameter.Container{},
					Value:     types.Text(property.ActionTypeEMail),
				},
			},
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
			p := NewParser(tc.input)
			actual, err := p.parseAlarm()
			tc.assertError(t, err)
			if diff := cmp.Diff(tc.expected, actual, cmp.AllowUnexported(types.DateTime{}, types.Date{})); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}
