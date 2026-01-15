// Package lexer
package lexer

import (
	"github.com/whererun3000/monkey/token"
)

type Lexer struct {
	src []rune

	offset   int
	rdOffset int

	line   int
	column int

	ch rune
}

func New(src string) *Lexer {
	l := &Lexer{
		src: []rune(src),

		line:   1,
		column: 1,
	}

	l.next()
	return l
}

func (l *Lexer) Next() token.Token {
	l.skipWhitespace()

	tok := token.Token{Pos: l.position()}
	switch ch := l.ch; {
	case isLetter(ch):
		tok.Lit = l.readIdent()
		tok.Type = token.Lookup(tok.Lit)
	case isDigit(ch):
		tok.Lit = l.readInt()
		tok.Type = token.INT
	default:
		l.next()
		switch ch {
		case '+':
			tok.Lit = "+"
			tok.Type = token.PLUS
		case '-':
			tok.Lit = "-"
			tok.Type = token.MINUS
		case '*':
			tok.Lit = "*"
			tok.Type = token.ASTERISK
		case '/':
			tok.Lit = "/"
			tok.Type = token.SLASH
		case '!':
			if l.ch == '=' {
				l.next()
				tok.Lit = "!="
				tok.Type = token.NEQ
			} else {
				tok.Lit = "!"
				tok.Type = token.BANG
			}
		case '=':
			if l.ch == '=' {
				l.next()
				tok.Lit = "=="
				tok.Type = token.EQ
			} else {
				tok.Lit = "="
				tok.Type = token.ASSIGN
			}
		case '>':
			tok.Lit = ">"
			tok.Type = token.GT
		case '<':
			tok.Lit = "<"
			tok.Type = token.LT
		case ',':
			tok.Lit = ","
			tok.Type = token.COMMA
		case ';':
			tok.Lit = ";"
			tok.Type = token.SEMICOLON
		case '(':
			tok.Lit = "("
			tok.Type = token.LPAREN
		case ')':
			tok.Lit = ")"
			tok.Type = token.RPAREN
		case '{':
			tok.Lit = "{"
			tok.Type = token.LBRACE
		case '}':
			tok.Lit = "}"
			tok.Type = token.RBRACE
		case 0:
			tok.Type = token.EOF
		default:
			tok.Type = token.ILLEGAL
		}
	}

	return tok
}

func (l *Lexer) next() {
	if l.rdOffset < len(l.src) {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}

		l.ch = l.src[l.rdOffset]

		l.offset = l.rdOffset
		l.rdOffset++

		l.column++

	} else {
		l.ch = 0

		l.offset = len(l.src)
		l.rdOffset = len(l.src)
	}
}

func (l *Lexer) position() token.Position {
	return token.Position{
		Offset: l.offset,

		Line:   l.line,
		Column: l.column,
	}
}

func (l *Lexer) readIdent() string {
	offset := l.offset
	for isLetter(l.ch) {
		l.next()
	}

	return string(l.src[offset:l.offset])
}

func (l *Lexer) readInt() string {
	offset := l.offset
	for isDigit(l.ch) {
		l.next()
	}

	return string(l.src[offset:l.offset])
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.next()
	}
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
