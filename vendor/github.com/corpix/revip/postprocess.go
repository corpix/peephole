package revip

import (
	"reflect"

	"github.com/fatih/structs"
)

func Postprocess(c Config, op ...Option) error {
	return postprocess(c, nil, op)
}

func postprocessApply(c Config, path []string, op []Option) error {
	var err error
	for _, f := range op {
		err = f(c, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func postprocess(c Config, path []string, op []Option) error {
	err := postprocessApply(c, path, op)
	if err != nil {
		return err
	}

	//n

	t := reflect.TypeOf(c)

	if indirectType(t).Kind() != reflect.Struct {
		return nil
	}

	//

	for _, v := range structs.Fields(c) {
		if !v.IsExported() {
			continue
		}

		err := postprocess(
			v.Value(),
			append(path, v.Name()),
			op,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

//

func WithDefaults() Option {
	return func(c Config, m ...OptionMeta) error {
		v, ok := c.(Defaultable)
		if ok {
			v.Default()
		}
		return nil
	}
}

func WithValidation() Option {
	return func (c Config, m ...OptionMeta) error {
		v, ok := c.(Validatable)
		if ok {
			err := v.Validate()
			if err != nil {
				var path []string
				if len(m) > 0 {
					path = m[0].([]string)
				}
				return &ErrPostprocess{
					Type: reflect.TypeOf(c).String(),
					Path: path,
					Err:  err,
				}
			}
		}
		return nil
	}
}
