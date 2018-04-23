package reflect

import (
	"reflect"
)

var (
	TypeOf = reflect.TypeOf
)

type Type = reflect.Type

var (
	Types = []Type{
		TypeBool,
		TypeInt,
		TypeInt8,
		TypeInt16,
		TypeInt32,
		TypeInt64,
		TypeUint,
		TypeUint8,
		TypeUint16,
		TypeUint32,
		TypeUint64,
		TypeFloat32,
		TypeFloat64,
		TypeComplex64,
		TypeComplex128,
		TypeUintptr,
		TypeString,
	}

	TypeInvalid    = Type(nil)
	TypeBool       = reflect.TypeOf(false)
	TypeInt        = reflect.TypeOf(int(0))
	TypeInt8       = reflect.TypeOf(int8(0))
	TypeInt16      = reflect.TypeOf(int16(0))
	TypeInt32      = reflect.TypeOf(int32(0))
	TypeInt64      = reflect.TypeOf(int64(0))
	TypeUint       = reflect.TypeOf(uint(0))
	TypeUint8      = reflect.TypeOf(uint8(0))
	TypeUint16     = reflect.TypeOf(uint16(0))
	TypeUint32     = reflect.TypeOf(uint32(0))
	TypeUint64     = reflect.TypeOf(uint64(0))
	TypeFloat32    = reflect.TypeOf(float32(0))
	TypeFloat64    = reflect.TypeOf(float64(0))
	TypeComplex64  = reflect.TypeOf(complex64(0))
	TypeComplex128 = reflect.TypeOf(complex128(0))
	TypeUintptr    = reflect.TypeOf(uintptr(0))
	TypeString     = reflect.TypeOf(string(""))
)
