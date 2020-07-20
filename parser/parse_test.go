package parser

import (
	"fmt"
	"testing"

	"github.com/knsh14/ical"
)

func TestContentLine(t *testing.T) {

	tests := []struct {
		input  string
		Expect ical.ContentLine
	}{
		{
			input: "BEGIN:VEVENT",
		},
		{
			input: "X-WR-TIMEZONE:Asia/Tokyo",
		},
		{
			input: "DTSTART;VALUE=DATE:20200301",
		},
		{
			input: "RDATE;VALUE=DATE:19970304,19970504,19970704,19970904",
		},
		{
			input: "ATTENDEE;RSVP=TRUE;ROLE=REQ-PARTICIPANT:mailto:jsmith@example.com",
		},
		{
			input: "EXAMPLE;AAA=\"BBBB;CCCC\":DDDD",
		},
		{
			input: "EXAMPLE;URL=\"https://github.com\":OCTOCAT",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
		})
	}
}
