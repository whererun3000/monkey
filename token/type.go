package token

import "fmt"

type Type uint8

const (
	_ Type = iota

	EOF
	ILLEGAL

	// Identifiers + literals
	INT   // 1343456
	IDENT // add, foobar, x, y, ...

	// Operators
	BANG
	PLUS
	MINUS
	SLASH
	ASSIGN
	ASTERISK

	LT
	GT

	EQ
	NEQ

	// Delimiters
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	IF
	LET
	ELSE
	TRUE
	FALSE
	RETURN
	FUNCTION
)

var tokens = [...]string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",

	// Identifiers + literals
	INT:   "INT",   // 1343456
	IDENT: "IDENT", // add, foobar, x, y, ...

	// Operators
	BANG:     "!",
	PLUS:     "+",
	MINUS:    "-",
	SLASH:    "/",
	ASSIGN:   "=",
	ASTERISK: "*",

	LT: "<",
	GT: ">",

	EQ:  "==",
	NEQ: "!=",

	// Delimiters
	COMMA:     ",",
	SEMICOLON: ";",

	LPAREN: "(",
	RPAREN: ")",
	LBRACE: "{",
	RBRACE: "}",

	// Keywords
	IF:       "IF",
	LET:      "LET",
	ELSE:     "ELSE",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
	RETURN:   "RETURN",
	FUNCTION: "FUNCTION",
}

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,

	"true":  TRUE,
	"false": FALSE,

	"if":   IF,
	"else": ELSE,

	"return": RETURN,
}

func (t Type) String() string {
	if 0 < t && t < Type(len(tokens)) {
		return tokens[t]
	}

	return fmt.Sprintf("token(%d)", t)
}

func Lookup(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
