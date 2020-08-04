package ical

import (
	"fmt"

	"github.com/knsh14/ical/component"
	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

func NewEvent() *Event {
	return &Event{UID: UID{}, DateTimeStamp: DateTimeStamp{}}
}

type Event struct {
	// required fields
	UID
	DateTimeStamp

	// required if Calender obj dont have METHOD property
	*DateTimeStart

	*Class
	*DateTimeCreated
	*Description
	*Geo
	*LastModified
	*Location
	*Organizer
	*Priority
	*SequenceNumber
	*Status
	*Summary
	*TimeTransparency
	*URL
	*RecurrenceID

	// The following is OPTIONAL,
	// but SHOULD NOT occur more than once.
	*RecurrenceRule

	// optional, but End or Duration.
	*DateTimeEnd
	*Duration

	*Attachment
	*Attendee
	*Categories
	*Comment
	*Contact
	*ExceptionDateTimes
	*RequestStatus
	*RelatedTo
	*Resources
	*RecurrenceDateTimes

	XProperties    []NonStandard
	IANAProperties []IANA
}

func (e *Event) implementCalender() {}

func (e *Event) Validate() error {
	if e.DateTimeEnd != nil && e.Duration != nil {
		return fmt.Errorf("DateTimeEnd and Duraion are not nil")
	}
	return nil
}

func (e *Event) SetDateTimeStart(params parameter.Container, value types.DateTime) error {
	if e.DateTimeStart == nil {
		e.DateTimeStart = &DateTimeStart{}
	}
	return e.DateTimeStart.SetStart(params, value)
}

func (e *Event) SetClass(params parameter.Container, value types.Text) error {
	if e.Class == nil {
		e.Class = &Class{}
	}
	return e.Class.SetClass(params, value)
}

func (e *Event) SetDateTimeCreated(params parameter.Container, value types.DateTime) error {
	if e.DateTimeCreated == nil {
		e.DateTimeCreated = &DateTimeCreated{}
	}
	return e.DateTimeCreated.SetDateTimeCreated(params, value)
}

func (e *Event) SetDescription(params parameter.Container, value types.Text) error {
	if e.Description == nil {
		e.Description = &Description{}
	}
	return e.Description.SetDescription(params, value)
}

func (e *Event) SetGeo(params parameter.Container, latitude, longitude types.Float) error {
	if e.Geo == nil {
		e.Geo = &Geo{}
	}
	return e.Geo.SetGeo(params, latitude, longitude)
}

func (e *Event) SetLastModified(params parameter.Container, value types.DateTime) error {
	if e.LastModified == nil {
		e.LastModified = &LastModified{}
	}
	return e.LastModified.SetLastModified(params, value)
}

func (e *Event) SetLocation(params parameter.Container, value types.Text) error {
	if e.Location == nil {
		e.Location = &Location{}
	}
	return e.Location.SetLocation(params, value)
}

func (e *Event) SetOrganizer(params parameter.Container, value types.CalenderUserAddress) error {
	if e.Organizer == nil {
		e.Organizer = &Organizer{}
	}
	return e.Organizer.SetOrganizer(params, value)
}

func (e *Event) SetPriority(params parameter.Container, value types.Integer) error {
	if e.Priority == nil {
		e.Priority = &Priority{}
	}
	return e.Priority.SetPriority(params, value)
}

func (e *Event) SetSequenceNumber(params parameter.Container, value types.Integer) error {
	if e.SequenceNumber == nil {
		e.SequenceNumber = &SequenceNumber{}
	}
	return e.SequenceNumber.SetSequenceNumber(params, value)
}

func (e *Event) SetStatus(params parameter.Container, value types.Text) error {
	if e.Status == nil {
		e.Status = &Status{}
	}
	return e.Status.SetStatus(params, value, component.ComponentTypeEvent)
}

func (e *Event) SetSummary(params parameter.Container, value types.Text) error {
	if e.Summary == nil {
		e.Summary = &Summary{}
	}
	return e.Summary.SetSummary(params, value)
}

func (e *Event) SetTimeTransparency(params parameter.Container, value types.Text) error {
	if e.TimeTransparency == nil {
		e.TimeTransparency = &TimeTransparency{}
	}
	return e.TimeTransparency.SetTransparency(params, value)
}

func (e *Event) SetURL(params parameter.Container, value types.URI) error {
	if e.URL == nil {
		e.URL = &URL{}
	}
	return e.URL.SetURL(params, value)
}

func (e *Event) SetRecurrenceID(params parameter.Container, value types.TimeType) error {
	if e.RecurrenceID == nil {
		e.RecurrenceID = &RecurrenceID{}
	}
	return e.RecurrenceID.SetReccuenceID(params, value)
}

func (e *Event) SetRecurrenceRule(params parameter.Container, value types.TimeType) error {
	if e.RecurrenceRule == nil {
		e.RecurrenceRule = &RecurrenceRule{}
	}
	return e.RecurrenceRule.SetRecurrenceRule(params, value)
}

func (e *Event) SetDateTimeEnd(params parameter.Container, value types.TimeType) error {
	if e.DateTimeEnd == nil {
		e.DateTimeEnd = &DateTimeEnd{}
	}
	return e.DateTimeEnd.SetEnd(params, value)
}

func (e *Event) SetDuration(params parameter.Container, value types.Duration) error {
	if e.Duration == nil {
		e.Duration = &Duration{}
	}
	return e.Duration.SetDuration(params, value)
}

func (e *Event) SetAttachment(params parameter.Container, value interface{}) error {
	if e.Attachment == nil {
		e.Attachment = &Attachment{}
	}
	return e.Attachment.SetAttachment(params, value)
}

func (e *Event) SetAttendee(params parameter.Container, value types.CalenderUserAddress) error {
	if e.Attendee == nil {
		e.Attendee = &Attendee{}
	}
	return e.Attendee.SetAttendee(params, value)
}

func (e *Event) SetCategories(params parameter.Container, values []types.Text) error {
	if e.Categories == nil {
		e.Categories = &Categories{}
	}
	return e.Categories.SetCategories(params, values)
}

func (e *Event) SetComment(params parameter.Container, value types.Text) error {
	if e.Comment == nil {
		e.Comment = &Comment{}
	}
	return e.Comment.SetComment(params, value)
}

func (e *Event) SetContact(params parameter.Container, value types.Text) error {
	if e.Contact == nil {
		e.Contact = &Contact{}
	}
	return e.Contact.SetContact(params, value)
}

func (e *Event) SetExceptionDateTimes(params parameter.Container, values []types.TimeType) error {
	if e.ExceptionDateTimes == nil {
		e.ExceptionDateTimes = &ExceptionDateTimes{}
	}
	return e.ExceptionDateTimes.SetExceptionDateTimes(params, values)
}

func (e *Event) SetRequestStatus(params parameter.Container, value types.Text) error {
	if e.RequestStatus == nil {
		e.RequestStatus = &RequestStatus{}
	}
	return e.RequestStatus.SetRequestStatus(params, value)
}

func (e *Event) SetRelatedTo(params parameter.Container, value types.Text) error {
	if e.RelatedTo == nil {
		e.RelatedTo = &RelatedTo{}
	}
	return e.RelatedTo.SetRelatedTo(params, value)
}

func (e *Event) SetResources(params parameter.Container, values []types.Text) error {
	if e.Resources == nil {
		e.Resources = &Resources{}
	}
	return e.Resources.SetResources(params, values)
}

func (e *Event) SetRecurrenceDateTimes(params parameter.Container, values []interface{}) error {
	if e.RecurrenceDateTimes == nil {
		e.RecurrenceDateTimes = &RecurrenceDateTimes{}
	}
	return e.RecurrenceDateTimes.SetRecurrenceDateTimes(params, values)
}
