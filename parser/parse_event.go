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
	p.currentComponentType = component.TypeEvent
	event := ical.NewEvent()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.Name(l.Name); pname {
		case property.NameEnd:
			if !p.isEndComponent(component.TypeEvent) {
				return nil, fmt.Errorf("Invalid END")
			}
			return event, nil
		case property.NameBegin:
			if !p.isBeginComponent(component.TypeAlarm) {
				return nil, fmt.Errorf("allow only BEGIN:VALARM, but %v", l)
			}
			a, err := p.parseAlarm()
			if err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
			event.AddAlarm(a)
			p.currentComponentType = component.TypeEvent
		case property.NameUID:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetUID(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameDateTimeStamp:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeStamp(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameDateTimeStart:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert date time for %s: %w", pname, err)
			}
			if err := event.SetDateTimeStart(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameClass:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetClass(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameDateTimeCreated:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeCreated(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameDescription:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetDescription(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameGeo:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetGeoWithText(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameLastModified:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			v, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("conbert value to DateTime: %w", err)
			}
			if err := event.SetLastModified(params, v); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameLocaiton:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetLocation(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameOrganizer:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into CalenderUserAddress: %w", l.Values[0], err)
			}
			if err := event.SetOrganizer(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}

		case property.NamePriority:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := event.SetPriority(params, i); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameSequenceNumber:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := event.SetSequenceNumber(params, i); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetStatus(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameSummary:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.SetSummary(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameTimeTransparency:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := property.TransparencyValueType(l.Values[0])
			if err := event.SetTimeTransparency(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameURL:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewURI(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into URI: %w", l.Values[0], err)
			}
			if err := event.SetURL(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameRecurrenceID:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := event.SetRecurrenceID(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameRecurrenceRule:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			rr, err := types.NewRecurrenceRule(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into RecurrenceRule: %w", l.Values[0], err)
			}
			if err := event.SetRecurrenceRule(params, rr); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameDateTimeEnd:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := event.SetDateTimeEnd(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameDuration:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := event.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameAttachment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := property.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := event.AddAttachment(params, a); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}

		case property.NameAttendee:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := event.AddAttendee(params, a); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}

		case property.NameCategories:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := event.AddCategories(params, ts); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameContact:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.AddContact(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameExceptionDateTimes:
			var ts []types.TimeValue
			for _, v := range l.Values {
				t, err := ical.NewTimeType(params, v)
				if err != nil {
					return nil, fmt.Errorf("convert %s into TimeType in %s: %w", v, pname, err)
				}
				ts = append(ts, t)
			}
			if err := event.AddExceptionDateTimes(params, ts); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameRequestStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.AddRequestStatus(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameRelatedTo:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := event.AddRelatedTo(params, t); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameResources:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := event.AddResources(params, ts); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
		case property.NameRecurrenceDateTimes:
			var rdts []types.RecurrenceDateTimeValue
			for _, v := range l.Values {
				rdt, err := property.NewRecurrenceDateTime(params, v)
				if err != nil {
					return nil, fmt.Errorf("convert %s to RecurrenceDateTime: %w", v, err)
				}
				rdts = append(rdts, rdt)
			}
			if err := event.AddRecurrenceDateTimes(params, rdts); err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
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
	return nil, NoEndError(component.TypeEvent)
}
