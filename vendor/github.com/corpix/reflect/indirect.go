package reflect

import (
	"reflect"
)

func Indirect(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	return IndirectValue(reflect.ValueOf(v)).Interface()
}

func IndirectValue(reflectValue Value) Value {
	if reflectValue.Kind() == reflect.Ptr {
		return reflectValue.Elem()
	}
	return reflectValue
}

func IndirectType(reflectType Type) Type {
	if reflectType == TypeInvalid {
		return TypeInvalid
	}

	if reflectType.Kind() == reflect.Ptr {
		return reflectType.Elem()
	}
	return reflectType
}
