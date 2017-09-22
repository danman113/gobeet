package re

type RE interface {
	Res() interface{}
	Err() error
}

// Example
type returnType struct {
	r string
	e error
}

func (rt *returnType) Res() interface{} {
	return rt.r
}

func (rt *returnType) Err() error {
	return rt.e
}
