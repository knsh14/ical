package ical

import (
	"io"
)

type CalenderComponent interface {
	implementCalender()
	Write(io.Writer) error
}
