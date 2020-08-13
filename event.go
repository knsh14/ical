package ical

import (
	"fmt"

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

	Attachment          *property.Attachment
	Attendee            *property.Attendee
	Categories          *property.Categories
	Comment             *property.Comment
	Contact             *property.Contact
	ExceptionDateTimes  *property.ExceptionDateTimes
	RequestStatus       *property.RequestStatus
	RelatedTo           *property.RelatedTo
	Resources           *property.Resources
	RecurrenceDateTimes *property.RecurrenceDateTimes

	XProperties    []*property.NonStandard
	IANAProperties []*property.IANA
}

func (e *Event) implementCalender() {}

func (e *Event) Validate() error {
	if e.DateTimeEnd != nil && e.Duration != nil {
		return fmt.Errorf("DateTimeEnd and Duraion are not nil")
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

func (e *Event) SetDateTimeStart(params parameter.Container, value types.TimeType) error {
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
		return e.Status.SetStatus(params, value, component.ComponentTypeEvent)
	}
	s := &property.Status{}
	if err := s.SetStatus(params, value, component.ComponentTypeEvent); err != nil {
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

func (e *Event) SetTimeTransparency(params parameter.Container, value types.Text) error {
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

func (e *Event) SetRecurrenceID(params parameter.Container, value types.TimeType) error {
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

func (e *Event) SetDateTimeEnd(params parameter.Container, value types.TimeType) error {
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

func (e *Event) SetAttachment(params parameter.Container, value types.Attachmentable) error {
	if e.Attachment != nil {
		return e.Attachment.SetAttachment(params, value)
	}
	a := &property.Attachment{}
	if err := a.SetAttachment(params, value); err != nil {
		return err
	}
	e.Attachment = a
	return nil
}

func (e *Event) SetAttendee(params parameter.Container, value types.CalenderUserAddress) error {
	if e.Attendee != nil {
		return e.Attendee.SetAttendee(params, value)
	}
	a := &property.Attendee{}
	if err := a.SetAttendee(params, value); err != nil {
		return err
	}
	e.Attendee = a
	return nil
}

func (e *Event) SetCategories(params parameter.Container, values []types.Text) error {
	if e.Categories != nil {
		return e.Categories.SetCategories(params, values)
	}
	c := &property.Categories{}
	if err := c.SetCategories(params, values); err != nil {
		return err
	}
	e.Categories = c
	return nil
}

func (e *Event) SetComment(params parameter.Container, value types.Text) error {
	if e.Comment != nil {
		return e.Comment.SetComment(params, value)
	}
	c := &property.Comment{}
	if err := c.SetComment(params, value); err != nil {
		return err
	}
	e.Comment = c
	return nil
}

func (e *Event) SetContact(params parameter.Container, value types.Text) error {
	if e.Contact != nil {
		return e.Contact.SetContact(params, value)
	}
	c := &property.Contact{}
	if err := c.SetContact(params, value); err != nil {
		return err
	}
	e.Contact = c
	return nil
}

func (e *Event) SetExceptionDateTimes(params parameter.Container, values []types.TimeType) error {
	if e.ExceptionDateTimes != nil {
		return e.ExceptionDateTimes.SetExceptionDateTimes(params, values)
	}
	edt := &property.ExceptionDateTimes{}
	if err := edt.SetExceptionDateTimes(params, values); err != nil {
		return err
	}
	e.ExceptionDateTimes = edt
	return nil
}

func (e *Event) SetRequestStatus(params parameter.Container, value types.Text) error {
	if e.RequestStatus != nil {
		return e.RequestStatus.SetRequestStatus(params, value)
	}
	rs := &property.RequestStatus{}
	if err := rs.SetRequestStatus(params, value); err != nil {
		return err
	}
	e.RequestStatus = rs
	return nil
}

func (e *Event) SetRelatedTo(params parameter.Container, value types.Text) error {
	if e.RelatedTo != nil {
		return e.RelatedTo.SetRelatedTo(params, value)
	}
	rt := &property.RelatedTo{}
	if err := rt.SetRelatedTo(params, value); err != nil {
		return err
	}
	e.RelatedTo = rt
	return nil
}

func (e *Event) SetResources(params parameter.Container, values []types.Text) error {
	if e.Resources != nil {
		return e.Resources.SetResources(params, values)
	}
	r := &property.Resources{}
	if err := r.SetResources(params, values); err != nil {
		return err
	}
	e.Resources = r
	return nil
}

func (e *Event) SetRecurrenceDateTimes(params parameter.Container, values []types.RecurrenceDateTime) error {
	if e.RecurrenceDateTimes != nil {
		return e.RecurrenceDateTimes.SetRecurrenceDateTimes(params, values)
	}
	rdt := &property.RecurrenceDateTimes{}
	if err := rdt.SetRecurrenceDateTimes(params, values); err != nil {
		return err
	}
	e.RecurrenceDateTimes = rdt
	return nil
}
