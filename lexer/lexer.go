package lexer

import (
	"unicode"

	"github.com/knsh14/ical/token"
)

// Lexer converts source code to tokens
type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune

	checkFunc func(rune) bool
}

// New returns lexer
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input), checkFunc: isName}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
		l.checkFunc = isParamValue
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
		l.checkFunc = isParamName
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Value = l.readString()
	case ':':
		tok = newToken(token.COLON, l.ch)
		l.checkFunc = isValue
	case 0:
		tok.Value = ""
		tok.Type = token.EOF
	default:
		if l.checkFunc(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = token.IDENT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.Type, ch rune) token.Token {
	return token.Token{Type: tokenType, Value: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.checkFunc(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDoubleQuoteSafeLetter(ch rune) bool {
	if unicode.IsControl(ch) {
		return ch == rune('\t')
	}
	return ch != rune('"')
}

func isName(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '-'
}

func isParamName(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '-'
}
func isParamValue(ch rune) bool {
	if unicode.IsControl(ch) {
		return ch == rune('\t')
	}
	switch ch {
	case '=', ';', ',', '"', ':':
		return false
	}
	return true
}

func isValue(ch rune) bool {
	if unicode.IsControl(ch) {
		return ch == rune('\t') || ch == rune('\n') || ch == rune('\r')
	}
	if ch == rune(',') {
		return false
	}
	return true
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if !isDoubleQuoteSafeLetter(l.ch) {
			break
		}
	}

	return string(l.input[position:l.position])
}
