package token

import (
	"fmt"
)

type (
	Prec uint8
	Type uint8
)

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

const (
	LOWEST Prec = iota + 1
	PREFIX
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

func (t Type) Prec() Prec {
	switch t {
	case EQ, NEQ:
		return 2
	case LT, GT:
		return 3
	case PLUS:
		return 4
	case ASTERISK, SLASH:
		return 5
	case MINUS, BANG:
		return 6
	case LPAREN:
		return 7
	default:
		return LOWEST
	}
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
