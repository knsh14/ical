package ical

import "time"

type Event struct {
	// required fields
	UID       string
	Timestamp string

	// required if Calender obj dont have METHOD property
	Start string

	// optional , must not set more than 1
	// class / created / description / geo /
	// last-mod / location / organizer / priority /
	// seq / status / summary / transp /
	// url / recurid /
	Class        string
	Created      string
	Description  string
	Geo          string
	LastModified string
	Location     string
	Organizer    string
	Priority     string
	Sequense     string
	Status       string
	Summary      string
	Transparent  string
	URL          string
	RecurID      string

	// The following is OPTIONAL,
	// but SHOULD NOT occur more than once.
	RRule string

	// optional, but End or Duration.
	End      time.Time
	Duration time.Duration

	// attach / attendee / categories / comment /
	// contact / exdate / rstatus / related /
	// resources / rdate / x-prop / iana-prop
	Attach     string
	Attendee   string
	Categories string
	Comment    string
	Contact    string
	ExDate     string
	RStatus    string
	Related    string
	Resources  string
	Rdate      string

	XProp    string
	IANAPror string
}

func (e *Event) implementCalender() {}
