package lexer

import (
	"fmt"
	"testing"

	"github.com/knsh14/ical/token"
)

func TestLexer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input  string
		expect []token.Token
	}{
		{
			input: "BEGIN:VEVENT",
			expect: []token.Token{
				{Type: token.IDENT, Value: "BEGIN"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "VEVENT"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "X-WR-TIMEZONE:Asia/Tokyo",
			expect: []token.Token{
				{Type: token.IDENT, Value: "X-WR-TIMEZONE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "Asia/Tokyo"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "DTSTART;VALUE=DATE:20200301",
			expect: []token.Token{
				{Type: token.IDENT, Value: "DTSTART"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "VALUE"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.IDENT, Value: "DATE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "20200301"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "RDATE;VALUE=DATE:19970304,19970504,19970704,19970904",
			expect: []token.Token{
				{Type: token.IDENT, Value: "RDATE"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "VALUE"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.IDENT, Value: "DATE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "19970304"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "19970504"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "19970704"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "19970904"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "ATTENDEE;RSVP=TRUE;ROLE=RASSIGN-PARTICIPANT:mailto:jsmith@example.com",
			expect: []token.Token{
				{Type: token.IDENT, Value: "ATTENDEE"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "RSVP"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.IDENT, Value: "TRUE"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "ROLE"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.IDENT, Value: "RASSIGN-PARTICIPANT"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "mailto:jsmith@example.com"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "EXAMPLE;AAA=\"BBBB;CCCC\":DDDD",
			expect: []token.Token{
				{Type: token.IDENT, Value: "EXAMPLE"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "AAA"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.STRING, Value: "BBBB;CCCC"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "DDDD"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "EXAMPLE;URL=\"https://github.com\":OCTOCAT",
			expect: []token.Token{
				{Type: token.IDENT, Value: "EXAMPLE"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "URL"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.STRING, Value: "https://github.com"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "OCTOCAT"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "EXAMPLE:DDDD,EEEE,FFFF",
			expect: []token.Token{
				{Type: token.IDENT, Value: "EXAMPLE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "DDDD"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "EEEE"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "FFFF"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "EX@MPLE:DDDD,EEEE,FFFF",
			expect: []token.Token{
				{Type: token.IDENT, Value: "EX"},
				{Type: token.ILLEGAL, Value: "@"},
				{Type: token.IDENT, Value: "MPLE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "DDDD"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "EEEE"},
				{Type: token.COMMA, Value: ","},
				{Type: token.IDENT, Value: "FFFF"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "DTSTART;VA`UE=DATE:20200301",
			expect: []token.Token{
				{Type: token.IDENT, Value: "DTSTART"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "VA"},
				{Type: token.ILLEGAL, Value: "`"},
				{Type: token.IDENT, Value: "UE"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.IDENT, Value: "DATE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "20200301"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: "DTSTART;VALUE=D@TE:20200301",
			expect: []token.Token{
				{Type: token.IDENT, Value: "DTSTART"},
				{Type: token.SEMICOLON, Value: ";"},
				{Type: token.IDENT, Value: "VALUE"},
				{Type: token.ASSIGN, Value: "="},
				{Type: token.IDENT, Value: "D@TE"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: "20200301"},
				{Type: token.EOF, Value: ""},
			},
		},
		{
			input: `DESCRIPTION:hello world
this is test`,
			expect: []token.Token{
				{Type: token.IDENT, Value: "DESCRIPTION"},
				{Type: token.COLON, Value: ":"},
				{Type: token.IDENT, Value: `hello world
this is test`},
				{Type: token.EOF, Value: ""},
			},
		},
	}

	for i, tt := range tests {
		tt := tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			lexer := New(tt.input)
			for i := range tt.expect {
				tok := lexer.NextToken()
				if tok.Type != tt.expect[i].Type {
					t.Fatalf("not expected type, expect=%v, got=%v", tt.expect[i].Type, tok.Type)
				}
				if tok.Value != tt.expect[i].Value {
					t.Fatalf("not expected value, expect=%v, got=%v", tt.expect[i].Value, tok.Value)
				}
			}
		})
	}
}
