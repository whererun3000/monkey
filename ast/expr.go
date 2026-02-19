package ast

import (
	"fmt"
	"strings"

	"github.com/whererun3000/monkey/token"
)

type Ident struct {
	Tok   token.Token
	Value string
}

type IntLit struct {
	Tok   token.Token
	Value int64
}

type FuncLit struct {
	Tok    token.Token
	Params []*Ident
	Body   *BlockStmt
}

type BoolLit struct {
	Tok   token.Token
	Value bool
}

type IfExpr struct {
	Tok  token.Token
	Cond Expr
	Body *BlockStmt
	Else *BlockStmt
}

type CallExpr struct {
	Tok token.Token

	Fn   Expr
	Args []Expr
}

type PrefixExpr struct {
	Op token.Token
	X  Expr
}

type InfixExpr struct {
	Op token.Token

	X, Y Expr
}

func (x *Ident) Token() token.Token      { return x.Tok }
func (x *IntLit) Token() token.Token     { return x.Tok }
func (x *BoolLit) Token() token.Token    { return x.Tok }
func (x *FuncLit) Token() token.Token    { return x.Tok }
func (x *IfExpr) Token() token.Token     { return x.Tok }
func (x *CallExpr) Token() token.Token   { return x.Tok }
func (x *InfixExpr) Token() token.Token  { return x.Op }
func (x *PrefixExpr) Token() token.Token { return x.Op }

func (x *Ident) String() string   { return x.Value }
func (x *IntLit) String() string  { return x.Tok.Lit }
func (x *BoolLit) String() string { return x.Tok.Lit }
func (x *FuncLit) String() string {
	var sb strings.Builder

	params := make([]string, 0, len(x.Params))
	for _, v := range x.Params {
		params = append(params, v.String())
	}

	sb.WriteString(x.Tok.Lit)
	sb.WriteString("(")
	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(") ")

	sb.WriteString(x.Body.String())

	return sb.String()
}

func (x *IfExpr) String() string {
	var sb strings.Builder

	sb.WriteString("if")
	sb.WriteString(x.Cond.String())
	sb.WriteString(" ")
	sb.WriteString(x.Body.String())

	if x.Else != nil {
		sb.WriteString(" else ")
		sb.WriteString(x.Else.String())
	}

	return sb.String()
}

func (x *CallExpr) String() string {
	var sb strings.Builder

	args := make([]string, 0, len(x.Args))
	for _, v := range x.Args {
		args = append(args, v.String())
	}

	sb.WriteString(x.Fn.String())
	sb.WriteString("(")
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString(")")

	return sb.String()
}

func (x *InfixExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", x.X.String(), x.Op.Lit, x.Y.String())
}
func (x *PrefixExpr) String() string { return fmt.Sprintf("(%s%s)", x.Op.Lit, x.X.String()) }

func (x *Ident) exprNode()      {}
func (x *IntLit) exprNode()     {}
func (x *BoolLit) exprNode()    {}
func (x *FuncLit) exprNode()    {}
func (x *IfExpr) exprNode()     {}
func (x *CallExpr) exprNode()   {}
func (x *PrefixExpr) exprNode() {}
func (x *InfixExpr) exprNode()  {}
