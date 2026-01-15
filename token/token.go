// Package token
package token

type Token struct {
	Pos Position

	Lit  string
	Type Type
}
