package parser

import (
	"fmt"

	"github.com/knsh14/ical"
	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/contentline"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/token"
	"github.com/knsh14/ical/types"
)

func (p *Parser) parseAlarm() (ical.Alarm, error) {
	p.nextLine()
	var lines []*contentline.ContentLine

	var parseFunc func([]*contentline.ContentLine) (ical.Alarm, error)
	for l := p.getCurrentLine(); l != nil; l = p.getCurrentLine() {
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameEnd:
			if !p.isEndComponent(component.ComponentTypeAlarm) {
				return nil, fmt.Errorf("Invalid END")
			}
			if parseFunc == nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, fmt.Errorf("required ACTION but not found"))
			}
			return parseFunc(lines)
		case property.PropertyNameAction:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			switch property.ActionType(l.Values[0]) {
			case property.ActionTypeAudio:
				parseFunc = p.parseAlarmAudio
			case property.ActionTypeDisplay:
				parseFunc = p.parseAlarmDisplay
			case property.ActionTypeEMail:
				parseFunc = p.parseAlarmEmail
			}
			lines = append(lines, l)
		default:
			lines = append(lines, l)
		}
		p.nextLine()
	}
	return nil, NoEndError(component.ComponentTypeAlarm)
}

func (p *Parser) parseAlarmAudio(lines []*contentline.ContentLine) (ical.Alarm, error) {

	aa := &ical.AlarmAudio{}
	for _, l := range lines {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameAction:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := aa.SetAction(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameTrigger:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := aa.SetTrigger(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameDuration:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := aa.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameRepeatCount:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			v, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := aa.SetRepeatCount(params, v); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameAttachment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := property.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := aa.SetAttachment(params, a); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				aa.XProperties = append(aa.XProperties, ns)
			}
		}
	}
	return aa, nil
}
func (p *Parser) parseAlarmDisplay(lines []*contentline.ContentLine) (ical.Alarm, error) {

	ad := &ical.AlarmDisplay{}
	for _, l := range lines {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameAction:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ad.SetAction(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameDescription:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ad.SetDescription(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameTrigger:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ad.SetTrigger(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameDuration:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := ad.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameRepeatCount:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			v, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := ad.SetRepeatCount(params, v); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				ad.XProperties = append(ad.XProperties, ns)
			}
		}
	}
	return ad, nil
}
func (p *Parser) parseAlarmEmail(lines []*contentline.ContentLine) (ical.Alarm, error) {

	ae := &ical.AlarmEmail{}
	for _, l := range lines {
		params, err := p.parseParameter(l)
		if err != nil {
			return nil, fmt.Errorf("parse parameter: %w", err)
		}
		switch pname := property.PropertyName(l.Name); pname {
		case property.PropertyNameAction:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ae.SetAction(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameDescription:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ae.SetDescription(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameTrigger:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ae.SetTrigger(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameSummary:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			t := types.Text(l.Values[0])
			if err := ae.SetSummary(params, t); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameAttendee:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			cua, err := types.NewCalenderUserAddress(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into CalenderUserAddress: %w", l.Values[0], err)
			}
			if err := ae.AddAttendee(params, cua); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameDuration:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			d, err := types.NewDuration(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Duration: %w", l.Values[0], err)
			}
			if err := ae.SetDuration(params, d); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameRepeatCount:
			if len(l.Values) != 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			v, err := types.NewInteger(l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Integer: %w", l.Values[0], err)
			}
			if err := ae.SetRepeatCount(params, v); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		case property.PropertyNameAttachment:
			if len(l.Values) > 1 {
				return nil, NewInvalidValueLengthError(1, len(l.Values))
			}
			a, err := property.NewAttachmentValue(params, l.Values[0])
			if err != nil {
				return nil, fmt.Errorf("convert %s into Attachment value: %w", l.Values[0], err)
			}
			if err := ae.AddAttachment(params, a); err != nil {
				return nil, NewParseError(component.ComponentTypeAlarm, pname, err)
			}
		default:
			if token.IsXName(l.Name) {
				ns, err := property.NewNonStandard(l.Name, params, l.Values)
				if err != nil {
					return nil, fmt.Errorf("value : %w", err)
				}
				ae.XProperties = append(ae.XProperties, ns)
			}
		}
	}
	return ae, nil
}
