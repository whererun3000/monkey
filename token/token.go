// Package token
package token

type Token struct {
	Pos Position

	Lit  string
	Type Type
}

func (t Token) Is(tt Type) bool {
	return t.Type == tt
}

func (t Token) Prec() Prec {
	return t.Type.Prec()
}
