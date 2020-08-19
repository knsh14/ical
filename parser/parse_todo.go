package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

func (p *Parser) parseTodo() (*ical.ToDo, error) {
	p.nextLine() // skip BEGIN:VEVENT line
	p.currentComponentType = component.ComponentTypeTODO
	todo := ical.NewToDo()

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
			return todo, nil
		case property.PropertyNameUID:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetUID(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
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
			if err := todo.SetDateTimeStamp(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameClass:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetClass(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameDateTimeCompleted:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := todo.SetDateTimeCompleted(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
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
			if err := todo.SetDateTimeCreated(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameDescription:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetDescription(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameDateTimeStart:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert date time for %s: %w", pname, err)
			}
			if err := todo.SetDateTimeStart(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameGeo:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetGeoWithText(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
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
			if err := todo.SetLastModified(params, v); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameLocaiton:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetLocation(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameOrganizer:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into CalenderUserAddress: %w", l.Values[0], err)
			}
			if err := todo.SetOrganizer(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNamePercentComplete:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := todo.SetPercentComplete(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNamePriority:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := todo.SetPriority(params, i); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameRecurrenceID:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := todo.SetRecurrenceID(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameSequenceNumber:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := todo.SetSequenceNumber(params, i); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetStatus(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameSummary:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetSummary(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameURL:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewURI(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into URI: %w", l.Values[0], err)
			}
			if err := todo.SetURL(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameDateTimeDue:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert date time for %s: %w", pname, err)
			}
			if err := todo.SetDateTimeDue(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameDuration:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := todo.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameAttachment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := property.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := todo.AddAttachment(params, a); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameAttendee:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := todo.AddAttendee(params, a); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameCategories:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := todo.AddCategories(params, ts); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameComment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.AddComment(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameContact:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.AddContact(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
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
			if err := todo.AddExceptionDateTimes(params, ts); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameRequestStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.AddRequestStatus(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameResources:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := todo.AddResources(params, ts); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
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
			if err := todo.AddRecurrenceDateTimes(params, rdts); err != nil {
				return nil, NewParseError(component.ComponentTypeTODO, pname, err)
			}
		case property.PropertyNameBegin:
			if !p.isBeginComponent(component.ComponentTypeAlarm) {
				return nil, fmt.Errorf("allow only BEGIN:VALARM, but %v", l)
			}
			a, err := p.parseAlarm()
			if err != nil {
				return nil, NewParseError(component.ComponentTypeEvent, pname, err)
			}
			todo.AddAlarm(a)
			p.currentComponentType = component.ComponentTypeTODO
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				todo.XProperties = append(todo.XProperties, ns)
			}
		}
		p.nextLine()
	}
	return nil, NoEndError(component.ComponentTypeEvent)
}
