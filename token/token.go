package token

type Token struct {
	Type  Type
	Value string
}

type Type string

const (
	ILLEGAL   Type = "ILLEGAL"
	EOF       Type = "EOF"
	IDENT     Type = "IDENT"
	STRING    Type = "STRING"
	ASSIGN    Type = "="
	COMMA     Type = ","
	SEMICOLON Type = ";"
	COLON     Type = ":"
)
