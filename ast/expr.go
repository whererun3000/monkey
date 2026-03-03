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

type StringLit struct {
	Tok   token.Token
	Value string
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

type ArrayLit struct {
	Tok   token.Token
	Elems []Expr
}

type HashLit struct {
	Tok   token.Token
	Pairs map[Expr]Expr
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

type IndexExpr struct {
	Tok token.Token

	X Expr
	I Expr
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
func (x *StringLit) Token() token.Token  { return x.Tok }
func (x *BoolLit) Token() token.Token    { return x.Tok }
func (x *FuncLit) Token() token.Token    { return x.Tok }
func (x *ArrayLit) Token() token.Token   { return x.Tok }
func (x *HashLit) Token() token.Token    { return x.Tok }
func (x *IfExpr) Token() token.Token     { return x.Tok }
func (x *CallExpr) Token() token.Token   { return x.Tok }
func (x *IndexExpr) Token() token.Token  { return x.Tok }
func (x *InfixExpr) Token() token.Token  { return x.Op }
func (x *PrefixExpr) Token() token.Token { return x.Op }

func (x *Ident) String() string     { return x.Value }
func (x *IntLit) String() string    { return x.Tok.Lit }
func (x *StringLit) String() string { return fmt.Sprintf("%q", x.Tok.Lit) }
func (x *BoolLit) String() string   { return x.Tok.Lit }
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

func (x *ArrayLit) String() string {
	var sb strings.Builder

	elems := make([]string, 0, len(x.Elems))
	for _, elem := range x.Elems {
		elems = append(elems, elem.String())
	}

	sb.WriteString("[")
	sb.WriteString(strings.Join(elems, ", "))
	sb.WriteString("]")

	return sb.String()
}

func (x *HashLit) String() string {
	var sb strings.Builder

	pairs := make([]string, 0, len(x.Pairs))
	for k, v := range x.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s:%s", k.String(), v.String()))
	}

	sb.WriteString("{")
	sb.WriteString(strings.Join(pairs, ", "))
	sb.WriteString("}")

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

func (x *IndexExpr) String() string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(x.X.String())
	sb.WriteString("[")
	sb.WriteString(x.I.String())
	sb.WriteString("])")

	return sb.String()
}

func (x *InfixExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", x.X.String(), x.Op.Lit, x.Y.String())
}
func (x *PrefixExpr) String() string { return fmt.Sprintf("(%s%s)", x.Op.Lit, x.X.String()) }

func (x *Ident) exprNode()      {}
func (x *IntLit) exprNode()     {}
func (x *StringLit) exprNode()  {}
func (x *BoolLit) exprNode()    {}
func (x *FuncLit) exprNode()    {}
func (x *ArrayLit) exprNode()   {}
func (x *HashLit) exprNode()    {}
func (x *IfExpr) exprNode()     {}
func (x *CallExpr) exprNode()   {}
func (x *IndexExpr) exprNode()  {}
func (x *PrefixExpr) exprNode() {}
func (x *InfixExpr) exprNode()  {}
