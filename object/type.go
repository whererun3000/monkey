package object

type Type uint8

const (
	_ Type = iota

	NULL

	INT
	BOOL
)

var objects = [...]string{
	NULL: "NULL",

	INT:  "INT",
	BOOL: "BOOL",
}
