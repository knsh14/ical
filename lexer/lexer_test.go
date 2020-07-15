package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input  string
		expect []Token
	}{
		{
			input: "BEGIN:VEVENT",
			expect: []Token{
				{Type: IDENT, Value: "BEGIN"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "VEVENT"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "X-WR-TIMEZONE:Asia/Tokyo",
			expect: []Token{
				{Type: IDENT, Value: "X-WR-TIMEZONE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "Asia/Tokyo"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "DTSTART;VALUE=DATE:20200301",
			expect: []Token{
				{Type: IDENT, Value: "DTSTART"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "VALUE"},
				{Type: ASSIGN, Value: "="},
				{Type: IDENT, Value: "DATE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "20200301"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "RDATE;VALUE=DATE:19970304,19970504,19970704,19970904",
			expect: []Token{
				{Type: IDENT, Value: "RDATE"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "VALUE"},
				{Type: ASSIGN, Value: "="},
				{Type: IDENT, Value: "DATE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "19970304"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "19970504"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "19970704"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "19970904"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "ATTENDEE;RSVP=TRUE;ROLE=RASSIGN-PARTICIPANT:mailto:jsmith@example.com",
			expect: []Token{
				{Type: IDENT, Value: "ATTENDEE"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "RSVP"},
				{Type: ASSIGN, Value: "="},
				{Type: IDENT, Value: "TRUE"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "ROLE"},
				{Type: ASSIGN, Value: "="},
				{Type: IDENT, Value: "RASSIGN-PARTICIPANT"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "mailto:jsmith@example.com"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "EXAMPLE;AAA=\"BBBB;CCCC\":DDDD",
			expect: []Token{
				{Type: IDENT, Value: "EXAMPLE"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "AAA"},
				{Type: ASSIGN, Value: "="},
				{Type: STRING, Value: "BBBB;CCCC"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "DDDD"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "EXAMPLE;URL=\"https://github.com\":OCTOCAT",
			expect: []Token{
				{Type: IDENT, Value: "EXAMPLE"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "URL"},
				{Type: ASSIGN, Value: "="},
				{Type: STRING, Value: "https://github.com"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "OCTOCAT"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "EXAMPLE:DDDD,EEEE,FFFF",
			expect: []Token{
				{Type: IDENT, Value: "EXAMPLE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "DDDD"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "EEEE"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "FFFF"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "EX@MPLE:DDDD,EEEE,FFFF",
			expect: []Token{
				{Type: IDENT, Value: "EX"},
				{Type: ILLEGAL, Value: "@"},
				{Type: IDENT, Value: "MPLE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "DDDD"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "EEEE"},
				{Type: COMMA, Value: ","},
				{Type: IDENT, Value: "FFFF"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "DTSTART;VA`UE=DATE:20200301",
			expect: []Token{
				{Type: IDENT, Value: "DTSTART"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "VA"},
				{Type: ILLEGAL, Value: "`"},
				{Type: IDENT, Value: "UE"},
				{Type: ASSIGN, Value: "="},
				{Type: IDENT, Value: "DATE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "20200301"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "DTSTART;VALUE=D@TE:20200301",
			expect: []Token{
				{Type: IDENT, Value: "DTSTART"},
				{Type: SEMICOLON, Value: ";"},
				{Type: IDENT, Value: "VALUE"},
				{Type: ASSIGN, Value: "="},
				{Type: IDENT, Value: "D@TE"},
				{Type: COLON, Value: ":"},
				{Type: IDENT, Value: "20200301"},
				{Type: EOF, Value: ""},
			},
		},
	}

	for i, tt := range tests {
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
