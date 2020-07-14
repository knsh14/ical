package lexer

import "fmt"

type ContentLine struct {
	Name       string
	Parameters []Parameter
	Values     []string
}

type Parameter struct {
	Name   string
	Values []string
}

func ConvertContentLine(line *Lexer) (*ContentLine, error) {
	var cl ContentLine
	// get name
	n, t, err := getName(line)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %w", err)
	}
	cl.Name = n

	// get parameters until get colon
	for t.Type == SEMICOLON {
		p, token, err := getParameter(line)
		if err != nil {
			return nil, fmt.Errorf("failed to get parameter: %w", err)
		}
		t = token
		cl.Parameters = append(cl.Parameters, p)
	}

	// get values until illegal or eof
	if t.Type != COLON {
		return nil, fmt.Errorf("expected \":\" but got %s[%s]", t.Type, t.Value)
	}
	for t.Type != EOF && t.Type != ILLEGAL {
		v, token, err := getValue(line)
		if err != nil {
			return nil, fmt.Errorf("failed to get value: %w", err)
		}
		t = token
		cl.Values = append(cl.Values, v)
	}
	if t.Type == ILLEGAL {
		return nil, fmt.Errorf("received ILLEGAL %v", t.Value)
	}
	return &cl, nil
}

func getName(lexer *Lexer) (string, Token, error) {
	var n string
	for {
		t := lexer.NextToken()
		switch t.Type {
		case IDENT:
			n += t.Value
		case SEMICOLON, COLON:
			return n, t, nil
		default:
			return "", t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
}

func getParameter(lexer *Lexer) (Parameter, Token, error) {
	var p Parameter
name:
	for {
		t := lexer.NextToken()
		switch t.Type {
		case IDENT:
			p.Name += t.Value
		case ASSIGN:
			break name
		default:
			return Parameter{}, t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
	var val string
	for {
		t := lexer.NextToken()
		switch t.Type {
		case IDENT, STRING:
			val += t.Value
		case COMMA:
			p.Values = append(p.Values, val)
			val = ""
		case COLON, SEMICOLON:
			p.Values = append(p.Values, val)
			return p, t, nil
		default:
			return Parameter{}, t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
}

func getValue(lexer *Lexer) (string, Token, error) {
	var val string
	for {
		t := lexer.NextToken()
		switch t.Type {
		case IDENT, STRING:
			val += t.Value
		case COMMA, EOF:
			return val, t, nil
		default:
			return "", t, fmt.Errorf("invalid token %s", t.Value)
		}
	}
}
