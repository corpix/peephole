package encoder

// Encoder is an interface which describes encoder for some
// type of data.
type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}
