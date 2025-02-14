package eval

import (
	"fmt"
	"go_lisp_interpreter/internal/parser"
	"strconv"
)

func eval(ast parser.Node, env *Env) interface{} {
	if ast.Type == "ATOM" {
		if val, ok := env.Get(ast.Value); ok {
			return val
		}
		// 数値変換を試みる
		if num, err := strconv.Atoi(ast.Value); err == nil {
			return num
		}
		// 真偽値変換を試みる
		if b, err := strconv.ParseBool(ast.Value); err == nil {
			return b
		}
		return ast.Value
	}

	if ast.Type == "LIST" {
		if len(ast.Children) == 0 {
			return nil
		}

		first := ast.Children[0]

		if first.Type == "ATOM" {
			switch first.Value {
			case "+":
				return toInt(eval(ast.Children[1], env)) + toInt(eval(ast.Children[2], env))
			case "-":
				return toInt(eval(ast.Children[1], env)) - toInt(eval(ast.Children[2], env))
			case "*":
				return toInt(eval(ast.Children[1], env)) * toInt(eval(ast.Children[2], env))
			case "/":
				return toInt(eval(ast.Children[1], env)) / toInt(eval(ast.Children[2], env))
			case "define":
				env.Set(ast.Children[1].Value, eval(ast.Children[2], env))
				return nil
			case "if":
				cond := toBool(eval(ast.Children[1], env))
				if cond {
					return eval(ast.Children[2], env)
				} else {
					return eval(ast.Children[3], env)
				}
			case "lambda":
				return ast // Lambda式を関数オブジェクトとして返す
			case "apply":
				fn := eval(ast.Children[1], env).(parser.Node)
				newEnv := NewEnv(env)
				for i, param := range fn.Children[1].Children {
					newEnv.Set(param.Value, eval(ast.Children[2].Children[i], env))
				}
				return eval(fn.Children[2], newEnv)
			case "cons":
				return append([]interface{}{eval(ast.Children[1], env)}, eval(ast.Children[2], env).([]interface{})...)
			case "car":
				list := eval(ast.Children[1], env).([]interface{})
				return list[0]
			case "cdr":
				list := eval(ast.Children[1], env).([]interface{})
				return list[1:]
			default:
				// 先頭要素が演算子でない場合、そのままリストとして評価
				result := []interface{}{}
				for _, child := range ast.Children {
					result = append(result, eval(child, env))
				}
				return result
			}
		}
		// 演算子ではなくデータとしてのリストである場合、すべての要素を評価する
		result := []interface{}{}
		for _, child := range ast.Children {
			result = append(result, eval(child, env))
		}
		return result
	}
	return nil
}

func toInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case string:
		if num, err := strconv.Atoi(v); err == nil {
			return num
		}
	}
	panic(fmt.Sprintf("Cannot convert %v to int", value))
}

func toBool(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		if b, err := strconv.ParseBool(v); err == nil {
			return b
		}
	}
	panic(fmt.Sprintf("Cannot convert %v to bool", value))
}
