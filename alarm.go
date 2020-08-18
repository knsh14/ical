package ical

import (
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/types"
)

// Alarm is VALARM
// https://tools.ietf.org/html/rfc5545#section-3.6.6
type Alarm interface {
	implementAlarm()
}

func NewAlarmAudio() *AlarmAudio {
	return &AlarmAudio{
		Action: &property.Action{
			Value: types.Text(property.ActionTypeAudio),
		},
	}
}

type AlarmAudio struct {
	// require
	Action  *property.Action
	Trigger *property.Trigger

	Duration    *property.Duration
	RepeatCount *property.RepeatCount
	Attachment  *property.Attachment

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (aa *AlarmAudio) implementAlarm() {}

func (aa *AlarmAudio) SetAction(params parameter.Container, value types.Text) error {
	if aa.Action != nil {
		return aa.Action.SetAction(params, value)
	}
	a := &property.Action{
		Value: types.Text(property.ActionTypeAudio),
	}
	if err := a.SetAction(params, value); err != nil {
		return err
	}
	aa.Action = a
	return nil
}
func (aa *AlarmAudio) SetTrigger(params parameter.Container, value interface{}) error {
	if aa.Trigger != nil {
		return aa.Trigger.SetTrigger(params, value)
	}
	t := &property.Trigger{
		Value: types.Text(property.ActionTypeAudio),
	}
	if err := t.SetTrigger(params, value); err != nil {
		return err
	}
	aa.Trigger = t
	return nil
}
func (aa *AlarmAudio) SetDuration(params parameter.Container, value types.Duration) error {
	if aa.Duration != nil {
		return aa.Duration.SetDuration(params, value)
	}
	d := &property.Duration{}
	if err := d.SetDuration(params, value); err != nil {
		return err
	}
	aa.Duration = d
	return nil
}
func (aa *AlarmAudio) SetRepeatCount(params parameter.Container, value types.Integer) error {
	if aa.RepeatCount != nil {
		return aa.RepeatCount.SetRepeatCount(params, value)
	}
	rc := &property.RepeatCount{}
	if err := rc.SetRepeatCount(params, value); err != nil {
		return err
	}
	aa.RepeatCount = rc
	return nil
}
func (aa *AlarmAudio) SetAttachment(params parameter.Container, value types.Attachmentable) error {
	if aa.Attachment != nil {
		return aa.Attachment.SetAttachment(params, value)
	}
	a := &property.Attachment{}
	if err := a.SetAttachment(params, value); err != nil {
		return err
	}
	aa.Attachment = a
	return nil
}

type AlarmDisplay struct {
	// require
	Action      *property.Action
	Description *property.Description
	Trigger     *property.Trigger

	Duration    *property.Duration
	RepeatCount *property.RepeatCount

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (ad *AlarmDisplay) implementAlarm() {}
func (ad *AlarmDisplay) SetAction(params parameter.Container, value types.Text) error {
	if ad.Action != nil {
		return ad.Action.SetAction(params, value)
	}
	a := &property.Action{
		Value: types.Text(property.ActionTypeAudio),
	}
	if err := a.SetAction(params, value); err != nil {
		return err
	}
	ad.Action = a
	return nil
}
func (ad *AlarmDisplay) SetDescription(params parameter.Container, value types.Text) error {
	if ad.Description != nil {
		return ad.Description.SetDescription(params, value)
	}
	d := &property.Description{}
	if err := d.SetDescription(params, value); err != nil {
		return err
	}
	ad.Description = d
	return nil
}
func (ad *AlarmDisplay) SetTrigger(params parameter.Container, value interface{}) error {
	if ad.Trigger != nil {
		return ad.Trigger.SetTrigger(params, value)
	}
	t := &property.Trigger{
		Value: types.Text(property.ActionTypeAudio),
	}
	if err := t.SetTrigger(params, value); err != nil {
		return err
	}
	ad.Trigger = t
	return nil
}
func (ad *AlarmDisplay) SetDuration(params parameter.Container, value types.Duration) error {
	if ad.Duration != nil {
		return ad.Duration.SetDuration(params, value)
	}
	d := &property.Duration{}
	if err := d.SetDuration(params, value); err != nil {
		return err
	}
	ad.Duration = d
	return nil
}
func (ad *AlarmDisplay) SetRepeatCount(params parameter.Container, value types.Integer) error {
	if ad.RepeatCount != nil {
		return ad.RepeatCount.SetRepeatCount(params, value)
	}
	rc := &property.RepeatCount{}
	if err := rc.SetRepeatCount(params, value); err != nil {
		return err
	}
	ad.RepeatCount = rc
	return nil
}

type AlarmEmail struct {
	// require
	Action      *property.Action
	Description *property.Description
	Trigger     *property.Trigger
	Summary     *property.Summary
	Attendees   []*property.Attendee

	Duration    *property.Duration
	RepeatCount *property.RepeatCount

	Attachments    []*property.Attachment
	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (ae *AlarmEmail) implementAlarm() {}
func (ae *AlarmEmail) SetAction(params parameter.Container, value types.Text) error {
	if ae.Action != nil {
		return ae.Action.SetAction(params, value)
	}
	a := &property.Action{
		Value: types.Text(property.ActionTypeAudio),
	}
	if err := a.SetAction(params, value); err != nil {
		return err
	}
	ae.Action = a
	return nil
}
func (ae *AlarmEmail) SetDescription(params parameter.Container, value types.Text) error {
	if ae.Description != nil {
		return ae.Description.SetDescription(params, value)
	}
	d := &property.Description{}
	if err := d.SetDescription(params, value); err != nil {
		return err
	}
	ae.Description = d
	return nil
}
func (ae *AlarmEmail) SetTrigger(params parameter.Container, value interface{}) error {
	if ae.Trigger != nil {
		return ae.Trigger.SetTrigger(params, value)
	}
	t := &property.Trigger{
		Value: types.Text(property.ActionTypeAudio),
	}
	if err := t.SetTrigger(params, value); err != nil {
		return err
	}
	ae.Trigger = t
	return nil
}
func (ae *AlarmEmail) SetSummary(params parameter.Container, value types.Text) error {
	if ae.Summary != nil {
		return ae.Summary.SetSummary(params, value)
	}
	s := &property.Summary{}
	if err := s.SetSummary(params, value); err != nil {
		return err
	}
	ae.Summary = s
	return nil
}
func (ae *AlarmEmail) AddAttendee(params parameter.Container, value types.CalenderUserAddress) error {
	a := &property.Attendee{}
	if err := a.SetAttendee(params, value); err != nil {
		return err
	}
	ae.Attendees = append(ae.Attendees, a)
	return nil
}
func (ae *AlarmEmail) SetDuration(params parameter.Container, value types.Duration) error {
	if ae.Duration != nil {
		return ae.Duration.SetDuration(params, value)
	}
	d := &property.Duration{}
	if err := d.SetDuration(params, value); err != nil {
		return err
	}
	ae.Duration = d
	return nil
}
func (ae *AlarmEmail) SetRepeatCount(params parameter.Container, value types.Integer) error {
	if ae.RepeatCount != nil {
		return ae.RepeatCount.SetRepeatCount(params, value)
	}
	rc := &property.RepeatCount{}
	if err := rc.SetRepeatCount(params, value); err != nil {
		return err
	}
	ae.RepeatCount = rc
	return nil
}
func (ae *AlarmEmail) AddAttachment(params parameter.Container, value types.Attachmentable) error {
	a := &property.Attachment{}
	if err := a.SetAttachment(params, value); err != nil {
		return err
	}
	ae.Attachments = append(ae.Attachments, a)
	return nil
}
