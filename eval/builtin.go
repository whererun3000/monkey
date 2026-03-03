package eval

import (
	"fmt"

	"github.com/whererun3000/monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if n := len(args); n != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", n)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Int{Value: int64(len(arg.Elems))}
			case *object.String:
				return &object.Int{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if n := len(args); n != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", n)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elems) > 0 {
					return arg.Elems[0]
				}

				return Null
			default:
				return newError("argument to `first` must be ARRAY, got %s", arg.Type().String())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if n := len(args); n != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", n)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if n := len(arg.Elems); n > 0 {
					return arg.Elems[n-1]
				}

				return Null
			default:
				return newError("arguments to `last` must be ARRAY, got %s", arg.Type().String())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if n := len(args); n != 1 {
				return newError("wrong number of arguments. got = %d, want = 1", n)
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if n := len(arg.Elems); n > 0 {
					elems := make([]object.Object, n-1)
					copy(elems, arg.Elems[1:n])
					return &object.Array{Elems: elems}
				}

				return Null
			default:
				return newError("arguments to `last` must be ARRAY, got %s", arg.Type().String())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if n := len(args); n != 2 {
				return newError("wrong number of arguments. got = %d, want = 2", n)
			}

			switch arg0 := args[0].(type) {
			case *object.Array:
				n0 := len(arg0.Elems)
				elems := make([]object.Object, n0+1)
				copy(elems, arg0.Elems)
				elems[n0] = args[1]
				return &object.Array{Elems: elems}
			default:
				return newError("arguments to `push` must be ARRAY, got %s", arg0.Type().String())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.String())
			}

			return Null
		},
	},
}
