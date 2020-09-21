package ical

import "io"

type CalenderComponent interface {
	CalendarNode
	implementCalender()
}

type CalendarNode interface {
	Decode(w io.Writer) error
	Validate() error
}
