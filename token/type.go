package token

type Type string

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	INT   = "INT"
	IDENT = "IDENT"

	PLUS   = "+"
	ASSIGN = "="

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	LET      = "LET"
	FUNCTION = "FUNCTION"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
