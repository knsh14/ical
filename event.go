package ical

import (
	"fmt"
	"io"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/types"
)

func NewEvent() *Event {
	return &Event{}
}

type Event struct {
	// required fields
	UID           *property.UID
	DateTimeStamp *property.DateTimeStamp

	// required if Calender obj dont have METHOD property
	DateTimeStart *property.DateTimeStart

	Class            *property.Class
	DateTimeCreated  *property.DateTimeCreated
	Description      *property.Description
	Geo              *property.Geo
	LastModified     *property.LastModified
	Location         *property.Location
	Organizer        *property.Organizer
	Priority         *property.Priority
	SequenceNumber   *property.SequenceNumber
	Status           *property.Status
	Summary          *property.Summary
	TimeTransparency *property.TimeTransparency
	URL              *property.URL
	RecurrenceID     *property.RecurrenceID

	// The following is OPTIONAL,
	// but SHOULD NOT occur more than once.
	RecurrenceRule *property.RecurrenceRule

	// optional, but End or Duration.
	DateTimeEnd *property.DateTimeEnd
	Duration    *property.Duration

	// optional but may occur more than once
	Attachments         []*property.Attachment
	Attendees           []*property.Attendee
	Categories          []*property.Categories
	Comments            []*property.Comment
	Contacts            []*property.Contact
	ExceptionDateTimes  []*property.ExceptionDateTimes
	RequestStatus       []*property.RequestStatus
	RelatedTos          []*property.RelatedTo
	Resources           []*property.Resources
	RecurrenceDateTimes []*property.RecurrenceDateTimes

	Alarms []Alarm

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (e *Event) implementCalender() {}

func (e *Event) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s:%s", property.NameBegin, component.TypeEvent)
	fmt.Fprintf(w, "%s:%s", property.NameEnd, component.TypeEvent)
	return nil
}

func (e *Event) Validate() error {
	if e.UID == nil {
		return NewValidationError(component.TypeEvent, property.NameUID, "must not to be nil")
	}
	if e.UID.Value == "" {
		return NewValidationError(component.TypeEvent, property.NameUID, "must not to be empty")
	}
	if e.DateTimeStamp == nil {
		return NewValidationError(component.TypeEvent, property.NameDateTimeStamp, "must not to be nil")
	}
	if e.DateTimeEnd != nil && e.Duration != nil {
		return NewValidationError(component.TypeEvent, property.NameDateTimeEnd, "one of DateTimeEnd or Duration must not be nil")
	}
	for _, alarm := range e.Alarms {
		if err := alarm.Validate(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

func (e *Event) SetUID(params parameter.Container, value types.Text) error {
	if e.UID != nil {
		return e.UID.SetUID(params, value)
	}
	uid := &property.UID{}
	if err := uid.SetUID(params, value); err != nil {
		return err
	}
	e.UID = uid
	return nil
}

func (e *Event) SetDateTimeStamp(params parameter.Container, value types.DateTime) error {
	if e.DateTimeStamp != nil {
		return e.DateTimeStamp.SetDateTimeStamp(params, value)
	}
	dts := &property.DateTimeStamp{}
	if err := dts.SetDateTimeStamp(params, value); err != nil {
		return err
	}
	e.DateTimeStamp = dts
	return nil
}

func (e *Event) SetDateTimeStart(params parameter.Container, value types.TimeValue) error {
	if e.DateTimeStart != nil {
		return e.DateTimeStart.SetStart(params, value)
	}
	dts := &property.DateTimeStart{}
	if err := dts.SetStart(params, value); err != nil {
		return err
	}
	e.DateTimeStart = dts
	return nil
}

func (e *Event) SetClass(params parameter.Container, value types.Text) error {
	if e.Class != nil {
		return e.Class.SetClass(params, value)
	}
	c := &property.Class{}
	if err := c.SetClass(params, value); err != nil {
		return err
	}
	e.Class = c
	return nil
}

func (e *Event) SetDateTimeCreated(params parameter.Container, value types.DateTime) error {
	if e.DateTimeCreated != nil {
		return e.DateTimeCreated.SetDateTimeCreated(params, value)
	}
	dtc := &property.DateTimeCreated{}
	if err := dtc.SetDateTimeCreated(params, value); err != nil {
		return err
	}
	e.DateTimeCreated = dtc
	return nil
}

func (e *Event) SetDescription(params parameter.Container, value types.Text) error {
	if e.Description != nil {
		return e.Description.SetDescription(params, value)
	}
	d := &property.Description{}
	if err := d.SetDescription(params, value); err != nil {
		return err
	}
	e.Description = d
	return nil
}

func (e *Event) SetGeo(params parameter.Container, latitude, longitude types.Float) error {
	if e.Geo != nil {
		return e.Geo.SetGeo(params, latitude, longitude)
	}
	g := &property.Geo{}
	if err := g.SetGeo(params, latitude, longitude); err != nil {
		return err
	}
	e.Geo = g
	return nil
}

func (e *Event) SetGeoWithText(params parameter.Container, value types.Text) error {
	if e.Geo != nil {
		return e.Geo.SetGeoWithText(params, value)
	}
	g := &property.Geo{}
	if err := g.SetGeoWithText(params, value); err != nil {
		return err
	}
	e.Geo = g
	return nil
}

func (e *Event) SetLastModified(params parameter.Container, value types.DateTime) error {
	if e.LastModified != nil {
		return e.LastModified.SetLastModified(params, value)
	}
	lm := &property.LastModified{}
	if err := lm.SetLastModified(params, value); err != nil {
		return err
	}
	e.LastModified = lm
	return nil
}

func (e *Event) SetLocation(params parameter.Container, value types.Text) error {
	if e.Location != nil {
		return e.Location.SetLocation(params, value)
	}
	l := &property.Location{}
	if err := l.SetLocation(params, value); err != nil {
		return err
	}
	e.Location = l
	return nil
}

func (e *Event) SetOrganizer(params parameter.Container, value types.CalenderUserAddress) error {
	if e.Organizer != nil {
		return e.Organizer.SetOrganizer(params, value)
	}
	o := &property.Organizer{}
	if err := o.SetOrganizer(params, value); err != nil {
		return err
	}
	e.Organizer = o
	return nil
}

func (e *Event) SetPriority(params parameter.Container, value types.Integer) error {
	if e.Priority != nil {
		return e.Priority.SetPriority(params, value)
	}
	p := &property.Priority{}
	if err := p.SetPriority(params, value); err != nil {
		return err
	}
	e.Priority = p
	return nil
}

func (e *Event) SetSequenceNumber(params parameter.Container, value types.Integer) error {
	if e.SequenceNumber != nil {
		return e.SequenceNumber.SetSequenceNumber(params, value)
	}
	sn := &property.SequenceNumber{}
	if err := sn.SetSequenceNumber(params, value); err != nil {
		return err
	}
	e.SequenceNumber = sn
	return nil
}

func (e *Event) SetStatus(params parameter.Container, value types.Text) error {
	if e.Status != nil {
		return e.Status.SetStatus(params, value, component.TypeEvent)
	}
	s := &property.Status{}
	if err := s.SetStatus(params, value, component.TypeEvent); err != nil {
		return err
	}
	e.Status = s
	return nil
}

func (e *Event) SetSummary(params parameter.Container, value types.Text) error {
	if e.Summary != nil {
		return e.Summary.SetSummary(params, value)
	}
	s := &property.Summary{}
	if err := s.SetSummary(params, value); err != nil {
		return err
	}
	e.Summary = s
	return nil
}

func (e *Event) SetTimeTransparency(params parameter.Container, value property.TransparencyValueType) error {
	if e.TimeTransparency != nil {
		return e.TimeTransparency.SetTransparency(params, value)
	}
	tt := &property.TimeTransparency{}
	if err := tt.SetTransparency(params, value); err != nil {
		return err
	}
	e.TimeTransparency = tt
	return nil
}

func (e *Event) SetURL(params parameter.Container, value types.URI) error {
	if e.URL != nil {
		return e.URL.SetURL(params, value)
	}
	url := &property.URL{}
	if err := url.SetURL(params, value); err != nil {
		return err
	}
	e.URL = url
	return nil
}

func (e *Event) SetRecurrenceID(params parameter.Container, value types.TimeValue) error {
	if e.RecurrenceID != nil {
		return e.RecurrenceID.SetRecurrenceID(params, value)
	}
	rid := &property.RecurrenceID{}
	if err := rid.SetRecurrenceID(params, value); err != nil {
		return err
	}
	e.RecurrenceID = rid
	return nil
}

func (e *Event) SetRecurrenceRule(params parameter.Container, value types.RecurrenceRule) error {
	if e.RecurrenceRule != nil {
		return e.RecurrenceRule.SetRecurrenceRule(params, value)
	}
	rr := &property.RecurrenceRule{}
	if err := rr.SetRecurrenceRule(params, value); err != nil {
		return err
	}
	e.RecurrenceRule = rr
	return nil
}

func (e *Event) SetDateTimeEnd(params parameter.Container, value types.TimeValue) error {
	if e.DateTimeEnd != nil {
		return e.DateTimeEnd.SetEnd(params, value)
	}
	dte := &property.DateTimeEnd{}
	if err := dte.SetEnd(params, value); err != nil {
		return err
	}
	e.DateTimeEnd = dte
	return nil
}

func (e *Event) SetDuration(params parameter.Container, value types.Duration) error {
	if e.Duration != nil {
		return e.Duration.SetDuration(params, value)
	}
	d := &property.Duration{}
	if err := d.SetDuration(params, value); err != nil {
		return err
	}
	e.Duration = d
	return nil
}

func (e *Event) AddAttachment(params parameter.Container, value types.AttachmentValue) error {
	a := &property.Attachment{}
	if err := a.SetAttachment(params, value); err != nil {
		return err
	}
	e.Attachments = append(e.Attachments, a)
	return nil
}

func (e *Event) AddAttendee(params parameter.Container, value types.CalenderUserAddress) error {
	a := &property.Attendee{}
	if err := a.SetAttendee(params, value); err != nil {
		return err
	}
	e.Attendees = append(e.Attendees, a)
	return nil
}

func (e *Event) AddCategories(params parameter.Container, values []types.Text) error {
	c := &property.Categories{}
	if err := c.SetCategories(params, values); err != nil {
		return err
	}
	e.Categories = append(e.Categories, c)
	return nil
}

func (e *Event) AddComment(params parameter.Container, value types.Text) error {
	c := &property.Comment{}
	if err := c.SetComment(params, value); err != nil {
		return err
	}
	e.Comments = append(e.Comments, c)
	return nil
}

func (e *Event) AddContact(params parameter.Container, value types.Text) error {
	c := &property.Contact{}
	if err := c.SetContact(params, value); err != nil {
		return err
	}
	e.Contacts = append(e.Contacts, c)
	return nil
}

func (e *Event) AddExceptionDateTimes(params parameter.Container, values []types.TimeValue) error {
	edt := &property.ExceptionDateTimes{}
	if err := edt.SetExceptionDateTimes(params, values); err != nil {
		return err
	}
	e.ExceptionDateTimes = append(e.ExceptionDateTimes, edt)
	return nil
}

func (e *Event) AddRequestStatus(params parameter.Container, value types.Text) error {
	rs := &property.RequestStatus{}
	if err := rs.SetRequestStatus(params, value); err != nil {
		return err
	}
	e.RequestStatus = append(e.RequestStatus, rs)
	return nil
}

func (e *Event) AddRelatedTo(params parameter.Container, value types.Text) error {
	rt := &property.RelatedTo{}
	if err := rt.SetRelatedTo(params, value); err != nil {
		return err
	}
	e.RelatedTos = append(e.RelatedTos, rt)
	return nil
}

func (e *Event) AddResources(params parameter.Container, values []types.Text) error {
	r := &property.Resources{}
	if err := r.SetResources(params, values); err != nil {
		return err
	}
	e.Resources = append(e.Resources, r)
	return nil
}

func (e *Event) AddRecurrenceDateTimes(params parameter.Container, values []types.RecurrenceDateTimeValue) error {
	rdt := &property.RecurrenceDateTimes{}
	if err := rdt.SetRecurrenceDateTimes(params, values); err != nil {
		return err
	}
	e.RecurrenceDateTimes = append(e.RecurrenceDateTimes, rdt)
	return nil
}

func (e *Event) AddAlarm(a Alarm) {
	e.Alarms = append(e.Alarms, a)
}
