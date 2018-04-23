package pool

type Result struct {
	Value interface{}
	Err   error
}

func NewResult(v interface{}, err error) *Result {
	return &Result{
		Value: v,
		Err:   err,
	}
}
