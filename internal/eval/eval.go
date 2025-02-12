package eval

import (
	"fmt"
	"go_lisp_interpreter/internal/parser"
)

func Eval(ast parser.Node, env *Env) interface{} {
	if ast.Type == "ATOM" {
		if val, ok := env.vars[ast.Value]; ok {
			return val
		}
		return ast.Value
	}

	operator := ast.Children[0].Value
	args := ast.Children[1:]

	switch operator {
	case "+":
		return Eval(args[0], env).(int) + Eval(args[1], env).(int)
	case "-":
		return Eval(args[0], env).(int) - Eval(args[1], env).(int)
	case "*":
		return Eval(args[0], env).(int) * Eval(args[1], env).(int)
	case "/":
		return Eval(args[0], env).(int) / Eval(args[1], env).(int)
	case "define":
		env.vars[args[0].Value] = Eval(args[1], env)
		return nil
	case "if":
		cond := Eval(args[0], env).(bool)
		if cond {
			return Eval(args[1], env)
		} else {
			return Eval(args[2], env)
		}
	case "lambda":
		return ast // Lambda式を関数オブジェクトとして返す
	case "apply":
		fn := Eval(args[0], env).(parser.Node)
		newEnv := NewEnv()
		for i, param := range fn.Children[1].Children {
			newEnv.vars[param.Value] = Eval(args[1].Children[i], env)
		}
		return Eval(fn.Children[2], newEnv)
	case "car":
		list := Eval(args[0], env).([]interface{})
		return list[0]
	case "cdr":
		list := Eval(args[0], env).([]interface{})
		return list[1:]
	case "cons":
		return append([]interface{}{Eval(args[0], env)}, Eval(args[1], env).([]interface{})...)
	default:
		panic(fmt.Sprintf("Unknown operator: %s", operator))
	}
}
