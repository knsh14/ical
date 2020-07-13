package lexer

import "unicode"

// Lexer converts source code to tokens
type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input ( after current char)
	ch           rune // current char under examination

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

// peekChar returns next char but does not increment position
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, l.ch)
		l.checkFunc = isParamValue
	case ';':
		tok = newToken(SEMICOLON, l.ch)
		l.checkFunc = isParamName
	case ',':
		tok = newToken(COMMA, l.ch)
	case '"':
		tok.Type = STRING
		tok.Value = l.readString()
	case ':':
		tok = newToken(COLON, l.ch)
		l.checkFunc = isValue
	case 0:
		tok.Value = ""
		tok.Type = EOF
	default:
		if l.checkFunc(l.ch) {
			tok.Value = l.readIdentifier()
			tok.Type = IDENT
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Value: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.checkFunc(l.ch) {
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isLetter(ch rune) bool {
	if unicode.IsControl(ch) {
		return ch == rune('\t')
	}
	if ch == rune(',') {
		return false
	}
	return true
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
		return ch == rune('\t')
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
