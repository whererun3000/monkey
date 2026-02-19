package parser

import (
	"testing"
)

func TestParseProgram(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "LetStmt",
			input:  "let x = 5;",
			output: "let x = 5;",
		},
		{
			name:   "ReturnStmt",
			input:  "return 999;",
			output: "return 999;",
		},
		{
			name:   "IdentLit",
			input:  "foobar",
			output: "foobar;",
		},
		{
			name:   "IntLit",
			input:  "119",
			output: "119;",
		},
		{
			name:   "PrefixExpr(-)",
			input:  "-88",
			output: "(-88);",
		},
		{
			name:   "PrefixExpr(!)",
			input:  "!5",
			output: "(!5);",
		},
		{
			name:   "InfixExpr(Prec)",
			input:  "3 + 4 * 5 == 3 * 1 + 4 * 5",
			output: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
		},
		{
			name:   "BoolLit(true)",
			input:  "let foobar = true;",
			output: "let foobar = true;",
		},
		{
			name:   "BoolLit(false)",
			input:  "let foobar = false;",
			output: "let foobar = false;",
		},
		{
			name:   "GroupExpr",
			input:  "(5 + 5) * 2",
			output: "((5 + 5) * 2);",
		},
		{
			name:   "IfExpr(only if)",
			input:  "if(x < y) { x; }",
			output: "if(x < y) { x; };",
		},
		{
			name:   "IfExpr(if else)",
			input:  "if(x < y) { x } else { y }",
			output: "if(x < y) { x; } else { y; };",
		},
		{
			name:   "CallExpr",
			input:  "add(a, b, c)",
			output: "add(a, b, c);",
		},
		{
			name:   "FuncLit",
			input:  "fn(x, y) { x + y; }",
			output: "fn(x, y) { (x + y); };",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := new(parser)
			p.init(tt.input)

			got := ParseProgram(tt.input)
			if errors := p.errors; len(errors) > 0 {
				t.Errorf("parser has %d errors", len(errors))
				for _, msg := range errors {
					t.Errorf("parser error: %q", msg)
				}
				t.FailNow()
			}
			if s := got.String(); s != tt.output {
				t.Errorf("ParseProgram() = %v, want %v", got.String(), tt.output)
			}
		})
	}
}
