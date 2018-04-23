package formats

import (
	"encoding/hex"
	"fmt"

	"github.com/corpix/reflect"
)

var (
	stringerType = reflect.TypeOf(NewStringer(""))
)

// HEXFormat is a HEX marshaler.
type HEXFormat struct{}

func (f *HEXFormat) indirectToFstPtr(v interface{}) reflect.Value {
	var (
		rv = reflect.ValueOf(v)
	)
	for {
		if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Ptr {
			rv = rv.Elem()
			continue
		}
		break
	}
	return rv
}

// Marshal serializes data represented by v into slice of bytes.
func (f *HEXFormat) Marshal(v interface{}) ([]byte, error) {
	var (
		rv = f.indirectToFstPtr(v)
		vs fmt.Stringer
	)

	if rv.Type() != stringerType || !rv.IsValid() {
		return nil, reflect.NewErrCanNotAssertType(
			v,
			stringerType,
		)
	}

	vs = rv.Interface().(*Stringer)

	var (
		vv  = []byte(vs.String())
		buf = make([]byte, hex.EncodedLen(len(vv)))
	)

	hex.Encode(buf, vv)
	return buf, nil
}

// Unmarshal deserializes data represented by data byte slice into v.
func (f *HEXFormat) Unmarshal(data []byte, v interface{}) error {
	var (
		buf = make([]byte, hex.DecodedLen(len(data)))
		rv  = f.indirectToFstPtr(v)
		err error
	)

	if rv.Kind() != reflect.Ptr {
		return reflect.NewErrWrongKind(reflect.Ptr, rv.Kind())

	}

	_, err = hex.Decode(buf, data)
	if err != nil {
		return err
	}

	if rv.Type() != stringerType {
		if rv.Elem().Kind() == reflect.Interface {
			rv.Elem().Set(reflect.ValueOf(*NewStringer(string(buf))))
			return nil
		}
		return reflect.NewErrCanNotAssertType(
			v,
			stringerType,
		)
	}

	rv.Set(reflect.ValueOf(NewStringer(string(buf))))

	return nil
}

// Name returns a format name which is used to identify this format
// in the package.
func (f *HEXFormat) Name() string {
	return HEX
}

// NewHEX constructs a new HEX format marshaler.
func NewHEX() *HEXFormat { return &HEXFormat{} }
