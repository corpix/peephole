package formats

type Stringer string

func (s *Stringer) String() string {
	return string(*s)
}

func NewStringer(s string) *Stringer {
	sr := Stringer(s)
	return &sr
}
