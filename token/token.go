package token

type Token struct {
	Type  TokenType
	Value string
}

type TokenType string

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	IDENT     TokenType = "IDENT"
	STRING    TokenType = "STRING"
	ASSIGN    TokenType = "="
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
)
