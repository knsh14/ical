package property

import (
	"fmt"
	"io"

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

func (a *Attendee) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameAttendee, a.Parameter.String(), a.Value); err != nil {
		return err
	}
	return nil
}

func (a *Attendee) Validate() error {
	// TODO
	return nil
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

func (c *Contact) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameContact, c.Parameter.String(), c.Value); err != nil {
		return err
	}
	return nil
}

func (c *Contact) Validate() error {
	// TODO
	return nil
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

func (o *Organizer) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameOrganizer, o.Parameter.String(), o.Value); err != nil {
		return err
	}
	return nil
}

func (o *Organizer) Validate() error {
	// TODO
	return nil
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
	Value     types.TimeValue
}

func (rid *RecurrenceID) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameRecurrenceID, rid.Parameter.String(), rid.Value); err != nil {
		return err
	}
	return nil
}

func (rid *RecurrenceID) Validate() error {
	// TODO
	return nil
}

func (rid RecurrenceID) SetRecurrenceID(params parameter.Container, value types.TimeValue) error {
	if len(params[parameter.TypeNameReferenceTimezone]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameReferenceTimezone)
	}
	if len(params[parameter.TypeNameRecurrenceIDRange]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameRecurrenceIDRange)
	}
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

func (rt *RelatedTo) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameRelatedTo, rt.Parameter.String(), rt.Value); err != nil {
		return err
	}
	return nil
}

func (rt *RelatedTo) Validate() error {
	// TODO
	return nil
}

func (rt *RelatedTo) SetRelatedTo(params parameter.Container, value types.Text) error {
	if len(params[parameter.TypeNameRelationshipType]) > 1 {
		return fmt.Errorf("too much values for parameter %s", parameter.TypeNameLanguage)
	}
	rt.Parameter = params
	rt.Value = value
	return nil
}

// URL is URL
// maybe name will be changed to UniformResourceLocator
// https://tools.ietf.org/html/rfc5545#section-3.8.4.6
type URL struct {
	Parameter parameter.Container
	Value     types.URI
}

func (u *URL) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameURL, u.Parameter.String(), u.Value); err != nil {
		return err
	}
	return nil
}

func (u *URL) Validate() error {
	// TODO
	return nil
}
func (url *URL) SetURL(params parameter.Container, value types.URI) error {
	url.Parameter = params
	url.Value = value
	return nil
}

// UID is UID
// https://tools.ietf.org/html/rfc5545#section-3.8.4.7
type UID struct {
	Parameter parameter.Container
	Value     types.Text
}

func (u *UID) Decoce(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s%s:%s", NameUID, u.Parameter.String(), u.Value); err != nil {
		return err
	}
	return nil
}

func (u *UID) Validate() error {
	// TODO
	return nil
}

func (uid *UID) SetUID(params parameter.Container, value types.Text) error {
	uid.Parameter = params
	uid.Value = value
	return nil
}
