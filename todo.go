package ical

import (
	"fmt"
	"io"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/property"
	"github.com/knsh14/ical/types"
)

func NewToDo() *ToDo {
	return &ToDo{}
}

// ToDo is VTODO
/// https://tools.ietf.org/html/rfc5545#section-3.6.2
type ToDo struct {
	// required fields
	UID           *property.UID
	DateTimeStamp *property.DateTimeStamp

	Class             *property.Class
	DateTimeCompleted *property.DateTimeCompleted
	DateTimeCreated   *property.DateTimeCreated
	Description       *property.Description
	DateTimeStart     *property.DateTimeStart
	Geo               *property.Geo
	LastModified      *property.LastModified
	Location          *property.Location
	Organizer         *property.Organizer
	PercentComplete   *property.PercentComplete
	Priority          *property.Priority
	RecurrenceID      *property.RecurrenceID
	SequenceNumber    *property.SequenceNumber
	Status            *property.Status
	Summary           *property.Summary
	URL               *property.URL

	// The following is OPTIONAL,
	// but SHOULD NOT occur more than once.
	RecurrenceRule *property.RecurrenceRule

	// optional, but Due or Duration.
	DateTimeDue *property.DateTimeDue // diff
	Duration    *property.Duration

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

func (todo *ToDo) implementCalender() {}

func (todo *ToDo) Decode(w io.Writer) error {
	fmt.Fprintf(w, "%s:%s", property.NameBegin, component.TypeTODO)
	fmt.Fprintf(w, "%s:%s", property.NameEnd, component.TypeTODO)
	return nil
}

func (todo *ToDo) Validate() error {
	if todo.UID == nil {
		return NewValidationError(component.TypeTODO, property.NameUID, "must not to be nil")
	}
	if todo.UID.Value == "" {
		return NewValidationError(component.TypeTODO, property.NameUID, "must not to be empty")
	}
	if todo.DateTimeStamp == nil {
		return NewValidationError(component.TypeTODO, property.NameDateTimeStamp, "must not to be nil")
	}
	if todo.DateTimeDue != nil && todo.Duration != nil {
		return fmt.Errorf("DateTimeEnd and Duraion are not nil")
	}
	for _, alarm := range todo.Alarms {
		if err := alarm.Validate(); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}

func (todo *ToDo) SetUID(params parameter.Container, value types.Text) error {
	if todo.UID != nil {
		return todo.UID.SetUID(params, value)
	}
	uid := &property.UID{}
	if err := uid.SetUID(params, value); err != nil {
		return err
	}
	todo.UID = uid
	return nil
}

func (todo *ToDo) SetDateTimeStamp(params parameter.Container, value types.DateTime) error {
	if todo.DateTimeStamp != nil {
		return todo.DateTimeStamp.SetDateTimeStamp(params, value)
	}
	dts := &property.DateTimeStamp{}
	if err := dts.SetDateTimeStamp(params, value); err != nil {
		return err
	}
	todo.DateTimeStamp = dts
	return nil
}

func (todo *ToDo) SetDateTimeStart(params parameter.Container, value types.TimeValue) error {
	if todo.DateTimeStart != nil {
		return todo.DateTimeStart.SetStart(params, value)
	}
	dts := &property.DateTimeStart{}
	if err := dts.SetStart(params, value); err != nil {
		return err
	}
	todo.DateTimeStart = dts
	return nil
}

func (todo *ToDo) SetClass(params parameter.Container, value types.Text) error {
	if todo.Class != nil {
		return todo.Class.SetClass(params, value)
	}
	c := &property.Class{}
	if err := c.SetClass(params, value); err != nil {
		return err
	}
	todo.Class = c
	return nil
}

func (todo *ToDo) SetDateTimeCompleted(params parameter.Container, value types.DateTime) error {
	if todo.DateTimeCompleted != nil {
		return todo.DateTimeCompleted.SetCompleted(params, value)
	}
	dtc := &property.DateTimeCompleted{}
	if err := dtc.SetCompleted(params, value); err != nil {
		return err
	}
	todo.DateTimeCompleted = dtc
	return nil
}

func (todo *ToDo) SetDateTimeCreated(params parameter.Container, value types.DateTime) error {
	if todo.DateTimeCreated != nil {
		return todo.DateTimeCreated.SetDateTimeCreated(params, value)
	}
	dtc := &property.DateTimeCreated{}
	if err := dtc.SetDateTimeCreated(params, value); err != nil {
		return err
	}
	todo.DateTimeCreated = dtc
	return nil
}

func (todo *ToDo) SetDescription(params parameter.Container, value types.Text) error {
	if todo.Description != nil {
		return todo.Description.SetDescription(params, value)
	}
	d := &property.Description{}
	if err := d.SetDescription(params, value); err != nil {
		return err
	}
	todo.Description = d
	return nil
}

func (todo *ToDo) SetGeo(params parameter.Container, latitude, longitude types.Float) error {
	if todo.Geo != nil {
		return todo.Geo.SetGeo(params, latitude, longitude)
	}
	g := &property.Geo{}
	if err := g.SetGeo(params, latitude, longitude); err != nil {
		return err
	}
	todo.Geo = g
	return nil
}

func (todo *ToDo) SetGeoWithText(params parameter.Container, value types.Text) error {
	if todo.Geo != nil {
		return todo.Geo.SetGeoWithText(params, value)
	}
	g := &property.Geo{}
	if err := g.SetGeoWithText(params, value); err != nil {
		return err
	}
	todo.Geo = g
	return nil
}

func (todo *ToDo) SetLastModified(params parameter.Container, value types.DateTime) error {
	if todo.LastModified != nil {
		return todo.LastModified.SetLastModified(params, value)
	}
	lm := &property.LastModified{}
	if err := lm.SetLastModified(params, value); err != nil {
		return err
	}
	todo.LastModified = lm
	return nil
}

func (todo *ToDo) SetLocation(params parameter.Container, value types.Text) error {
	if todo.Location != nil {
		return todo.Location.SetLocation(params, value)
	}
	l := &property.Location{}
	if err := l.SetLocation(params, value); err != nil {
		return err
	}
	todo.Location = l
	return nil
}

func (todo *ToDo) SetOrganizer(params parameter.Container, value types.CalenderUserAddress) error {
	if todo.Organizer != nil {
		return todo.Organizer.SetOrganizer(params, value)
	}
	o := &property.Organizer{}
	if err := o.SetOrganizer(params, value); err != nil {
		return err
	}
	todo.Organizer = o
	return nil
}

func (todo *ToDo) SetPercentComplete(params parameter.Container, value types.Integer) error {
	if todo.PercentComplete != nil {
		return todo.PercentComplete.SetPercentComplete(params, value)
	}
	pc := &property.PercentComplete{}
	if err := pc.SetPercentComplete(params, value); err != nil {
		return err
	}
	todo.PercentComplete = pc
	return nil
}

func (todo *ToDo) SetPriority(params parameter.Container, value types.Integer) error {
	if todo.Priority != nil {
		return todo.Priority.SetPriority(params, value)
	}
	p := &property.Priority{}
	if err := p.SetPriority(params, value); err != nil {
		return err
	}
	todo.Priority = p
	return nil
}

func (todo *ToDo) SetSequenceNumber(params parameter.Container, value types.Integer) error {
	if todo.SequenceNumber != nil {
		return todo.SequenceNumber.SetSequenceNumber(params, value)
	}
	sn := &property.SequenceNumber{}
	if err := sn.SetSequenceNumber(params, value); err != nil {
		return err
	}
	todo.SequenceNumber = sn
	return nil
}

func (todo *ToDo) SetStatus(params parameter.Container, value types.Text) error {
	if todo.Status != nil {
		return todo.Status.SetStatus(params, value, component.TypeEvent)
	}
	s := &property.Status{}
	if err := s.SetStatus(params, value, component.TypeEvent); err != nil {
		return err
	}
	todo.Status = s
	return nil
}

func (todo *ToDo) SetSummary(params parameter.Container, value types.Text) error {
	if todo.Summary != nil {
		return todo.Summary.SetSummary(params, value)
	}
	s := &property.Summary{}
	if err := s.SetSummary(params, value); err != nil {
		return err
	}
	todo.Summary = s
	return nil
}

func (todo *ToDo) SetURL(params parameter.Container, value types.URI) error {
	if todo.URL != nil {
		return todo.URL.SetURL(params, value)
	}
	url := &property.URL{}
	if err := url.SetURL(params, value); err != nil {
		return err
	}
	todo.URL = url
	return nil
}

func (todo *ToDo) SetRecurrenceID(params parameter.Container, value types.TimeValue) error {
	if todo.RecurrenceID != nil {
		return todo.RecurrenceID.SetRecurrenceID(params, value)
	}
	rid := &property.RecurrenceID{}
	if err := rid.SetRecurrenceID(params, value); err != nil {
		return err
	}
	todo.RecurrenceID = rid
	return nil
}

func (todo *ToDo) SetRecurrenceRule(params parameter.Container, value types.RecurrenceRule) error {
	if todo.RecurrenceRule != nil {
		return todo.RecurrenceRule.SetRecurrenceRule(params, value)
	}
	rr := &property.RecurrenceRule{}
	if err := rr.SetRecurrenceRule(params, value); err != nil {
		return err
	}
	todo.RecurrenceRule = rr
	return nil
}

func (todo *ToDo) SetDateTimeDue(params parameter.Container, value types.TimeValue) error {
	if todo.DateTimeDue != nil {
		return todo.DateTimeDue.SetDue(params, value)
	}
	dtd := &property.DateTimeDue{}
	if err := dtd.SetDue(params, value); err != nil {
		return err
	}
	todo.DateTimeDue = dtd
	return nil
}

func (todo *ToDo) SetDuration(params parameter.Container, value types.Duration) error {
	if todo.Duration != nil {
		return todo.Duration.SetDuration(params, value)
	}
	d := &property.Duration{}
	if err := d.SetDuration(params, value); err != nil {
		return err
	}
	todo.Duration = d
	return nil
}

func (todo *ToDo) AddAttachment(params parameter.Container, value types.AttachmentValue) error {
	a := &property.Attachment{}
	if err := a.SetAttachment(params, value); err != nil {
		return err
	}
	todo.Attachments = append(todo.Attachments, a)
	return nil
}

func (todo *ToDo) AddAttendee(params parameter.Container, value types.CalenderUserAddress) error {
	a := &property.Attendee{}
	if err := a.SetAttendee(params, value); err != nil {
		return err
	}
	todo.Attendees = append(todo.Attendees, a)
	return nil
}

func (todo *ToDo) AddCategories(params parameter.Container, values []types.Text) error {
	c := &property.Categories{}
	if err := c.SetCategories(params, values); err != nil {
		return err
	}
	todo.Categories = append(todo.Categories, c)
	return nil
}

func (todo *ToDo) AddComment(params parameter.Container, value types.Text) error {
	c := &property.Comment{}
	if err := c.SetComment(params, value); err != nil {
		return err
	}
	todo.Comments = append(todo.Comments, c)
	return nil
}

func (todo *ToDo) AddContact(params parameter.Container, value types.Text) error {
	c := &property.Contact{}
	if err := c.SetContact(params, value); err != nil {
		return err
	}
	todo.Contacts = append(todo.Contacts, c)
	return nil
}

func (todo *ToDo) AddExceptionDateTimes(params parameter.Container, values []types.TimeValue) error {
	edt := &property.ExceptionDateTimes{}
	if err := edt.SetExceptionDateTimes(params, values); err != nil {
		return err
	}
	todo.ExceptionDateTimes = append(todo.ExceptionDateTimes, edt)
	return nil
}

func (todo *ToDo) AddRequestStatus(params parameter.Container, value types.Text) error {
	rs := &property.RequestStatus{}
	if err := rs.SetRequestStatus(params, value); err != nil {
		return err
	}
	todo.RequestStatus = append(todo.RequestStatus, rs)
	return nil
}

func (todo *ToDo) AddRelatedTo(params parameter.Container, value types.Text) error {
	rt := &property.RelatedTo{}
	if err := rt.SetRelatedTo(params, value); err != nil {
		return err
	}
	todo.RelatedTos = append(todo.RelatedTos, rt)
	return nil
}

func (todo *ToDo) AddResources(params parameter.Container, values []types.Text) error {
	r := &property.Resources{}
	if err := r.SetResources(params, values); err != nil {
		return err
	}
	todo.Resources = append(todo.Resources, r)
	return nil
}

func (todo *ToDo) AddRecurrenceDateTimes(params parameter.Container, values []types.RecurrenceDateTimeValue) error {
	rdt := &property.RecurrenceDateTimes{}
	if err := rdt.SetRecurrenceDateTimes(params, values); err != nil {
		return err
	}
	todo.RecurrenceDateTimes = append(todo.RecurrenceDateTimes, rdt)
	return nil
}

func (todo *ToDo) AddAlarm(a Alarm) {
	todo.Alarms = append(todo.Alarms, a)
}
