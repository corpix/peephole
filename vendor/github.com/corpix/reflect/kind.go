package reflect

import (
	"reflect"
)

type Kind = reflect.Kind

const (
	Invalid       = reflect.Invalid
	Bool          = reflect.Bool
	Int           = reflect.Int
	Int8          = reflect.Int8
	Int16         = reflect.Int16
	Int32         = reflect.Int32
	Int64         = reflect.Int64
	Uint          = reflect.Uint
	Uint8         = reflect.Uint8
	Uint16        = reflect.Uint16
	Uint32        = reflect.Uint32
	Uint64        = reflect.Uint64
	Uintptr       = reflect.Uintptr
	Float32       = reflect.Float32
	Float64       = reflect.Float64
	Complex64     = reflect.Complex64
	Complex128    = reflect.Complex128
	Array         = reflect.Array
	Chan          = reflect.Chan
	Func          = reflect.Func
	Interface     = reflect.Interface
	Map           = reflect.Map
	Ptr           = reflect.Ptr
	Slice         = reflect.Slice
	String        = reflect.String
	Struct        = reflect.Struct
	UnsafePointer = reflect.UnsafePointer
)
