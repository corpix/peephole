package reflect

import (
	"reflect"
)

func ParseType(p string) (Type, error) {
	var (
		buf string
		kt  Type
		vt  Type
		err error
	)

	for k, c := range p {
		buf += string(c)

		switch buf {
		case "map":
			first := k + 2
			last := first
			for kk, vv := range p[first:] {
				switch vv {
				case ']':
					last = first + kk
				default:
					continue
				}
			}

			kt, err = ParseType(string(p[first]) + p[first+1:last])
			if err != nil {
				return TypeInvalid, err
			}
			vt, err = ParseType(p[last+1:])
			if err != nil {
				return TypeInvalid, err
			}

			return reflect.MapOf(kt, vt), nil
		case "[]":
			vt, err = ParseType(p[2:])
			if err != nil {
				return TypeInvalid, err
			}
			return reflect.SliceOf(vt), nil
		case "*":
			vt, err = ParseType(p[1:])
			if err != nil {
				return TypeInvalid, err
			}
			return reflect.PtrTo(vt), nil
		default:
			continue
		}
	}

	switch p {
	case "bool":
		return TypeBool, nil
	case "int":
		return TypeInt, nil
	case "int8":
		return TypeInt8, nil
	case "int16":
		return TypeInt16, nil
	case "int32":
		return TypeInt32, nil
	case "int64":
		return TypeInt64, nil
	case "uint":
		return TypeUint, nil
	case "uint8":
		return TypeUint8, nil
	case "uint16":
		return TypeUint16, nil
	case "uint32":
		return TypeUint32, nil
	case "uint64":
		return TypeUint64, nil
	case "float32":
		return TypeFloat32, nil
	case "float64":
		return TypeFloat64, nil
	case "complex64":
		return TypeComplex64, nil
	case "complex128":
		return TypeComplex128, nil
	case "uintptr":
		return TypeUintptr, nil
	case "string":
		return TypeString, nil
	default:
		return TypeInvalid, NewErrCanNotParseType(p, "unknown type")
	}
}
