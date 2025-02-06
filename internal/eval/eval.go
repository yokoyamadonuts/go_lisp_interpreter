package eval

func eval(expr interface{}, env map[string]interface{}) interface{} {
	switch v := expr.(type) {
	case int, float64:
		return v
	case string:
		return env[v]
	case []interface{}:
		switch v[0] {
		case "+":
			return eval(v[1], env).(int) + eval(v[2], env).(int)
		case "-":
			return eval(v[1], env).(int) - eval(v[2], env).(int)
		case "*":
			return eval(v[1], env).(int) * eval(v[2], env).(int)
		case "/":
			return eval(v[1], env).(int) / eval(v[2], env).(int)
		case "<":
			return eval(v[1], env).(int) < eval(v[2], env).(int)
		case ">":
			return eval(v[1], env).(int) > eval(v[2], env).(int)
		case "=":
			return eval(v[1], env).(int) == eval(v[2], env).(int)
		case "if":
			cond := eval(v[1], env).(bool)
			if cond {
				return eval(v[2], env)
			} else {
				return eval(v[3], env)
			}
		case "lambda":
			return v // 関数オブジェクトとしてそのまま返す
		case "define":
			env[v[1].(string)] = eval(v[2], env)
			return nil
		case "apply":
			fn := eval(v[1], env).([]interface{})
			args := v[2].([]interface{})
			newEnv := make(map[string]interface{})
			for i, param := range fn[1].([]interface{}) {
				newEnv[param.(string)] = eval(args[i], env)
			}
			return eval(fn[2], newEnv)
		case "car":
			return eval(v[1], env).([]interface{})[0]
		case "cdr":
			return eval(v[1], env).([]interface{})[1:]
		case "cons":
			return append([]interface{}{eval(v[1], env)}, eval(v[2], env).([]interface{})...)
		default:
			return "未知の演算"
		}
	}
	return nil
}
