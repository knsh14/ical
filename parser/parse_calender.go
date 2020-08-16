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
	p.currentComponentType = component.ComponentTypeCalendar
	c := ical.NewCalender()

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameCalScale:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetCalScale(params, t)
			if err != nil {
				return nil, NewParseError(component.ComponentTypeCalendar, pname, err)
			}
		case property.PropertyNameMethod:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetMethod(params, t)
			if err != nil {
				return nil, NewParseError(component.ComponentTypeCalendar, pname, err)
			}
		case property.PropertyNameProdID:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetProdID(params, t)
			if err != nil {
				return nil, NewParseError(component.ComponentTypeCalendar, pname, err)
			}
		case property.PropertyNameVersion:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.NewText(l.Values[0])
			err = c.SetVersion(params, t)
			if err != nil {
				return nil, NewParseError(component.ComponentTypeCalendar, pname, err)
			}
		case property.PropertyNameBegin:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			switch ct := component.ComponentType(l.Values[0]); ct {
			case component.ComponentTypeEvent:
				e, err := p.parseEvent()
				if err != nil {
					return nil, fmt.Errorf("parse %s: %w", component.ComponentTypeEvent, err)
				}
				c.Component = append(c.Component, e)
				break
			case component.ComponentTypeTODO:
				todo, err := p.parseTodo()
				if err != nil {
					return nil, fmt.Errorf("parse %s: %w", component.ComponentTypeTODO, err)
				}
				c.Component = append(c.Component, todo)
				break
			case component.ComponentTypeJournal:
				for !p.isEndComponent(ct) {
					p.nextLine()
				}
			case component.ComponentTypeFreeBusy:
				for !p.isEndComponent(ct) {
					p.nextLine()
				}
			case component.ComponentTypeTimezone:
				tz, err := p.parseTimezone()
				if err != nil {
					return nil, fmt.Errorf("parse %s: %w", component.ComponentTypeTimezone, err)
				}
				c.Component = append(c.Component, tz)
				break
			default:
				return nil, fmt.Errorf("unknown component type %s", ct)
			}
		case "END":
			if !p.isEndComponent(component.ComponentTypeCalendar) {
				return nil, fmt.Errorf("Invalid END")
			}
			if err := c.Validate(); err != nil {
				return nil, fmt.Errorf("validation error: %w", err)
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
	return nil, NoEndError(component.ComponentTypeCalendar)
}
