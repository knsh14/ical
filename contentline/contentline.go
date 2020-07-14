package contentline

import (
	"fmt"

	"github.com/knsh14/ical/lexer"
)

type ContentLine struct {
	Name       string
	Parameters []Parameter
	Values     []string
}

type Parameter struct {
	Name   string
	Values []string
}

func ConvertContentLine(l *lexer.Lexer) (*ContentLine, error) {
	var cl ContentLine
	// get name
	n, t, err := getName(l)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %w", err)
	}
	cl.Name = n

	// get parameters until get colon
	for t.Type == lexer.SEMICOLON {
		p, token, err := getParameter(l)
		if err != nil {
			return nil, fmt.Errorf("failed to get parameter: %w", err)
		}
		t = token
		cl.Parameters = append(cl.Parameters, p)
	}

	// get values until illegal or eof
	if t.Type != lexer.COLON {
		return nil, fmt.Errorf("expected \":\" but got %s[%s]", t.Type, t.Value)
	}
	for t.Type != lexer.EOF && t.Type != lexer.ILLEGAL {
		v, token, err := getValue(l)
		if err != nil {
			return nil, fmt.Errorf("failed to get value: %w", err)
		}
		t = token
		cl.Values = append(cl.Values, v)
	}
	if t.Type == lexer.ILLEGAL {
		return nil, fmt.Errorf("received ILLEGAL %v", t.Value)
	}
	return &cl, nil
}

func getName(l *lexer.Lexer) (string, lexer.Token, error) {
	var n string
	for {
		t := l.NextToken()
		switch t.Type {
		case lexer.IDENT:
			n += t.Value
		case lexer.SEMICOLON, lexer.COLON:
			return n, t, nil
		default:
			return "", t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
}

func getParameter(l *lexer.Lexer) (Parameter, lexer.Token, error) {
	var p Parameter
name:
	for {
		t := l.NextToken()
		switch t.Type {
		case lexer.IDENT:
			p.Name += t.Value
		case lexer.ASSIGN:
			break name
		default:
			return Parameter{}, t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
	var val string
	for {
		t := l.NextToken()
		switch t.Type {
		case lexer.IDENT, lexer.STRING:
			val += t.Value
		case lexer.COMMA:
			p.Values = append(p.Values, val)
			val = ""
		case lexer.COLON, lexer.SEMICOLON:
			p.Values = append(p.Values, val)
			return p, t, nil
		default:
			return Parameter{}, t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
}

func getValue(l *lexer.Lexer) (string, lexer.Token, error) {
	var val string
	for {
		t := l.NextToken()
		switch t.Type {
		case lexer.IDENT, lexer.STRING:
			val += t.Value
		case lexer.COMMA, lexer.EOF:
			return val, t, nil
		default:
			return "", t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
}
