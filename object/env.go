package object

type Env struct {
	out *Env

	values map[string]Object
}

func (e *Env) Set(key string, value Object) {
	e.values[key] = value
}

func (e *Env) Get(key string) (Object, bool) {
	value, ok := e.values[key]
	if !ok && e.out != nil {
		value, ok = e.out.Get(key)
	}

	return value, ok
}

func NewEnv(out *Env) *Env {
	return &Env{
		out: out,

		values: make(map[string]Object),
	}
}
