package contentline

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knsh14/ical/lexer"
)

func TestContentLine(t *testing.T) {
	tests := []struct {
		input       string
		expectValue *ContentLine
		expectError error
	}{
		{
			input: "BEGIN:VEVENT",
			expectValue: &ContentLine{
				Name:       "BEGIN",
				Parameters: nil,
				Values:     []string{"VEVENT"},
			},
			expectError: nil,
		},
		{
			input: "X-WR-TIMEZONE:Asia/Tokyo",
			expectValue: &ContentLine{
				Name:       "X-WR-TIMEZONE",
				Parameters: nil,
				Values:     []string{"Asia/Tokyo"},
			},
			expectError: nil,
		},
		{
			input: "DTSTART;VALUE=DATE:20200301",
			expectValue: &ContentLine{
				Name: "DTSTART",
				Parameters: []Parameter{
					{
						Name:   "VALUE",
						Values: []string{"DATE"},
					},
				},
				Values: []string{"20200301"},
			},
			expectError: nil,
		},
		{
			input: "RDATE;VALUE=DATE:19970304,19970504,19970704,19970904",
			expectValue: &ContentLine{
				Name: "RDATE",
				Parameters: []Parameter{
					{
						Name:   "VALUE",
						Values: []string{"DATE"},
					},
				},
				Values: []string{"19970304", "19970504", "19970704", "19970904"},
			},
			expectError: nil,
		},
		{
			input: "ATTENDEE;RSVP=TRUE;ROLE=RASSIGN-PARTICIPANT:mailto:jsmith@example.com",
			expectValue: &ContentLine{
				Name: "ATTENDEE",
				Parameters: []Parameter{
					{
						Name:   "RSVP",
						Values: []string{"TRUE"},
					},
					{
						Name:   "ROLE",
						Values: []string{"RASSIGN-PARTICIPANT"},
					},
				},
				Values: []string{"mailto:jsmith@example.com"},
			},
			expectError: nil,
		},
		{
			input: "EXAMPLE;AAA=\"BBBB;CCCC\":DDDD",
			expectValue: &ContentLine{
				Name: "EXAMPLE",
				Parameters: []Parameter{
					{
						Name:   "AAA",
						Values: []string{"BBBB;CCCC"},
					},
				},
				Values: []string{"DDDD"},
			},
			expectError: nil,
		},
		{
			input: "EXAMPLE;URL=\"https://github.com\":OCTOCAT",
			expectValue: &ContentLine{
				Name: "EXAMPLE",
				Parameters: []Parameter{
					{
						Name:   "URL",
						Values: []string{"https://github.com"},
					},
				},
				Values: []string{"OCTOCAT"},
			},
			expectError: nil,
		},
		{
			input: "EXAMPLE:DDDD,EEEE,FFFF",
			expectValue: &ContentLine{
				Name:       "EXAMPLE",
				Parameters: nil,
				Values:     []string{"DDDD", "EEEE", "FFFF"},
			},
			expectError: nil,
		},
		{
			input:       "EX@MPLE:DDDD,EEEE,FFFF",
			expectValue: nil,
			expectError: fmt.Errorf("failed to get name: invalid token @"),
		},
		{
			input:       "DTSTART;VA`UE=DATE:20200301",
			expectValue: nil,
			expectError: fmt.Errorf("failed to get parameter: invalid token `"),
		},
		{
			input: "DTSTART;VALUE=D@TE:20200301",
			expectValue: &ContentLine{
				Name: "DTSTART",
				Parameters: []Parameter{
					{
						Name:   "VALUE",
						Values: []string{"D@TE"},
					},
				},
				Values: []string{"20200301"},
			},
			expectError: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			lexer := lexer.New(tt.input)
			cl, err := ConvertContentLine(lexer)
			if err != nil {
				if tt.expectError == nil {
					t.Fatal(err)
				}
				if err.Error() != tt.expectError.Error() {
					t.Fatalf("unexpected error:\n\texpect: %v\n\tgot: %v", tt.expectError.Error(), err.Error())
				}
			}
			if diff := cmp.Diff(tt.expectValue, cl); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
