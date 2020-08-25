package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

func (p *Parser) parseTimezone() (*ical.Timezone, error) {
	p.nextLine() // skip BEGIN:VTIMEZONE line
	p.currentComponentType = component.TypeTimezone
	timezone := ical.NewTimezone()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameEnd:
			if !p.isEndComponent(component.TypeTimezone) {
				return nil, fmt.Errorf("Invalid END")
			}
			return timezone, nil
		case property.PropertyNameTimezoneIdentifier:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t := types.NewText(l.Values[0])
			if err := timezone.SetTimezoneID(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case property.PropertyNameLastModified:
			var tz string
			tzs := params[parameter.TypeNameReferenceTimezone]
			if len(tzs) > 0 {
				tz = tzs[0].(*parameter.ReferenceTimezone).Value
			}
			v, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("conbert value to DateTime: %w", err)
			}
			if err := timezone.SetLastModified(params, v); err != nil {
				return nil, fmt.Errorf("set value to %s: %w", pname, err)
			}
		case property.PropertyNameTimezoneURL:
			if len(l.Values) > 1 {
				return nil, fmt.Errorf("")
			}
			t, err := types.NewURI(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into URI: %w", l.Values[0], err)
			}
			if err := timezone.SetTimezoneURL(params, t); err != nil {
				return nil, fmt.Errorf("failed to set value to %s: %w", pname, err)
			}
		case property.PropertyNameBegin:
			if len(l.Values) != 1 {
				return nil, fmt.Errorf("")
			}
			switch cname := component.Type(l.Values[0]); cname {
			case component.TypeStandard:
				s, err := p.parseStandard()
				if err != nil {
					return nil, fmt.Errorf("parse Standard: %w", err)
				}
				timezone.Standards = append(timezone.Standards, s)
			case component.TypeDaylight:
				d, err := p.parseDaylight()
				if err != nil {
					return nil, fmt.Errorf("parse Daylight: %w", err)
				}
				timezone.Daylights = append(timezone.Daylights, d)
			default:
				return nil, UnknownComponentTypeError(cname)
			}
			p.currentComponentType = component.TypeTimezone
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				timezone.XProperties = append(timezone.XProperties, ns)
			}
		}
		p.nextLine()
	}
	return nil, NoEndError(component.TypeTimezone)
}

func (p *Parser) parseStandard() (*ical.Standard, error) {
	p.nextLine() // skip BEGIN line
	p.currentComponentType = component.TypeStandard
	standard := ical.NewStandard()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameEnd:
			if !p.isEndComponent(component.TypeStandard) {
				return nil, fmt.Errorf("Invalid END")
			}
			return standard, nil
		case property.PropertyNameDateTimeStart:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("parse DATE-TIME: %w", err)
			}
			if err := standard.SetStart(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameTimezoneOffsetFrom:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewUTCOffset(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse UTC offset: %w", err)
			}
			if err := standard.SetTimezoneOffsetFrom(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameTimezoneOffsetTo:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewUTCOffset(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse UTC offset: %w", err)
			}
			if err := standard.SetTimezoneOffsetTo(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameRecurrenceRule:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			rr, err := types.NewRecurrenceRule(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse recurrence rule: %w", err)
			}
			if err := standard.SetRecurrenceRule(params, rr); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameComment:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := standard.SetComment(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameRecurrenceDateTimes:
			var ts []types.RecurrenceDateTimeValue
			for _, v := range l.Values {
				t, err := property.NewRecurrenceDateTime(params, l.Values[0])
				if err != nil {
					return nil, fmt.Errorf("parse %s to DATE-TIME: %w", v, err)
				}
				ts = append(ts, t)
			}
			if err := standard.SetRecurrenceDateTimes(params, ts); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameTimezoneName:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := standard.SetTimezoneName(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				standard.XProperties = append(standard.XProperties, ns)
			}
		}
		p.nextLine()
	}
	return nil, NoEndError(component.TypeStandard)
}
func (p *Parser) parseDaylight() (*ical.Daylight, error) {
	p.nextLine() // skip BEGIN line
	p.currentComponentType = component.TypeDaylight
	daylight := ical.NewDaylight()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameEnd:
			if !p.isEndComponent(component.TypeDaylight) {
				return nil, fmt.Errorf("finished without END:%s", component.TypeDaylight)
			}
			return daylight, nil
		case property.PropertyNameDateTimeStart:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			tz := params.GetTimezone()
			t, err := types.NewDateTime(l.Values[0], tz)
			if err != nil {
				return nil, fmt.Errorf("parse DATE-TIME: %w", err)
			}
			if err := daylight.SetStart(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameTimezoneOffsetFrom:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewUTCOffset(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse UTC offset: %w", err)
			}
			if err := daylight.SetTimezoneOffsetFrom(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameTimezoneOffsetTo:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t, err := types.NewUTCOffset(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse UTC offset: %w", err)
			}
			if err := daylight.SetTimezoneOffsetTo(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameRecurrenceRule:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			rr, err := types.NewRecurrenceRule(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("parse recurrence rule: %w", err)
			}
			if err := daylight.SetRecurrenceRule(params, rr); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameComment:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := daylight.SetComment(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameRecurrenceDateTimes:
			var ts []types.RecurrenceDateTimeValue
			for _, v := range l.Values {
				t, err := property.NewRecurrenceDateTime(params, l.Values[0])
				if err != nil {
					return nil, fmt.Errorf("parse %s to DATE-TIME: %w", v, err)
				}
				ts = append(ts, t)
			}
			if err := daylight.SetRecurrenceDateTimes(params, ts); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		case property.PropertyNameTimezoneName:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			if err := daylight.SetTimezoneName(params, t); err != nil {
				return nil, NewParseError(component.TypeStandard, pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				daylight.XProperties = append(daylight.XProperties, ns)
			}
		}
		p.nextLine()
	}
	return nil, NoEndError(component.TypeDaylight)
}
