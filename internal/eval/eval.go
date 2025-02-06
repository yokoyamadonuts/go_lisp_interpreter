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
		case "if":
			cond := eval(v[1], env).(bool)
			if cond {
				return eval(v[2], env)
			} else {
				return eval(v[3], env)
			}
		}
	}
	return nil
}
