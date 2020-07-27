package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/types"
)

func (p *Parser) parseCalender() (*ical.Calender, error) {
	p.currentComponentType = component.ComponentTypeCalender
	c := &ical.Calender{}

	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		_, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch l.Name {
		case "CALSCALE":
			params, err := p.parseParameter(l)
			if err != nil {
				p.errors = append(p.errors, err)
			}
			t := types.NewText(l.Values[0])
			err = c.SetCalScale(params, t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.CALSCALE: %w", err)
			}
		case "METHOD":
			params, err := p.parseParameter(l)
			if err != nil {
				p.errors = append(p.errors, err)
			}
			t := types.NewText(l.Values[0])
			err = c.SetMethod(params, t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.METHOD: %w", err)
			}
		case "PRODID":
			params, err := p.parseParameter(l)
			if err != nil {
				p.errors = append(p.errors, err)
			}
			t := types.NewText(l.Values[0])
			err = c.SetMethod(params, t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.METHOD: %w", err)
			}
		case "VERSION":
			params, err := p.parseParameter(l)
			if err != nil {
				p.errors = append(p.errors, err)
			}
			t := types.NewText(l.Values[0])
			err = c.SetVersion(params, t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.VERSION: %w", err)
			}
		case "BEGIN":
		case "END":
			if p.isEndComponent(component.ComponentTypeCalender) {
				return c, nil
			}
		default:
			// if isXProp {
			// }
			// if isIANAProp {
			// }
			return nil, fmt.Errorf("no property matched,LINE:%d %v", p.CurrentIndex+1, l)
		}
		p.nextLine()
	}
	return nil, fmt.Errorf("finished withour end")
}
