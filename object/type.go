package object

type Type uint8

func (t Type) String() string {
	return objects[t]
}

const (
	_ Type = iota

	NULL

	INT
	BOOL
	HASH
	ARRAY
	STRING

	FUNC
	BUILTIN

	RETURN
	ERROR
)

var objects = [...]string{
	NULL: "NULL",

	INT:    "INT",
	BOOL:   "BOOL",
	HASH:   "HASH",
	ARRAY:  "ARRAY",
	STRING: "STRING",

	FUNC:    "FUNC",
	BUILTIN: "BUILTIN",

	RETURN: "RETURN",
	ERROR:  "ERROR",
}
