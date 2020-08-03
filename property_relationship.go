package ical

import (
	"fmt"

	"github.com/knsh14/ical/parameter"
	"github.com/knsh14/ical/types"
)

// https://tools.ietf.org/html/rfc5545#section-3.8.4

// Attendee is ATTENDEE
// https://tools.ietf.org/html/rfc5545#section-3.8.4.1
type Attendee struct {
	Parameter parameter.Container
	Value     types.CalenderUserAddress
}

func (a *Attendee) SetAttendee(params parameter.Container, value types.CalenderUserAddress) error {

	for _, pname := range []parameter.TypeName{parameter.TypeNameCalenderUserType, parameter.TypeNameMembership, parameter.TypeNameParticipationRole, parameter.TypeNameParticipationStatus, parameter.TypeNameRSVP, parameter.TypeNameDelegatee, parameter.TypeNameDelegatee, parameter.TypeNameSentBy, parameter.TypeNameCommonName, parameter.TypeNameDirectoryEntry, parameter.TypeNameLanguage} {
		if len(params[pname]) > 1 {
			return fmt.Errorf("too much values for parameter %s", pname)
		}
	}
	a.Parameter = params
	a.Value = value
	return nil
}

// Contact is CONTACT
// https://tools.ietf.org/html/rfc5545#section-3.8.4.2
type Contact struct {
	Parameter parameter.Container
	Value     types.Text
}

func (c *Contact) SetContact(params parameter.Container, value types.Text) error {
	if len(params[parameter.TypeNameAlternateTextRepresentation]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameAlternateTextRepresentation)
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameLanguage)
	}
	c.Parameter = params
	c.Value = value
	return nil
}

// Organizer is ORGANIZER
// https://tools.ietf.org/html/rfc5545#section-3.8.4.3
type Organizer struct {
	Parameter parameter.Container
	Value     types.CalenderUserAddress
}

func (o *Organizer) SetOrganizer(params parameter.Container, value types.CalenderUserAddress) error {

	if len(params[parameter.TypeNameCommonName]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameCommonName)
	}
	if len(params[parameter.TypeNameDirectoryEntry]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameDirectoryEntry)
	}
	if len(params[parameter.TypeNameSentBy]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameSentBy)
	}
	if len(params[parameter.TypeNameLanguage]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameLanguage)
	}
	o.Parameter = params
	o.Value = value
	return nil
}

// RecurrenceID is RECURRENCE-ID
// https://tools.ietf.org/html/rfc5545#section-3.8.4.4
type RecurrenceID struct {
	Parameter parameter.Container
	Value     types.TimeType
}

// TODO: implement
func (rid RecurrenceID) SetReccuenceID(params parameter.Container, value types.TimeType) error {

	rid.Parameter = params
	rid.Value = value
	return nil
}

// RelatedTo is RELATED-TO
// https://tools.ietf.org/html/rfc5545#section-3.8.4.5
type RelatedTo struct {
	Parameter parameter.Container
	Value     types.Text
}

// URL is URL
// maybe name will be changed to UniformResourceLocator
// https://tools.ietf.org/html/rfc5545#section-3.8.4.6
type URL struct {
	Parameter parameter.Container
	Value     types.URI
}

// UID is UID
// https://tools.ietf.org/html/rfc5545#section-3.8.4.7
type UID struct {
	Parameter parameter.Container
	Value     types.Text
}