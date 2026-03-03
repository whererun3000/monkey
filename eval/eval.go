package eval

import (
	"fmt"

	"github.com/whererun3000/monkey/ast"
	"github.com/whererun3000/monkey/object"
	"github.com/whererun3000/monkey/token"
)

var (
	Null  = &object.Null{}
	True  = &object.Bool{Value: true}
	False = &object.Bool{Value: false}
)

func Eval(node ast.Node, env *object.Env) object.Object {
	switch n := node.(type) {
	case *ast.Program:
		return evalProgram(n, env)
	// statements
	case *ast.BlockStmt:
		return evalBlockStmt(n, env)
	case *ast.ExprStmt:
		return Eval(n.Expr, env)
	case *ast.ReturnStmt:
		return evalReturnStmt(n, env)
	case *ast.LetStmt:
		evalLetStmt(n, env)
	// expressions
	case *ast.CallExpr:
		return evalCallExpr(n, env)
	case *ast.IfExpr:
		return evalIfExpr(n, env)
	case *ast.PrefixExpr:
		return evalPrefixExpr(n, env)
	case *ast.InfixExpr:
		return evalInfixExpr(n, env)
	case *ast.FuncLit:
		return evalFunc(n, env)
	case *ast.Ident:
		return evalIdent(n, env)
	case *ast.IndexExpr:
		return evalIndexExpr(n, env)
	// literals
	case *ast.IntLit:
		return &object.Int{Value: n.Value}
	case *ast.BoolLit:
		return evalBool(n.Value)
	case *ast.HashLit:
		return evalHash(n, env)
	case *ast.ArrayLit:
		return evalArray(n, env)
	case *ast.StringLit:
		return &object.String{Value: n.Value}
	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Env) object.Object {
	var result object.Object
	for _, stmt := range program.Stmts {
		result = Eval(stmt, env)
		switch v := result.(type) {
		case *object.Error:
			return v
		case *object.Return:
			return v.Value
		}
	}

	return result
}

func evalBlockStmt(block *ast.BlockStmt, env *object.Env) object.Object {
	var result object.Object
	for _, stmt := range block.List {
		if result = Eval(stmt, env); result != nil {
			switch result.Type() {
			case object.RETURN, object.ERROR:
				return result
			}
		}
	}

	return result
}

func evalReturnStmt(stmt *ast.ReturnStmt, env *object.Env) object.Object {
	v := Eval(stmt.Result, env)
	if v.Type() == object.ERROR {
		return v
	}

	return &object.Return{Value: v}
}

func evalLetStmt(stmt *ast.LetStmt, env *object.Env) object.Object {
	v := Eval(stmt.Value, env)
	if v.Type() == object.ERROR {
		return v
	}

	env.Set(stmt.Name.Value, v)
	return Null
}

func evalCallExpr(expr *ast.CallExpr, env *object.Env) object.Object {
	obj := Eval(expr.Fn, env)
	if obj.Type() == object.ERROR {
		return obj
	}

	args := make([]object.Object, 0, len(expr.Args))
	for _, v := range expr.Args {
		arg := Eval(v, env)
		if arg.Type() == object.ERROR {
			return arg
		}

		args = append(args, arg)
	}

	switch fn := obj.(type) {
	case *object.Func:
		fnEnv := object.NewEnv(fn.Env)
		for i, arg := range fn.Args {
			fnEnv.Set(arg.Value, args[i])
		}

		result := Eval(fn.Body, fnEnv)
		if v, ok := result.(*object.Return); ok {
			return v.Value
		}

		return result
	case *object.Builtin:
		return fn.Fn(args...)
	}

	return newError("not a function: %s", obj.Type())
}

func evalIfExpr(expr *ast.IfExpr, env *object.Env) object.Object {
	cond := Eval(expr.Cond, env)
	if isError(cond) {
		return cond
	}

	if isTruthy(cond) {
		return Eval(expr.Body, env)
	} else if expr.Else != nil {
		return Eval(expr.Else, env)
	} else {
		return Null
	}
}

func evalPrefixExpr(expr *ast.PrefixExpr, env *object.Env) object.Object {
	x := Eval(expr.X, env)
	if x.Type() == object.ERROR {
		return x
	}

	switch op := expr.Op.Type; op {
	case token.BANG:
		return evalBangPrefixExpr(x)
	case token.MINUS:
		return evalMinusPrefixExpr(x)
	default:
		return newError("unknown prefix operator: %s", op.String())
	}
}

func evalInfixExpr(expr *ast.InfixExpr, env *object.Env) object.Object {
	x := Eval(expr.X, env)
	if x.Type() == object.ERROR {
		return x
	}

	y := Eval(expr.Y, env)
	if y.Type() == object.ERROR {
		return y
	}

	op := expr.Op.Type

	switch {
	case x.Type() == object.INT && y.Type() == object.INT:
		v1 := x.(*object.Int).Value
		v2 := y.(*object.Int).Value

		switch op {
		case token.PLUS:
			return &object.Int{Value: v1 + v2}
		case token.MINUS:
			return &object.Int{Value: v1 - v2}
		case token.SLASH:
			return &object.Int{Value: v1 / v2}
		case token.ASTERISK:
			return &object.Int{Value: v1 * v2}
		case token.LT:
			return evalBool(v1 < v2)
		case token.GT:
			return evalBool(v1 > v2)
		case token.EQ:
			return evalBool(v1 == v2)
		case token.NEQ:
			return evalBool(v1 != v2)
		default:
			return newError("unknown operator: %s %s %s", x.Type().String(), op.String(), y.Type().String())
		}
	case x.Type() == object.BOOL && y.Type() == object.BOOL:
		o1 := x.(*object.Bool)
		o2 := y.(*object.Bool)

		switch op {
		case token.EQ:
			return evalBool(o1 == o2)
		case token.NEQ:
			return evalBool(o1 != o2)
		default:
			return newError("unknown operator: %s %s %s", x.Type().String(), op.String(), y.Type().String())
		}
	case x.Type() == object.STRING && y.Type() == object.STRING:
		s1 := x.(*object.String).Value
		s2 := y.(*object.String).Value

		switch op {
		case token.PLUS:
			return &object.String{Value: s1 + s2}
		default:
			return newError("unknown operator: %s %s %s", x.Type().String(), op.String(), y.Type().String())
		}
	default:
		return newError("type mismatch: %s %s %s", x.Type().String(), op.String(), y.Type().String())
	}
}

func evalFunc(fn *ast.FuncLit, env *object.Env) object.Object {
	return &object.Func{
		Args: fn.Params,
		Body: fn.Body,

		Env: env,
	}
}

func evalBangPrefixExpr(x object.Object) object.Object {
	switch x {
	case Null, False:
		return True
	default:
		return False
	}
}

func evalMinusPrefixExpr(x object.Object) object.Object {
	switch o := x.(type) {
	case *object.Int:
		return &object.Int{Value: -o.Value}
	default:
		return newError("unknown operator: -%s", x.Type().String())
	}
}

func evalIdent(ident *ast.Ident, env *object.Env) object.Object {
	if v, ok := env.Get(ident.Value); ok {
		return v
	}

	if builtin, ok := builtins[ident.Value]; ok {
		return builtin
	}

	return newError("ident not found: %s", ident.Value)
}

func evalIndexExpr(expr *ast.IndexExpr, env *object.Env) object.Object {
	x := Eval(expr.X, env)
	if x.Type() == object.ERROR {
		return x
	}

	i := Eval(expr.I, env)
	if i.Type() == object.ERROR {
		return i
	}

	switch {
	case x.Type() == object.ARRAY && i.Type() == object.INT:
		xv := x.(*object.Array)
		iv := i.(*object.Int)
		return evalArrayIndex(xv, iv)
	case x.Type() == object.HASH:
		return evalHashIndex(x.(*object.Hash), i)
	default:
		return newError("index operator not supported: %s[%s]", x.Type().String(), i.Type().String())
	}
}

func evalArrayIndex(x *object.Array, i *object.Int) object.Object {
	if i.Value < 0 || i.Value >= int64(len(x.Elems)) {
		return Null
	}

	return x.Elems[i.Value]
}

func evalHashIndex(x *object.Hash, k object.Object) object.Object {
	key, ok := k.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", k.Type().String())
	}

	pair, ok := x.Pairs[key.HashKey()]
	if !ok {
		return Null
	}

	return pair.Value
}

func evalBool(b bool) *object.Bool {
	if b {
		return True
	}

	return False
}

func evalHash(n *ast.HashLit, env *object.Env) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for k, v := range n.Pairs {
		key := Eval(k, env)
		if key.Type() == object.ERROR {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type().String())
		}

		value := Eval(v, env)
		if value.Type() == object.ERROR {
			return value
		}

		pairs[hashKey.HashKey()] = object.HashPair{
			Key:   key,
			Value: value,
		}
	}

	return &object.Hash{Pairs: pairs}
}

func evalArray(n *ast.ArrayLit, env *object.Env) object.Object {
	arr := &object.Array{}
	for _, v := range n.Elems {
		obj := Eval(v, env)
		if obj.Type() == object.ERROR {
			return obj
		}

		arr.Elems = append(arr.Elems, obj)
	}

	return arr
}

func isTruthy(o object.Object) bool {
	switch o {
	case Null, False:
		return false
	default:
		return true
	}
}

func newError(format string, a ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(o object.Object) bool {
	return o != nil && o.Type() == object.ERROR
}
