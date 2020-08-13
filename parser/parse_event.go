package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

func (p *Parser) parseEvent() (*ical.Event, error) {
	p.nextLine() // skip BEGIN:VEVENT line

	event := ical.NewEvent()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameEnd:
			if !p.isEndComponent(component.ComponentTypeEvent) {
				return nil, fmt.Errorf("Invalid END")
			}
			if err := event.Validate(); err != nil {
				return nil, fmt.Errorf("validation error: %w", err)
			}
			return event, nil

		case property.PropertyNameUID:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetUID(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameDateTimeStamp:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeStamp(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameDateTimeStart:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert date time for %s: %w", pname, err)
			}
			if err := event.SetDateTimeStart(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameClass:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetClass(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameDateTimeCreated:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeCreated(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameDescription:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetDescription(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameGeo:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetGeoWithText(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameLastModified:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			v, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("conbert value to DateTime: %w", err)
			}
			if err := event.SetLastModified(params, v); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameLocaiton:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetLocation(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameOrganizer:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into CalenderUserAddress: %w", l.Values[0], err)
			}
			if err := event.SetOrganizer(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}

		case property.PropertyNamePriority:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := event.SetPriority(params, i); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameSequenceNumber:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := event.SetSequenceNumber(params, i); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetStatus(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameSummary:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetSummary(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameTimeTransparency:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetTimeTransparency(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameURL:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewURI(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into URI: %w", l.Values[0], err)
			}
			if err := event.SetURL(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameRecurrenceID:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := event.SetRecurrenceID(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameRecurrenceRule:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			rr, err := types.NewRecurrenceRule(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into RecurrenceRule: %w", l.Values[0], err)
			}
			if err := event.SetRecurrenceRule(params, rr); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameDateTimeEnd:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := event.SetDateTimeEnd(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameDuration:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := event.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameAttachment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := property.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := event.SetAttachment(params, a); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}

		case property.PropertyNameAttendee:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := event.SetAttendee(params, a); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}

		case property.PropertyNameCategories:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := event.SetCategories(params, ts); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameContact:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetContact(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameExceptionDateTimes:
			var ts []types.TimeType
			for _, v := range l.Values {
				t, err := ical.NewTimeType(params, v)
				if err != nil {
					return nil, fmt.Errorf("convert %s into TimeType in %s: %w", v, pname, err)
				}
				ts = append(ts, t)
			}
			if err := event.SetExceptionDateTimes(params, ts); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameRequestStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetRequestStatus(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameRelatedTo:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetRelatedTo(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameResources:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := event.SetResources(params, ts); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		case property.PropertyNameRecurrenceDateTimes:
			var rdts []types.RecurrenceDateTime
			for _, v := range l.Values {
				rdt, err := property.NewRecurrenceDateTime(params, v)
				if err != nil {
					return nil, fmt.Errorf("convert %s to RecurrenceDateTime: %w", v, err)
				}
				rdts = append(rdts, rdt)
			}
			if err := event.SetRecurrenceDateTimes(params, rdts); err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				event.XProperties = append(event.XProperties, ns)
			}
		}
		p.nextLine()
	}
	return nil, NoEndError(component.ComponentTypeEvent)
}
