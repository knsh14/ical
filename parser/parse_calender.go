package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

func (p *Parser) parseCalender() (*ical.Calender, error) {
	p.nextLine()
	p.currentComponentType = component.TypeCalendar
	c := ical.NewCalender()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.Name(l.Name); pname {
		case property.NameCalScale:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetCalScale(params, t)
			if err != nil {
				return nil, NewParseError(component.TypeCalendar, pname, err)
			}
		case property.NameMethod:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetMethod(params, t)
			if err != nil {
				return nil, NewParseError(component.TypeCalendar, pname, err)
			}
		case property.NameProdID:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetProdID(params, t)
			if err != nil {
				return nil, NewParseError(component.TypeCalendar, pname, err)
			}
		case property.NameVersion:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetVersion(params, t)
			if err != nil {
				return nil, NewParseError(component.TypeCalendar, pname, err)
			}
		case property.NameBegin:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			switch ct := component.Type(l.Values[0]); ct {
			case component.TypeEvent:
				e, err := p.parseEvent()
				if err != nil {
					return nil, fmt.Errorf("parse %s: %w", ct, err)
				}
				c.Components = append(c.Components, e)
			case component.TypeTODO:
				todo, err := p.parseTodo()
				if err != nil {
					return nil, fmt.Errorf("parse %s: %w", ct, err)
				}
				c.Components = append(c.Components, todo)
			case component.TypeJournal:
				for !p.isEndComponent(ct) {
					p.nextLine()
				}
			case component.TypeFreeBusy:
				for !p.isEndComponent(ct) {
					p.nextLine()
				}
			case component.TypeTimezone:
				tz, err := p.parseTimezone()
				if err != nil {
					return nil, fmt.Errorf("parse %s: %w", ct, err)
				}
				c.Components = append(c.Components, tz)
			default:
				return nil, fmt.Errorf("unknown component type %s", ct)
			}
			p.currentComponentType = component.TypeCalendar
		case property.NameEnd:
			if !p.isEndComponent(component.TypeCalendar) {
				return nil, fmt.Errorf("Invalid END")
			}
			return c, nil
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				c.XProperties = append(c.XProperties, ns)
				break
			}
			// if isIANAProp {
			// }
			return nil, fmt.Errorf("no property matched,LINE:%d %v", p.CurrentIndex+1, l)
		}
		p.nextLine()
	}
	return nil, NoEndError(component.TypeCalendar)
}
