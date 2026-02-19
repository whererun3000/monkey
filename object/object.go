package object

import "fmt"

type Object interface {
	Type() Type
	String() string
}

type (
	Null struct{}

	Int struct {
		Value int64
	}

	Bool struct {
		Value bool
	}
)

func (n *Null) Type() Type     { return NULL }
func (n *Null) String() string { return "null" }

func (i *Int) Type() Type     { return INT }
func (i *Int) String() string { return fmt.Sprintf("%d", i.Value) }

func (b *Bool) Type() Type     { return BOOL }
func (b *Bool) String() string { return fmt.Sprintf("%t", b.Value) }
