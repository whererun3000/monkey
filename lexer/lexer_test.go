package lexer

import (
	"testing"

	"github.com/whererun3000/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;

		let add = fn(x, y) {
			x + y;
		};

		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;

		if (5 < 10) {
			return true;
		} else {
			return false;
		}

		10 == 10;
		10 != 9;
	`

	tests := []struct {
		expect token.Token
	}{
		{expect: token.Token{Pos: token.Position{Offset: 3, Line: 2, Column: 3}, Lit: "let", Type: token.LET}},
		{expect: token.Token{Pos: token.Position{Offset: 7, Line: 2, Column: 7}, Lit: "five", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 12, Line: 2, Column: 12}, Lit: "=", Type: token.ASSIGN}},
		{expect: token.Token{Pos: token.Position{Offset: 14, Line: 2, Column: 14}, Lit: "5", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 15, Line: 2, Column: 15}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 19, Line: 3, Column: 3}, Lit: "let", Type: token.LET}},
		{expect: token.Token{Pos: token.Position{Offset: 23, Line: 3, Column: 7}, Lit: "ten", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 27, Line: 3, Column: 11}, Lit: "=", Type: token.ASSIGN}},
		{expect: token.Token{Pos: token.Position{Offset: 29, Line: 3, Column: 13}, Lit: "10", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 31, Line: 3, Column: 15}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 36, Line: 5, Column: 3}, Lit: "let", Type: token.LET}},
		{expect: token.Token{Pos: token.Position{Offset: 40, Line: 5, Column: 7}, Lit: "add", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 44, Line: 5, Column: 11}, Lit: "=", Type: token.ASSIGN}},
		{expect: token.Token{Pos: token.Position{Offset: 46, Line: 5, Column: 13}, Lit: "fn", Type: token.FUNCTION}},
		{expect: token.Token{Pos: token.Position{Offset: 48, Line: 5, Column: 15}, Lit: "(", Type: token.LPAREN}},
		{expect: token.Token{Pos: token.Position{Offset: 49, Line: 5, Column: 16}, Lit: "x", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 50, Line: 5, Column: 17}, Lit: ",", Type: token.COMMA}},
		{expect: token.Token{Pos: token.Position{Offset: 52, Line: 5, Column: 19}, Lit: "y", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 53, Line: 5, Column: 20}, Lit: ")", Type: token.RPAREN}},
		{expect: token.Token{Pos: token.Position{Offset: 55, Line: 5, Column: 22}, Lit: "{", Type: token.LBRACE}},
		{expect: token.Token{Pos: token.Position{Offset: 60, Line: 6, Column: 4}, Lit: "x", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 62, Line: 6, Column: 6}, Lit: "+", Type: token.PLUS}},
		{expect: token.Token{Pos: token.Position{Offset: 64, Line: 6, Column: 8}, Lit: "y", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 65, Line: 6, Column: 9}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 69, Line: 7, Column: 3}, Lit: "}", Type: token.RBRACE}},
		{expect: token.Token{Pos: token.Position{Offset: 70, Line: 7, Column: 4}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 75, Line: 9, Column: 3}, Lit: "let", Type: token.LET}},
		{expect: token.Token{Pos: token.Position{Offset: 79, Line: 9, Column: 7}, Lit: "result", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 86, Line: 9, Column: 14}, Lit: "=", Type: token.ASSIGN}},
		{expect: token.Token{Pos: token.Position{Offset: 88, Line: 9, Column: 16}, Lit: "add", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 91, Line: 9, Column: 19}, Lit: "(", Type: token.LPAREN}},
		{expect: token.Token{Pos: token.Position{Offset: 92, Line: 9, Column: 20}, Lit: "five", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 96, Line: 9, Column: 24}, Lit: ",", Type: token.COMMA}},
		{expect: token.Token{Pos: token.Position{Offset: 98, Line: 9, Column: 26}, Lit: "ten", Type: token.IDENT}},
		{expect: token.Token{Pos: token.Position{Offset: 101, Line: 9, Column: 29}, Lit: ")", Type: token.RPAREN}},
		{expect: token.Token{Pos: token.Position{Offset: 102, Line: 9, Column: 30}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 106, Line: 10, Column: 3}, Lit: "!", Type: token.BANG}},
		{expect: token.Token{Pos: token.Position{Offset: 107, Line: 10, Column: 4}, Lit: "-", Type: token.MINUS}},
		{expect: token.Token{Pos: token.Position{Offset: 108, Line: 10, Column: 5}, Lit: "/", Type: token.SLASH}},
		{expect: token.Token{Pos: token.Position{Offset: 109, Line: 10, Column: 6}, Lit: "*", Type: token.ASTERISK}},
		{expect: token.Token{Pos: token.Position{Offset: 110, Line: 10, Column: 7}, Lit: "5", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 111, Line: 10, Column: 8}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 115, Line: 11, Column: 3}, Lit: "5", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 117, Line: 11, Column: 5}, Lit: "<", Type: token.LT}},
		{expect: token.Token{Pos: token.Position{Offset: 119, Line: 11, Column: 7}, Lit: "10", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 122, Line: 11, Column: 10}, Lit: ">", Type: token.GT}},
		{expect: token.Token{Pos: token.Position{Offset: 124, Line: 11, Column: 12}, Lit: "5", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 125, Line: 11, Column: 13}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 130, Line: 13, Column: 3}, Lit: "if", Type: token.IF}},
		{expect: token.Token{Pos: token.Position{Offset: 133, Line: 13, Column: 6}, Lit: "(", Type: token.LPAREN}},
		{expect: token.Token{Pos: token.Position{Offset: 134, Line: 13, Column: 7}, Lit: "5", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 136, Line: 13, Column: 9}, Lit: "<", Type: token.LT}},
		{expect: token.Token{Pos: token.Position{Offset: 138, Line: 13, Column: 11}, Lit: "10", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 140, Line: 13, Column: 13}, Lit: ")", Type: token.RPAREN}},
		{expect: token.Token{Pos: token.Position{Offset: 142, Line: 13, Column: 15}, Lit: "{", Type: token.LBRACE}},
		{expect: token.Token{Pos: token.Position{Offset: 147, Line: 14, Column: 4}, Lit: "return", Type: token.RETURN}},
		{expect: token.Token{Pos: token.Position{Offset: 154, Line: 14, Column: 11}, Lit: "true", Type: token.TRUE}},
		{expect: token.Token{Pos: token.Position{Offset: 158, Line: 14, Column: 15}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 162, Line: 15, Column: 3}, Lit: "}", Type: token.RBRACE}},
		{expect: token.Token{Pos: token.Position{Offset: 164, Line: 15, Column: 5}, Lit: "else", Type: token.ELSE}},
		{expect: token.Token{Pos: token.Position{Offset: 169, Line: 15, Column: 10}, Lit: "{", Type: token.LBRACE}},
		{expect: token.Token{Pos: token.Position{Offset: 174, Line: 16, Column: 4}, Lit: "return", Type: token.RETURN}},
		{expect: token.Token{Pos: token.Position{Offset: 181, Line: 16, Column: 11}, Lit: "false", Type: token.FALSE}},
		{expect: token.Token{Pos: token.Position{Offset: 186, Line: 16, Column: 16}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 190, Line: 17, Column: 3}, Lit: "}", Type: token.RBRACE}},
		{expect: token.Token{Pos: token.Position{Offset: 195, Line: 19, Column: 3}, Lit: "10", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 198, Line: 19, Column: 6}, Lit: "==", Type: token.EQ}},
		{expect: token.Token{Pos: token.Position{Offset: 201, Line: 19, Column: 9}, Lit: "10", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 203, Line: 19, Column: 11}, Lit: ";", Type: token.SEMICOLON}},
		{expect: token.Token{Pos: token.Position{Offset: 207, Line: 20, Column: 3}, Lit: "10", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 210, Line: 20, Column: 6}, Lit: "!=", Type: token.NEQ}},
		{expect: token.Token{Pos: token.Position{Offset: 213, Line: 20, Column: 9}, Lit: "9", Type: token.INT}},
		{expect: token.Token{Pos: token.Position{Offset: 214, Line: 20, Column: 10}, Lit: ";", Type: token.SEMICOLON}},
		// {expect: token.Token{Pos: token.Position{Offset: 216, Line: 21, Column: 2}, Type: token.EOF}},
	}

	l := New(input)

	for i, tt := range tests {
		if got := l.Next(); got != tt.expect {
			t.Errorf("%d: got(%+v) != expect(%+v)", i, got, tt.expect)
		}
	}
}
