package lexer

type Token struct {
	Type  TokenType
	Value string
}

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	STRING    = "STRING"
	ASSIGN    = "="
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
)
