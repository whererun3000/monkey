package parser

import (
	"strconv"

	"github.com/whererun3000/monkey/ast"
	"github.com/whererun3000/monkey/token"
)

func (p *parser) parseIdent() ast.Expr {
	return &ast.Ident{
		Tok:   p.tok,
		Value: p.tok.Lit,
	}
}

func (p *parser) parseIntLit() ast.Expr {
	value, err := strconv.ParseInt(p.tok.Lit, 10, 64)
	if err != nil {
		p.errorf("could not parse %q as integer", p.tok.Lit)
		return nil
	}

	lit := &ast.IntLit{
		Tok:   p.tok,
		Value: value,
	}

	return lit
}

func (p *parser) parseBoolLit() ast.Expr {
	value, err := strconv.ParseBool(p.tok.Lit)
	if err != nil {
		p.errorf("could not parse %q as bool", p.tok.Lit)
		return nil
	}

	lit := &ast.BoolLit{
		Tok:   p.tok,
		Value: value,
	}

	return lit
}

func (p *parser) parseFuncLit() ast.Expr {
	lit := &ast.FuncLit{
		Tok: p.tok,
	}

	if !p.expect(token.LPAREN) {
		return nil
	}

	lit.Params = p.parseFuncParams()

	if !p.expect(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStmt()

	return lit
}

func (p *parser) parseFuncParams() []*ast.Ident {
	p.next()

	if p.tok.Is(token.RPAREN) {
		return nil
	}

	idents := []*ast.Ident{
		{
			Tok:   p.tok,
			Value: p.tok.Lit,
		},
	}

	for p.rdTok.Is(token.COMMA) {
		p.next()
		p.next()

		idents = append(idents, &ast.Ident{
			Tok:   p.tok,
			Value: p.tok.Lit,
		})
	}

	if !p.expect(token.RPAREN) {
		return nil
	}

	return idents
}

func (p *parser) parseGroupExpr() ast.Expr {
	p.next()

	x := p.parseExpr(token.LOWEST)

	if !p.expect(token.RPAREN) {
		return nil
	}

	return x
}

func (p *parser) parseIfExpr() ast.Expr {
	expr := &ast.IfExpr{
		Tok: p.tok,
	}

	if !p.expect(token.LPAREN) {
		return nil
	}

	p.next()
	expr.Cond = p.parseExpr(token.LOWEST)

	if !p.expect(token.RPAREN) {
		return nil
	}

	if !p.expect(token.LBRACE) {
		return nil
	}

	expr.Body = p.parseBlockStmt()

	if p.rdTok.Is(token.ELSE) {
		p.next()

		if !p.expect(token.LBRACE) {
			return nil
		}

		expr.Else = p.parseBlockStmt()
	}

	return expr
}

func (p *parser) parseCallExpr(fn ast.Expr) ast.Expr {
	expr := &ast.CallExpr{
		Tok: p.tok,

		Fn:   fn,
		Args: p.parseCallArgs(),
	}

	return expr
}

func (p *parser) parseCallArgs() []ast.Expr {
	if p.next(); p.tok.Is(token.RPAREN) {
		return nil
	}

	args := []ast.Expr{
		p.parseExpr(token.LOWEST),
	}

	for p.rdTok.Is(token.COMMA) {
		p.next()
		p.next()
		args = append(args, p.parseExpr(token.LOWEST))
	}

	if !p.expect(token.RPAREN) {
		return nil
	}

	return args
}

func (p *parser) parseExpr(prec token.Prec) ast.Expr {
	prefix := p.prefixParseFns[p.tok.Type]
	if prefix == nil {
		p.errorf("no prefix parse function for %s found", p.tok.Type)
		return nil
	}

	x := prefix()

	for p.rdTok.Prec() > prec {
		infix := p.infixParseFns[p.rdTok.Type]
		if infix == nil {
			break
		}

		p.next()

		x = infix(x)
	}

	return x
}

func (p *parser) parsePrefixExpr() ast.Expr {
	expr := &ast.PrefixExpr{
		Op: p.tok,
	}

	p.next()

	expr.X = p.parseExpr(p.tok.Prec())

	return expr
}

func (p *parser) parseInfixExpr(x ast.Expr) ast.Expr {
	expr := &ast.InfixExpr{
		Op: p.tok,

		X: x,
	}

	prec := p.tok.Prec()
	p.next()
	expr.Y = p.parseExpr(prec)

	return expr
}
