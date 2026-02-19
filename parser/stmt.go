package parser

import (
	"github.com/whererun3000/monkey/ast"
	"github.com/whererun3000/monkey/token"
)

func (p *parser) parseStmt() ast.Stmt {
	switch p.tok.Type {
	case token.LET:
		return p.parseLetStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	default:
		return p.parseExprStmt()
	}
}

func (p *parser) parseLetStmt() *ast.LetStmt {
	stmt := &ast.LetStmt{
		Tok: p.tok,
	}

	if !p.expect(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Ident{
		Tok:   p.tok,
		Value: p.tok.Lit,
	}

	if !p.expect(token.ASSIGN) {
		return nil
	}

	p.next()
	stmt.Value = p.parseExpr(token.LOWEST)

	if p.rdTok.Is(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *parser) parseReturnStmt() *ast.ReturnStmt {
	stmt := &ast.ReturnStmt{
		Tok: p.tok,
	}

	p.next()
	stmt.Result = p.parseExpr(token.LOWEST)

	if p.rdTok.Is(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *parser) parseExprStmt() *ast.ExprStmt {
	stmt := &ast.ExprStmt{
		Expr: p.parseExpr(token.LOWEST),
	}

	if p.rdTok.Is(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *parser) parseBlockStmt() *ast.BlockStmt {
	block := &ast.BlockStmt{
		Tok: p.tok,
	}

	p.next()

	for !p.tok.Is(token.RBRACE) && !p.tok.Is(token.EOF) {
		if stmt := p.parseStmt(); stmt != nil {
			block.List = append(block.List, stmt)
		}

		p.next()
	}

	return block
}
