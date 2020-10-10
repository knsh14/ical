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
	p.currentComponentType = component.TypeTODO
	todo := ical.NewToDo()

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
			return todo, nil
		case property.NameUID:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetUID(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
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
			if err := todo.SetDateTimeStamp(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameClass:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetClass(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameDateTimeCompleted:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("convert date time: %w", err)
			}
			if err := todo.SetDateTimeCompleted(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
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
			if err := todo.SetDateTimeCreated(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameDescription:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetDescription(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameDateTimeStart:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert date time for %s: %w", pname, err)
			}
			if err := todo.SetDateTimeStart(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameGeo:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetGeoWithText(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
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
			if err := todo.SetLastModified(params, v); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameLocation:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetLocation(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameOrganizer:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into CalenderUserAddress: %w", l.Values[0], err)
			}
			if err := todo.SetOrganizer(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NamePercentComplete:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := todo.SetPercentComplete(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NamePriority:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := todo.SetPriority(params, i); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameRecurrenceID:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into TimeType: %w", l.Values[0], err)
			}
			if err := todo.SetRecurrenceID(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameSequenceNumber:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			i, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := todo.SetSequenceNumber(params, i); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetStatus(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameSummary:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.SetSummary(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameURL:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewURI(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into URI: %w", l.Values[0], err)
			}
			if err := todo.SetURL(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameDateTimeDue:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := ical.NewTimeType(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert date time for %s: %w", pname, err)
			}
			if err := todo.SetDateTimeDue(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameDuration:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := todo.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameAttachment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := property.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := todo.AddAttachment(params, a); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameAttendee:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := todo.AddAttendee(params, a); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameCategories:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := todo.AddCategories(params, ts); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameComment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.AddComment(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameContact:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.AddContact(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
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
			if err := todo.AddExceptionDateTimes(params, ts); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameRequestStatus:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := todo.AddRequestStatus(params, t); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameResources:
			var ts []types.Text
			for _, v := range l.Values {
				ts = append(ts, types.NewText(v))
			}
			if err := todo.AddResources(params, ts); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
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
			if err := todo.AddRecurrenceDateTimes(params, rdts); err != nil {
				return nil, NewParseError(component.TypeTODO, pname, err)
			}
		case property.NameBegin:
			if !p.isBeginComponent(component.TypeAlarm) {
				return nil, fmt.Errorf("allow only BEGIN:VALARM, but %v", l)
			}
			a, err := p.parseAlarm()
			if err != nil {
				return nil, NewParseError(component.TypeEvent, pname, err)
			}
			todo.AddAlarm(a)
			p.currentComponentType = component.TypeTODO
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
	return nil, NoEndError(component.TypeEvent)
}
