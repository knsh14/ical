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
					Values: []string{string(component.ComponentTypeEvent)},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeEvent)},
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
					Values: []string{string(component.ComponentTypeEvent)},
				},
			},
			expected: nil,
			assertError: func(t *testing.T, err error) {
				expected := NoEndError(component.ComponentTypeEvent)
				if !errors.Is(err, expected) {
					t.Fatalf("unexpected: %s\nactual: %s", expected, err)
				}
			},
		},
		"simple event": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeEvent)},
				},
				{
					Name:   "UID",
					Values: []string{"hello.world@kns14.dev"},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeEvent)},
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
