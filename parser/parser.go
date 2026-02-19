package parser

import (
	"errors"
	"fmt"

	"github.com/whererun3000/monkey/ast"
	"github.com/whererun3000/monkey/lexer"
	"github.com/whererun3000/monkey/token"
)

type (
	infixParseFn  func(ast.Expr) ast.Expr
	prefixParseFn func() ast.Expr
)

func ParseProgram(src string) *ast.Program {
	p := new(parser)
	p.init(src)

	program := &ast.Program{}

	for !p.tok.Is(token.EOF) {
		if stmt := p.parseStmt(); stmt != nil {
			program.Stmts = append(program.Stmts, stmt)
		}

		p.next()
	}

	return program
}

type parser struct {
	lexer lexer.Lexer

	tok   token.Token
	rdTok token.Token

	errors []error

	infixParseFns  map[token.Type]infixParseFn
	prefixParseFns map[token.Type]prefixParseFn
}

func (p *parser) init(src string) {
	p.lexer = *lexer.New(src)

	p.infixParseFns = map[token.Type]infixParseFn{
		token.PLUS:     p.parseInfixExpr,
		token.MINUS:    p.parseInfixExpr,
		token.SLASH:    p.parseInfixExpr,
		token.ASTERISK: p.parseInfixExpr,

		token.LT:  p.parseInfixExpr,
		token.GT:  p.parseInfixExpr,
		token.EQ:  p.parseInfixExpr,
		token.NEQ: p.parseInfixExpr,

		token.LPAREN: p.parseCallExpr,
	}

	p.prefixParseFns = map[token.Type]prefixParseFn{
		token.INT:      p.parseIntLit,
		token.FUNCTION: p.parseFuncLit,

		token.IDENT: p.parseIdent,

		token.TRUE:  p.parseBoolLit,
		token.FALSE: p.parseBoolLit,

		token.BANG:  p.parsePrefixExpr,
		token.MINUS: p.parsePrefixExpr,

		token.LPAREN: p.parseGroupExpr,

		token.IF: p.parseIfExpr,
	}

	p.next()
	p.next()
}

func (p *parser) next() {
	p.tok = p.rdTok
	p.rdTok = p.lexer.Next()
}

func (p *parser) expect(t token.Type) bool {
	if p.rdTok.Is(t) {
		p.next()
		return true
	} else {
		p.errorExpect(t)
		return false
	}
}

func (p *parser) errorExpect(t token.Type) {
	p.errors = append(p.errors, fmt.Errorf("expected %s, found %s", t.String(), p.rdTok.Type.String()))
}

func (p *parser) errorf(format string, a ...any) {
	p.error(fmt.Sprintf(format, a...))
}

func (p *parser) error(msg string) {
	p.errors = append(p.errors, errors.New(msg))
}
