package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/contentline"
)

func (p *Parser) parseCalender() (*ical.Calender, error) {
	c := &ical.Calender{}
	for l := p.getCurrentLine(); !isEndComponent(l, ical.ComponentTypeCalender) || l != nil; {
		switch l.Name {
		case "CALSCALE":
			t, err := contentline.NewText(l)
			if err != nil {
				return nil, fmt.Errorf("failed to parse value for VCALENDER.CALSCALE: %w", err)
			}
			err = c.SetCalScale(t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.CALSCALE: %w", err)
			}
		case "METHOD":
			t, err := contentline.NewText(l)
			if err != nil {
				return nil, fmt.Errorf("failed to parse value for VCALENDER.METHOD: %w", err)
			}
			err = c.SetMethod(t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.METHOD: %w", err)
			}
		case "PRODID":
			t, err := contentline.NewText(l)
			if err != nil {
				return nil, fmt.Errorf("failed to parse value for VCALENDER.PRODID: %w", err)
			}
			err = c.SetMethod(t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.METHOD: %w", err)
			}
		case "VERSION":
			t, err := contentline.NewText(l)
			if err != nil {
				return nil, fmt.Errorf("failed to parse value for VCALENDER.VERSION: %w", err)
			}
			err = c.SetVersion(t)
			if err != nil {
				return nil, fmt.Errorf("failed to set value to VCALENDER.VERSION: %w", err)
			}
		case "BEGIN":
		default:
			// if isXProp {
			// }
			// if isIANAProp {
			// }
			return nil, fmt.Errorf("no property matched,LINE:%d %v", p.CurrentIndex, l)
		}
	}
	return c, nil
}
