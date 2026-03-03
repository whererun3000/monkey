package eval

import (
	"testing"

	"github.com/whererun3000/monkey/lexer"
	"github.com/whererun3000/monkey/object"
	"github.com/whererun3000/monkey/parser"
)

func TestEval(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		input  string
		output string
	}{
		{
			name:   "int",
			input:  "5",
			output: "5",
		},
		{
			name:   "negative int",
			input:  "-5",
			output: "-5",
		},
		{
			name:   "bool true",
			input:  "true",
			output: "true",
		},
		{
			name:   "bool false",
			input:  "false",
			output: "false",
		},
		{
			name:   "",
			input:  "!true",
			output: "false",
		},
		{
			name:   "",
			input:  "!false",
			output: "true",
		},
		{
			name:   "",
			input:  "!5",
			output: "false",
		},
		{
			name:   "",
			input:  "!!true",
			output: "true",
		},
		{
			name:   "",
			input:  "!!false",
			output: "false",
		},
		{
			name:   "",
			input:  "!!5",
			output: "true",
		},
		{
			name:   "",
			input:  "5 + 5 + 5 + 5 - 10",
			output: "10",
		},
		{
			name:   "",
			input:  "2 * 2 * 2 * 2 * 2",
			output: "32",
		},
		{
			name:   "",
			input:  "-50 + 100 + -50",
			output: "0",
		},
		{
			name:   "",
			input:  "5 * 2 + 10",
			output: "20",
		},
		{
			name:   "",
			input:  "5 + 2 * 10",
			output: "25",
		},
		{
			name:   "",
			input:  "20 + 2 * -10",
			output: "0",
		},
		{
			name:   "",
			input:  "50 / 2 * 2 + 10",
			output: "60",
		},
		{
			name:   "",
			input:  "2 * (5 + 10)",
			output: "30",
		},
		{
			name:   "",
			input:  "3 * 3 * 3 + 10",
			output: "37",
		},
		{
			name:   "",
			input:  "3 * (3 * 3) + 10",
			output: "37",
		},
		{
			name:   "",
			input:  "(5 + 10 * 2 + 15 / 3) * 2 + -10",
			output: "50",
		},
		{
			name:   "",
			input:  "true == true",
			output: "true",
		},
		{
			name:   "",
			input:  "false == false",
			output: "true",
		},
		{
			name:   "",
			input:  "true != false",
			output: "true",
		},
		{
			name:   "",
			input:  "false != true",
			output: "true",
		},
		{
			name:   "",
			input:  "(1 < 2) == true",
			output: "true",
		},
		{
			name:   "",
			input:  "(1 < 2) == false",
			output: "false",
		},
		{
			name:   "",
			input:  "(1 > 2) == true",
			output: "false",
		},
		{
			name:   "",
			input:  "(1 > 2) == false",
			output: "true",
		},
		{
			name:   "",
			input:  "if (true) { 10 }",
			output: "10",
		},
		{
			name:   "",
			input:  "if (false) { 10 }",
			output: "null",
		},
		{
			name:   "",
			input:  "if (1) { 10 }",
			output: "10",
		},
		{
			name:   "",
			input:  "if (1 < 2) { 10 }",
			output: "10",
		},
		{
			name:   "",
			input:  "if (1 > 2) { 10 }",
			output: "null",
		},
		{
			name:   "",
			input:  "if (1 > 2) { 10 } else { 20 }",
			output: "20",
		},
		{
			name:   "",
			input:  "if (1 < 2) { 10 } else { 20 }",
			output: "10",
		},
		{
			name:   "",
			input:  "return 10;",
			output: "10",
		},
		{
			name:   "",
			input:  "return 10; 9;",
			output: "10",
		},
		{
			name:   "",
			input:  "return 2 * 5; 9;",
			output: "10",
		},
		{
			name:   "",
			input:  "9; return 2 * 5; 9;",
			output: "10",
		},
		{
			name: "",
			input: `
			if (10 > 1) {
				return 10;
			}

			return 1;
			`,
			output: "10",
		},
		{
			name:   "",
			input:  "5 + true",
			output: "ERROR: type mismatch: INT + BOOL",
		},
		{
			name:   "",
			input:  "5 + true; 5",
			output: "ERROR: type mismatch: INT + BOOL",
		},
		{
			name:   "",
			input:  "-true",
			output: "ERROR: unknown operator: -BOOL",
		},
		{
			name:   "",
			input:  "true + false",
			output: "ERROR: unknown operator: BOOL + BOOL",
		},
		{
			name:   "",
			input:  "5; true + false; 5",
			output: "ERROR: unknown operator: BOOL + BOOL",
		},
		{
			name:   "",
			input:  "if (10 > 1) { true + false; }",
			output: "ERROR: unknown operator: BOOL + BOOL",
		},
		{
			name: "",
			input: `
			if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}

				return 1;
			}
			`,
			output: "ERROR: unknown operator: BOOL + BOOL",
		},
		{
			name:   "",
			input:  "let a = 5; a",
			output: "5",
		},
		{
			name:   "",
			input:  "let a = 5 * 5; a;",
			output: "25",
		},
		{
			name:   "",
			input:  "let a = 5; let b = 5; b",
			output: "5",
		},
		{
			name:   "",
			input:  "let a = 5; let b = a; let c = a + b + 5; c;",
			output: "15",
		},
		{
			name:   "",
			input:  "fn(x) { x + 2; }",
			output: "fn(x) { (x + 2); }",
		},
		{
			name:   "",
			input:  "let identify = fn(x) { x; }; identify(5)",
			output: "5",
		},
		{
			name: "",
			input: `
			let newAdder = fn(x) {
				fn(y) { x + y; };
			};

			let addTwo = newAdder(2);
			addTwo(2);
			`,
			output: "4",
		},
		{
			name:   "string",
			input:  `"Hello World"`,
			output: `"Hello World"`,
		},
		{
			name:   "string + string",
			input:  `"Hello" + " " + "World!"`,
			output: `"Hello World!"`,
		},
		{
			name:   "len zero",
			input:  `len("")`,
			output: "0",
		},
		{
			name:   "len string",
			input:  `len("four")`,
			output: "4",
		},
		{
			name:   "len int",
			input:  `len(1)`,
			output: "ERROR: argument to `len` not supported, got INT",
		},
		{
			name:   "len multiple args",
			input:  `len("one", "two")`,
			output: "ERROR: wrong number of arguments. got = 2, want = 1",
		},
		{
			name:   "array lit",
			input:  `[1, 2 * 2, 3 + 3, "hello world"]`,
			output: `[1, 4, 6, "hello world"]`,
		},
		{
			name:   "",
			input:  "[1, 2, 3][0]",
			output: "1",
		},
		{
			name:   "",
			input:  "let arr = [1, 2, 3]; arr[2];",
			output: "3",
		},
		{
			name:   "",
			input:  `{"foo": 5}["foo"]`,
			output: "5",
		},
		{
			name:   "",
			input:  `{"foo": 5}["bar"]`,
			output: "null",
		},
		{
			name:   "",
			input:  `let key = "foo"; {"foo": 5}[key]`,
			output: "5",
		},
		{
			name:   "",
			input:  `{}["foo"]`,
			output: "null",
		},
		{
			name:   "",
			input:  `{5: 5}[5]`,
			output: "5",
		},
		{
			name:   "",
			input:  `{true: 5}[true]`,
			output: "5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.input)
			p := parser.New(l)

			program := p.Parse()

			env := object.NewEnv(nil)
			got := Eval(program, env)
			if s := got.String(); s != tt.output {
				t.Errorf("program: %+v", program)
				t.Errorf("Eval(%s) = %s, want %s", tt.input, s, tt.output)
			}
		})
	}
}
