package parallel

type Runner interface {
	Run() (interface{}, error)
}

type mockRunner struct {
	result interface{}
	err    error
}

func (r mockRunner) Run() (interface{}, error) {
	return r.result, r.err
}
