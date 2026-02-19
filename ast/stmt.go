package ast

import (
	"strings"

	"github.com/whererun3000/monkey/token"
)

type (
	LetStmt struct {
		Tok   token.Token
		Name  *Ident
		Value Expr
	}

	ExprStmt struct {
		Expr Expr
	}

	ReturnStmt struct {
		Tok    token.Token
		Result Expr
	}

	BlockStmt struct {
		Tok  token.Token
		List []Stmt
	}
)

func (s *LetStmt) Token() token.Token    { return s.Tok }
func (s *ExprStmt) Token() token.Token   { return s.Expr.Token() }
func (s *BlockStmt) Token() token.Token  { return s.Tok }
func (s *ReturnStmt) Token() token.Token { return s.Tok }

func (s *LetStmt) String() string {
	var out strings.Builder
	out.WriteString(s.Tok.Lit)
	out.WriteString(" ")
	out.WriteString(s.Name.Tok.Lit)
	if s.Value != nil {
		out.WriteString(" = ")
		out.WriteString(s.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (s *ExprStmt) String() string {
	return s.Expr.String() + ";"
}

func (s *BlockStmt) String() string {
	var sb strings.Builder
	sb.WriteString("{ ")
	for _, v := range s.List {
		sb.WriteString(v.String())
	}
	sb.WriteString(" }")
	return sb.String()
}

func (s *ReturnStmt) String() string {
	var out strings.Builder
	out.WriteString(s.Tok.Lit)
	if s.Result != nil {
		out.WriteString(" ")
		out.WriteString(s.Result.String())
	}
	out.WriteString(";")
	return out.String()
}

func (s *LetStmt) stmtNode()    {}
func (s *ExprStmt) stmtNode()   {}
func (s *BlockStmt) stmtNode()  {}
func (s *ReturnStmt) stmtNode() {}
