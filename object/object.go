package object

import (
	"fmt"
	"hash/fnv"
	"slices"
	"strings"

	"github.com/whererun3000/monkey/ast"
)

type (
	Object interface {
		Type() Type
		String() string
	}

	Hashable interface {
		HashKey() HashKey
	}

	BuiltinFunc func(args ...Object) Object
)

type (
	HashKey struct {
		Type  Type
		Value int64
	}

	HashPair struct {
		Key   Object
		Value Object
	}
)

// Object Implementations
type (
	Null struct{}

	Bool struct {
		Value bool
	}

	Int struct {
		Value int64
	}

	Array struct {
		Elems []Object
	}

	Hash struct {
		Pairs map[HashKey]HashPair
	}

	String struct {
		Value string
	}

	Func struct {
		Args []*ast.Ident
		Body *ast.BlockStmt

		Env *Env
	}

	Builtin struct {
		Fn BuiltinFunc
	}

	Return struct {
		Value Object
	}

	Error struct {
		Message string
	}
)

func (n *Null) Type() Type     { return NULL }
func (n *Null) String() string { return "null" }

func (i *Int) Type() Type       { return INT }
func (i *Int) String() string   { return fmt.Sprintf("%d", i.Value) }
func (i *Int) HashKey() HashKey { return HashKey{Type: i.Type(), Value: i.Value} }

func (b *Bool) Type() Type     { return BOOL }
func (b *Bool) String() string { return fmt.Sprintf("%t", b.Value) }
func (b *Bool) HashKey() HashKey {
	var value int64
	if b.Value {
		value = 1
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (a *Array) Type() Type { return ARRAY }
func (a *Array) String() string {
	var sb strings.Builder

	elems := make([]string, 0, len(a.Elems))
	for _, v := range a.Elems {
		elems = append(elems, v.String())
	}

	sb.WriteString("[")
	sb.WriteString(strings.Join(elems, ", "))
	sb.WriteString("]")

	return sb.String()
}

func (h *Hash) Type() Type { return HASH }
func (h *Hash) String() string {
	var sb strings.Builder

	pairs := make([]string, 0, len(h.Pairs))
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.String(), pair.Value.String()))
	}

	slices.Sort(pairs)

	sb.WriteString("{")
	sb.WriteString(strings.Join(pairs, ", "))
	sb.WriteString("}")

	return sb.String()
}

func (s *String) Type() Type     { return STRING }
func (s *String) String() string { return fmt.Sprintf("%q", s.Value) }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: int64(h.Sum64())}
}

func (f *Func) Type() Type { return FUNC }
func (f *Func) String() string {
	var sb strings.Builder

	args := make([]string, 0, len(f.Args))
	for _, arg := range f.Args {
		args = append(args, arg.String())
	}

	sb.WriteString("fn(")
	sb.WriteString(strings.Join(args, ", "))
	sb.WriteString(") ")
	sb.WriteString(f.Body.String())

	return sb.String()
}

func (b *Builtin) Type() Type     { return BUILTIN }
func (b *Builtin) String() string { return "builtin function" }

func (r *Return) Type() Type     { return RETURN }
func (r *Return) String() string { return r.Value.String() }

func (e *Error) Type() Type     { return ERROR }
func (e *Error) String() string { return fmt.Sprintf("ERROR: %s", e.Message) }
