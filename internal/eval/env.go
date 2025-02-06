package eval

type Env struct {
	vars map[string]interface{}
}

func NewEnv() *Env {
	return &Env{vars: make(map[string]interface{})}
}

func (e *Env) Set(key string, val interface{}) {
	e.vars[key] = val
}

func (e *Env) Get(key string) interface{} {
	return e.vars[key]
}
