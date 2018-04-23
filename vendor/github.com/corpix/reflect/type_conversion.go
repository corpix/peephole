package reflect

import (
	"reflect"
	"strconv"
)

func ConvertToType(v interface{}, t Type) (interface{}, error) {
	var (
		tt  = reflect.TypeOf(v)
		vv  = reflect.ValueOf(v)
		r   interface{}
		k   int
		err error
	)

	if t.Kind() == Interface {
		return v, nil
	}

	switch tt.Kind() {
	case Slice:
		switch t.Kind() {
		case Slice:
			var (
				res = reflect.MakeSlice(
					t,
					vv.Len(),
					vv.Cap(),
				)
			)

			for k = 0; k < vv.Len(); k++ {
				r, err = ConvertToType(
					vv.Index(k).Interface(),
					t.Elem(),
				)
				if err != nil {
					return nil, err
				}
				res.Index(k).Set(reflect.ValueOf(r))
			}

			return res.Interface(), nil
		}
	case Map:
		switch t.Kind() {
		case Map:
			var (
				res = reflect.MakeMapWithSize(
					t,
					vv.Len(),
				)
				ck interface{}
			)
			for _, k := range vv.MapKeys() {
				ck, err = ConvertToType(
					k.Interface(),
					t.Key(),
				)
				if err != nil {
					return nil, err
				}

				r, err = ConvertToType(
					vv.MapIndex(k).Interface(),
					t.Elem(),
				)
				if err != nil {
					return nil, err
				}

				res.SetMapIndex(
					reflect.ValueOf(ck),
					reflect.ValueOf(r),
				)
			}

			return res.Interface(), nil
		}
	}

	switch t {
	case TypeBool:
		return ConvertToBool(v)
	case TypeInt:
		return ConvertToInt(v)
	case TypeInt8:
		return ConvertToInt8(v)
	case TypeInt16:
		return ConvertToInt16(v)
	case TypeInt32:
		return ConvertToInt32(v)
	case TypeInt64:
		return ConvertToInt64(v)
	case TypeUint:
		return ConvertToUint(v)
	case TypeUint8:
		return ConvertToUint8(v)
	case TypeUint16:
		return ConvertToUint16(v)
	case TypeUint32:
		return ConvertToUint32(v)
	case TypeUint64:
		return ConvertToUint64(v)
	case TypeFloat32:
		return ConvertToFloat32(v)
	case TypeFloat64:
		return ConvertToFloat64(v)
	case TypeComplex64:
		return ConvertToComplex64(v)
	case TypeComplex128:
		return ConvertToComplex128(v)
	case TypeString:
		return ConvertToString(v)
	default:
		return nil, NewErrCanNotConvertType(
			v,
			tt,
			t,
		)
	}
}

func ConvertToBool(vv interface{}) (bool, error) {
	var (
		v = Indirect(vv)
	)

	switch value := v.(type) {
	case bool:
		return value, nil
	case nil:
		return false, nil
	case int:
		if value != 0 {
			return true, nil
		}
		return false, nil
	case string:
		if value == "" {
			return false, nil
		}
		return strconv.ParseBool(value)
	default:
		return false, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeBool,
		)
	}
}

func ConvertToInt(vv interface{}) (int, error) {
	var (
		v   = Indirect(vv)
		res int64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return value, nil
	case int8:
		return int(value), nil
	case int16:
		return int(value), nil
	case int32:
		return int(value), nil
	case int64:
		return int(value), nil
	case uint:
		return int(value), nil
	case uint8:
		return int(value), nil
	case uint16:
		return int(value), nil
	case uint32:
		return int(value), nil
	case uint64:
		return int(value), nil
	case float32:
		return int(value), nil
	case float64:
		return int(value), nil
	case complex64:
		return int(real(value)), nil
	case complex128:
		return int(real(value)), nil
	case uintptr:
		return int(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseInt(value, 0, 0)
		if err == nil {
			return int(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeInt,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeInt,
		)
	}
}

func ConvertToInt8(vv interface{}) (int8, error) {
	var (
		v   = Indirect(vv)
		res int64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return int8(value), nil
	case int8:
		return value, nil
	case int16:
		return int8(value), nil
	case int32:
		return int8(value), nil
	case int64:
		return int8(value), nil
	case uint:
		return int8(value), nil
	case uint8:
		return int8(value), nil
	case uint16:
		return int8(value), nil
	case uint32:
		return int8(value), nil
	case uint64:
		return int8(value), nil
	case float32:
		return int8(value), nil
	case float64:
		return int8(value), nil
	case complex64:
		return int8(real(value)), nil
	case complex128:
		return int8(real(value)), nil
	case uintptr:
		return int8(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseInt(value, 0, 8)
		if err == nil {
			return int8(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeInt8,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeInt8,
		)
	}
}

func ConvertToInt16(vv interface{}) (int16, error) {
	var (
		v   = Indirect(vv)
		res int64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return int16(value), nil
	case int8:
		return int16(value), nil
	case int16:
		return value, nil
	case int32:
		return int16(value), nil
	case int64:
		return int16(value), nil
	case uint:
		return int16(value), nil
	case uint8:
		return int16(value), nil
	case uint16:
		return int16(value), nil
	case uint32:
		return int16(value), nil
	case uint64:
		return int16(value), nil
	case float32:
		return int16(value), nil
	case float64:
		return int16(value), nil
	case complex64:
		return int16(real(value)), nil
	case complex128:
		return int16(real(value)), nil
	case uintptr:
		return int16(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseInt(value, 0, 16)
		if err == nil {
			return int16(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeInt16,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeInt16,
		)
	}
}

func ConvertToInt32(vv interface{}) (int32, error) {
	var (
		v   = Indirect(vv)
		res int64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return int32(value), nil
	case int8:
		return int32(value), nil
	case int16:
		return int32(value), nil
	case int32:
		return value, nil
	case int64:
		return int32(value), nil
	case uint:
		return int32(value), nil
	case uint8:
		return int32(value), nil
	case uint16:
		return int32(value), nil
	case uint32:
		return int32(value), nil
	case uint64:
		return int32(value), nil
	case float32:
		return int32(value), nil
	case float64:
		return int32(value), nil
	case complex64:
		return int32(real(value)), nil
	case complex128:
		return int32(real(value)), nil
	case uintptr:
		return int32(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseInt(value, 0, 32)
		if err == nil {
			return int32(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeInt32,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeInt32,
		)
	}
}

func ConvertToInt64(vv interface{}) (int64, error) {
	var (
		v   = Indirect(vv)
		res int64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return int64(value), nil
	case int8:
		return int64(value), nil
	case int16:
		return int64(value), nil
	case int32:
		return int64(value), nil
	case int64:
		return value, nil
	case uint:
		return int64(value), nil
	case uint8:
		return int64(value), nil
	case uint16:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case uint64:
		return int64(value), nil
	case float32:
		return int64(value), nil
	case float64:
		return int64(value), nil
	case complex64:
		return int64(real(value)), nil
	case complex128:
		return int64(real(value)), nil
	case uintptr:
		return int64(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseInt(value, 0, 64)
		if err == nil {
			return int64(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeInt64,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeInt64,
		)
	}
}

func ConvertToUint(vv interface{}) (uint, error) {
	var (
		v   = Indirect(vv)
		res uint64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return uint(value), nil
	case int8:
		return uint(value), nil
	case int16:
		return uint(value), nil
	case int32:
		return uint(value), nil
	case int64:
		return uint(value), nil
	case uint:
		return value, nil
	case uint8:
		return uint(value), nil
	case uint16:
		return uint(value), nil
	case uint32:
		return uint(value), nil
	case uint64:
		return uint(value), nil
	case float32:
		return uint(value), nil
	case float64:
		return uint(value), nil
	case complex64:
		return uint(real(value)), nil
	case complex128:
		return uint(real(value)), nil
	case uintptr:
		return uint(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseUint(value, 0, 0)
		if err == nil {
			return uint(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeUint,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeUint,
		)
	}
}

func ConvertToUint8(vv interface{}) (uint8, error) {
	var (
		v   = Indirect(vv)
		res uint64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return uint8(value), nil
	case int8:
		return uint8(value), nil
	case int16:
		return uint8(value), nil
	case int32:
		return uint8(value), nil
	case int64:
		return uint8(value), nil
	case uint:
		return uint8(value), nil
	case uint8:
		return value, nil
	case uint16:
		return uint8(value), nil
	case uint32:
		return uint8(value), nil
	case uint64:
		return uint8(value), nil
	case float32:
		return uint8(value), nil
	case float64:
		return uint8(value), nil
	case complex64:
		return uint8(real(value)), nil
	case complex128:
		return uint8(real(value)), nil
	case uintptr:
		return uint8(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseUint(value, 0, 8)
		if err == nil {
			return uint8(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeUint8,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeUint8,
		)
	}
}

func ConvertToUint16(vv interface{}) (uint16, error) {
	var (
		v   = Indirect(vv)
		res uint64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return uint16(value), nil
	case int8:
		return uint16(value), nil
	case int16:
		return uint16(value), nil
	case int32:
		return uint16(value), nil
	case int64:
		return uint16(value), nil
	case uint:
		return uint16(value), nil
	case uint8:
		return uint16(value), nil
	case uint16:
		return value, nil
	case uint32:
		return uint16(value), nil
	case uint64:
		return uint16(value), nil
	case float32:
		return uint16(value), nil
	case float64:
		return uint16(value), nil
	case complex64:
		return uint16(real(value)), nil
	case complex128:
		return uint16(real(value)), nil
	case uintptr:
		return uint16(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseUint(value, 0, 16)
		if err == nil {
			return uint16(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeUint16,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeUint16,
		)
	}

}

func ConvertToUint32(vv interface{}) (uint32, error) {
	var (
		v   = Indirect(vv)
		res uint64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return uint32(value), nil
	case int8:
		return uint32(value), nil
	case int16:
		return uint32(value), nil
	case int32:
		return uint32(value), nil
	case int64:
		return uint32(value), nil
	case uint:
		return uint32(value), nil
	case uint8:
		return uint32(value), nil
	case uint16:
		return uint32(value), nil
	case uint32:
		return value, nil
	case uint64:
		return uint32(value), nil
	case float32:
		return uint32(value), nil
	case float64:
		return uint32(value), nil
	case complex64:
		return uint32(real(value)), nil
	case complex128:
		return uint32(real(value)), nil
	case uintptr:
		return uint32(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseUint(value, 0, 32)
		if err == nil {
			return uint32(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeUint32,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeUint32,
		)
	}
}

func ConvertToUint64(vv interface{}) (uint64, error) {
	var (
		v   = Indirect(vv)
		res uint64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return uint64(value), nil
	case int8:
		return uint64(value), nil
	case int16:
		return uint64(value), nil
	case int32:
		return uint64(value), nil
	case int64:
		return uint64(value), nil
	case uint:
		return uint64(value), nil
	case uint8:
		return uint64(value), nil
	case uint16:
		return uint64(value), nil
	case uint32:
		return uint64(value), nil
	case uint64:
		return value, nil
	case float32:
		return uint64(value), nil
	case float64:
		return uint64(value), nil
	case complex64:
		return uint64(real(value)), nil
	case complex128:
		return uint64(real(value)), nil
	case uintptr:
		return uint64(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseUint(value, 0, 64)
		if err == nil {
			return uint64(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeUint64,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeUint64,
		)
	}
}

func ConvertToFloat32(vv interface{}) (float32, error) {
	var (
		v   = Indirect(vv)
		res float64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return float32(value), nil
	case int8:
		return float32(value), nil
	case int16:
		return float32(value), nil
	case int32:
		return float32(value), nil
	case int64:
		return float32(value), nil
	case uint:
		return float32(value), nil
	case uint8:
		return float32(value), nil
	case uint16:
		return float32(value), nil
	case uint32:
		return float32(value), nil
	case uint64:
		return float32(value), nil
	case float32:
		return value, nil
	case float64:
		return float32(value), nil
	case complex64:
		return float32(real(value)), nil
	case complex128:
		return float32(real(value)), nil
	case uintptr:
		return float32(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseFloat(value, 32)
		if err == nil {
			return float32(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeFloat32,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeFloat32,
		)
	}
}

func ConvertToFloat64(vv interface{}) (float64, error) {
	var (
		v   = Indirect(vv)
		res float64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return float64(value), nil
	case int8:
		return float64(value), nil
	case int16:
		return float64(value), nil
	case int32:
		return float64(value), nil
	case int64:
		return float64(value), nil
	case uint:
		return float64(value), nil
	case uint8:
		return float64(value), nil
	case uint16:
		return float64(value), nil
	case uint32:
		return float64(value), nil
	case uint64:
		return float64(value), nil
	case float32:
		return float64(value), nil
	case float64:
		return value, nil
	case complex64:
		return float64(real(value)), nil
	case complex128:
		return float64(real(value)), nil
	case uintptr:
		return float64(value), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseFloat(value, 32)
		if err == nil {
			return float64(res), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeFloat64,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeFloat64,
		)
	}
}

func ConvertToComplex64(vv interface{}) (complex64, error) {
	var (
		v   = Indirect(vv)
		res float64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return complex64(complex(float32(value), 0)), nil
	case int8:
		return complex64(complex(float32(value), 0)), nil
	case int16:
		return complex64(complex(float32(value), 0)), nil
	case int32:
		return complex64(complex(float32(value), 0)), nil
	case int64:
		return complex64(complex(float32(value), 0)), nil
	case uint:
		return complex64(complex(float32(value), 0)), nil
	case uint8:
		return complex64(complex(float32(value), 0)), nil
	case uint16:
		return complex64(complex(float32(value), 0)), nil
	case uint32:
		return complex64(complex(float32(value), 0)), nil
	case uint64:
		return complex64(complex(float32(value), 0)), nil
	case float32:
		return complex64(complex(float32(value), 0)), nil
	case float64:
		return complex64(complex(float32(value), 0)), nil
	case complex64:
		return value, nil
	case complex128:
		return complex64(value), nil
	case uintptr:
		return complex64(complex(float32(value), 0)), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseFloat(value, 32)
		if err == nil {
			return complex64(complex(float32(res), 0)), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeComplex64,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeComplex64,
		)
	}
}

func ConvertToComplex128(vv interface{}) (complex128, error) {
	var (
		v   = Indirect(vv)
		res float64
		err error
	)

	switch value := v.(type) {
	case bool:
		if value {
			return 1, nil
		}
		return 0, nil
	case int:
		return complex128(complex(float32(value), 0)), nil
	case int8:
		return complex128(complex(float32(value), 0)), nil
	case int16:
		return complex128(complex(float32(value), 0)), nil
	case int32:
		return complex128(complex(float32(value), 0)), nil
	case int64:
		return complex128(complex(float32(value), 0)), nil
	case uint:
		return complex128(complex(float32(value), 0)), nil
	case uint8:
		return complex128(complex(float32(value), 0)), nil
	case uint16:
		return complex128(complex(float32(value), 0)), nil
	case uint32:
		return complex128(complex(float32(value), 0)), nil
	case uint64:
		return complex128(complex(float32(value), 0)), nil
	case float32:
		return complex128(complex(float32(value), 0)), nil
	case float64:
		return complex128(complex(float32(value), 0)), nil
	case complex64:
		return complex128(value), nil
	case complex128:
		return value, nil
	case uintptr:
		return complex128(complex(float32(value), 0)), nil
	case string:
		if value == "" {
			return 0, nil
		}
		res, err = strconv.ParseFloat(value, 32)
		if err == nil {
			return complex128(complex(float32(res), 0)), nil
		}
		return 0, NewErrCanNotConvertType(
			v,
			reflect.TypeOf(vv),
			TypeComplex128,
			err.Error(),
		)
	case nil:
		return 0, nil
	default:
		return 0, NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeComplex128,
		)
	}
}

func ConvertToString(vv interface{}) (string, error) {
	var (
		v = Indirect(vv)
	)

	switch value := v.(type) {
	case bool:
		return strconv.FormatBool(value), nil
	case int:
		return strconv.FormatInt(int64(value), 10), nil
	case int8:
		return strconv.FormatInt(int64(value), 10), nil
	case int16:
		return strconv.FormatInt(int64(value), 10), nil
	case int32:
		return strconv.FormatInt(int64(value), 10), nil
	case int64:
		return strconv.FormatInt(value, 0), nil
	case uint:
		return strconv.FormatInt(int64(value), 10), nil
	case uint8:
		return strconv.FormatInt(int64(value), 10), nil
	case uint16:
		return strconv.FormatInt(int64(value), 10), nil
	case uint32:
		return strconv.FormatInt(int64(value), 10), nil
	case uint64:
		return strconv.FormatInt(int64(value), 10), nil
	case float32:
		return strconv.FormatFloat(float64(value), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64), nil
	case complex64:
		return strconv.FormatFloat(float64(real(value)), 'f', -1, 32), nil
	case complex128:
		return strconv.FormatFloat(real(value), 'f', -1, 64), nil
	case uintptr:
		return strconv.FormatInt(int64(value), 10), nil
	case string:
		return value, nil
	case nil:
		return "", nil
	default:
		return "", NewErrCanNotConvertType(
			vv,
			reflect.TypeOf(vv),
			TypeString,
		)
	}
}
