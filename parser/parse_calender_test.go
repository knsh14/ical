package parser

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

func TestParseCalender(t *testing.T) {
	t.Parallel()

	testcases := map[string]struct {
		input         []*contentline.ContentLine
		expected      *ical.Calender
		expectedError error
	}{
		"empty calender": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
			},
			expected:      &ical.Calender{},
			expectedError: nil,
		},
		"calender version": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
				{
					Name:   "VERSION",
					Values: []string{"2.0"},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
			},
			expected: &ical.Calender{
				Version: struct {
					Param parameter.Container
					Max   types.Text
					Min   types.Text
				}{
					Max: types.NewText("2.0"),
				},
			},
			expectedError: nil,
		},
		"calender min max version": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
				{
					Name:   "VERSION",
					Values: []string{"1.2;2.0"},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
			},
			expected: &ical.Calender{
				Version: struct {
					Param parameter.Container
					Max   types.Text
					Min   types.Text
				}{
					Min: types.NewText("1.2"),
					Max: types.NewText("2.0"),
				},
			},
			expectedError: nil,
		},
		"calender X-TOKEN": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
				{
					Name:   "X-WR-TIMEZONE",
					Values: []string{"UTC"},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
			},
			expected: &ical.Calender{
				XProperty: map[string]struct {
					Name   string
					Param  parameter.Container
					Values []types.Text
				}{
					"X-WR-TIMEZONE": {
						Name:   "X-WR-TIMEZONE",
						Values: []types.Text{"UTC"},
					},
				},
			},
			expectedError: nil,
		},
		"no end": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
				{
					Name:   "VERSION",
					Values: []string{"1.2;2.0"},
				},
			},
			expected:      nil,
			expectedError: NoEndError(component.ComponentTypeCalendar),
		},
		"calender with event": {
			input: []*contentline.ContentLine{
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
				{
					Name:   "VERSION",
					Values: []string{"1.2;2.0"},
				},
				{
					Name:   "BEGIN",
					Values: []string{string(component.ComponentTypeEvent)},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeEvent)},
				},
				{
					Name:   "END",
					Values: []string{string(component.ComponentTypeCalendar)},
				},
			},
			expected: &ical.Calender{
				Component: []ical.CalenderComponent{
					&ical.Event{},
				},
				Version: struct {
					Param parameter.Container
					Max   types.Text
					Min   types.Text
				}{
					Min: types.NewText("1.2"),
					Max: types.NewText("2.0"),
				},
			},
			expectedError: nil,
		},
	}

	for title, tc := range testcases {
		tc := tc
		t.Run(title, func(t *testing.T) {
			t.Parallel()
			p := NewParser(tc.input)
			cal, err := p.parse()
			if tc.expectedError == nil && err != nil {
				t.Fatal(err)
			}
			if tc.expectedError != nil && !errors.Is(err, tc.expectedError) {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.expected, cal, cmp.AllowUnexported(types.DateTime{}, types.Date{})); diff != "" {
				t.Errorf("(-got, +want)\n%s", diff)
			}
		})
	}
}
