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

func TestParseEvent(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		input       []*contentline.ContentLine
		expected    *ical.Event
		assertError func(*testing.T, error)
	}{
		"empty event": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeEvent)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeEvent)},
				},
			},
			expected: &ical.Event{},
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		"no end": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeEvent)},
				},
			},
			expected: nil,
			assertError: func(t *testing.T, err error) {
				expected := NoEndError(component.TypeEvent)
				if !errors.Is(err, expected) {
					t.Fatalf("unexpected: %s\nactual: %s", expected, err)
				}
			},
		},
		"simple event": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeEvent)},
				},
				{
					Name:   "UID",
					Values: []string{"hello.world@kns14.dev"},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeEvent)},
				},
			},
			expected: &ical.Event{
				UID: &property.UID{
					Parameter: parameter.Container{},
					Value:     types.Text("hello.world@kns14.dev"),
				},
			},
			assertError: func(t *testing.T, err error) {
				if err != nil {
					t.Fatal(err)
				}
			},
		},
		"with alarm": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeEvent)},
				},
				{
					Name:   "UID",
					Values: []string{"hello.world@kns14.dev"},
				},
				{
					Name:   "BEGIN",
					Values: []string{string(component.TypeAlarm)},
				},
				{
					Name:   "ACTION",
					Values: []string{string(property.ActionTypeAudio)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeAlarm)},
				},
				{
					Name:   "END",
					Values: []string{string(component.TypeEvent)},
				},
			},
			expected: &ical.Event{
				UID: &property.UID{
					Parameter: parameter.Container{},
					Value:     types.Text("hello.world@kns14.dev"),
				},
				Alarms: []ical.Alarm{
					&ical.AlarmAudio{
						Action: &property.Action{
							Parameter: parameter.Container{},
							Value:     types.Text(property.ActionTypeAudio),
						},
					},
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
			event, err := p.parseEvent()
			tc.assertError(t, err)
			if diff := cmp.Diff(tc.expected, event, cmp.AllowUnexported(types.DateTime{}, types.Date{})); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}
