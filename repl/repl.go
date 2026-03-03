package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/whererun3000/monkey/eval"
	"github.com/whererun3000/monkey/lexer"
	"github.com/whererun3000/monkey/object"
	"github.com/whererun3000/monkey/parser"
)

func Start(r io.Reader, w io.Writer) {
	env := object.NewEnv(nil)
	scanner := bufio.NewScanner(r)

	for {
		_, _ = fmt.Fprintf(w, ">> ")

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.Parse()
		if errs := p.Errors(); len(errs) > 0 {
			continue
		}

		if result := eval.Eval(program, env); result != nil {
			_, _ = fmt.Fprintf(w, "%s\n", result.String())
		}
	}
}
