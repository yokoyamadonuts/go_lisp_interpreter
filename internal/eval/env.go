package eval

type Env struct {
	vars   map[string]interface{}
	parent *Env
}

func NewEnv(parent *Env) *Env {
	return &Env{vars: make(map[string]interface{}), parent: parent}
}

func (e *Env) Set(key string, val interface{}) {
	e.vars[key] = val
}

func (e *Env) Get(key string) (interface{}, bool) {
	if val, ok := e.vars[key]; ok {
		return val, true
	} else if e.parent != nil {
		return e.parent.Get(key)
	}
	return nil, false
}
