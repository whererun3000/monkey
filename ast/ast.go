package ast

import (
	"strings"

	"github.com/whererun3000/monkey/token"
)

type Node interface {
	Token() token.Token
	String() string
}

type Expr interface {
	Node
	String() string
	exprNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type Program struct {
	Stmts []Stmt
}

func (p *Program) Token() token.Token {
	var tok token.Token
	if len(p.Stmts) > 0 {
		tok = p.Stmts[0].Token()
	}

	return tok
}

func (p *Program) String() string {
	var out strings.Builder
	for _, stmt := range p.Stmts {
		out.WriteString(stmt.String())
	}

	return out.String()
}
