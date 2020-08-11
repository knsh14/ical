package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
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
		switch pname := ical.PropertyName(l.Name); pname {
		case ical.PropertyNameEnd:
			if !p.isEndComponent(component.ComponentTypeEvent) {
				return nil, fmt.Errorf("Invalid END")
			}
			if err := event.Validate(); err != nil {
				return nil, fmt.Errorf("validation error: %w", err)
			}
			return event, nil

		case ical.PropertyNameUID:
			t := types.NewText(l.Values[0])
			if err := event.SetUID(params, t); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameDateTimeStamp:
			var tz string
			if len(params[parameter.TypeNameReferenceTimezone]) == 1 {
				tzv, ok := params[parameter.TypeNameReferenceTimezone][0].(*parameter.ReferenceTimezone)
				if !ok {
					return nil, fmt.Errorf("not %s but %T", parameter.TypeNameReferenceTimezone, params[parameter.TypeNameReferenceTimezone][0])
				}
				tz = tzv.Value
			}
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeStamp(params, t); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameDateTimeStart:
			var tz string
			if len(params[parameter.TypeNameReferenceTimezone]) == 1 {
				tz = params[parameter.TypeNameReferenceTimezone][0].(*parameter.ReferenceTimezone).Value
			}
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeStart(params, t); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameClass:
			t := types.NewText(l.Values[0])
			if err := event.SetClass(params, t); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameDateTimeCreated:
			var tz string
			if len(params[parameter.TypeNameReferenceTimezone]) == 1 {
				tz = params[parameter.TypeNameReferenceTimezone][0].(*parameter.ReferenceTimezone).Value
			}
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := event.SetDateTimeCreated(params, t); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameDescription:
			t := types.NewText(l.Values[0])
			if err := event.SetDescription(params, t); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameGeo:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetGeoWithText(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameLastModified:
			var tz string
			tzs := params[parameter.TypeNameReferenceTimezone]
			if len(tzs) > 0 {
				tz = tzs[0].(*parameter.ReferenceTimezone).Value
			}
			v, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("conbert value to DateTime: %w", err)
			}
			if err := event.SetLastModified(params, v); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case ical.PropertyNameLocaiton:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetLocation(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameOrganizer:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into CalenderUserAddress: %w", l.Values[0], err)
			}
			if err := event.SetOrganizer(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}

		case ical.PropertyNamePriority:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := event.SetPriority(params, i); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameSequenceNumber:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := event.SetSequenceNumber(params, i); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameStatus:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetStatus(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameSummary:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetSummary(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameTimeTransparency:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetTimeTransparency(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameURL:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t, err := types.NewURI(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into URI: %w", l.Values[0], err)
			}
			if err := event.SetURL(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameRecurrenceID:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := event.SetRecurrenceID(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameRecurrenceRule:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			rr, err := types.NewRecurrenceRule(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into RecurrenceRule: %w", l.Values[0], err)
			}
			if err := event.SetRecurrenceRule(params, rr); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameDateTimeEnd:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := event.SetDateTimeEnd(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameDuration:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := event.SetDuration(params, d); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameAttachment:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			a, err := ical.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := event.SetAttachment(params, a); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}

		case ical.PropertyNameAttendee:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			a, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := event.SetAttendee(params, a); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}

		case ical.PropertyNameCategories:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := event.SetCategories(params, ts); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameContact:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetContact(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameExceptionDateTimes:
			var ts []types.TimeType
			for _, v := range l.Values {
				t, err := ical.NewTimeType(params, v)
				if err != nil {
					return nil, fmt.Errorf("convert %s into TimeType in %s: %w", v, pname, err)
				}
				ts = append(ts, t)
			}
			if err := event.SetExceptionDateTimes(params, ts); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameRequestStatus:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetRequestStatus(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameRelatedTo:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := event.SetRelatedTo(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameResources:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := event.SetResources(params, ts); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case ical.PropertyNameRecurrenceDateTimes:
			var rdts []types.RecurrenceDateTime
			for _, v := range l.Values {
				rdt, err := ical.NewRecurrenceDateTime(params, v)
				if err != nil {
					return nil, fmt.Errorf("convert %s to RecurrenceDateTime: %w", v, err)
				}
				rdts = append(rdts, rdt)
			}
			if err := event.SetRecurrenceDateTimes(params, rdts); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := ical.NewNonStandard(l.Name, params, l.Values)
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
