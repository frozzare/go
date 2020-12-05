package pipe

import (
	"fmt"
	"reflect"

	"github.com/frozzare/go/reflect2"
)

func callFunc(f interface{}, a ...interface{}) ([]reflect.Value, error) {
	fv := reflect.ValueOf(f)
	if reflect2.IsZero(fv) || fv.Kind() != reflect.Func {
		return nil, fmt.Errorf("Function is not callable: %v", f)
	}

	ft := reflect.TypeOf(f)
	ai := ft.NumIn()
	av := make([]reflect.Value, 0, ai)
	for i := 0; i < ai; i++ {
		av = append(av, reflect.ValueOf(a[i]))
	}

	res := fv.Call(av)

	err, ok := res[len(res)-1].Interface().(error)
	if ok && err != nil {
		return nil, err
	}

	return res, nil
}

// Pipeline is the func type for the pipeline result.
type Pipeline func(a ...interface{}) (interface{}, error)

// Pipe performs left-to-right function composition.
// It accepts zero or more funcs fs and creates a pipeline.
func Pipe(fa ...interface{}) Pipeline {
	return func(a ...interface{}) (out interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("pipeline panicked: %v", r)
			}
		}()

		for _, f := range fa {
			res, err := callFunc(f, a...)
			if err != nil {
				return nil, err
			}

			for i, r := range res {
				if reflect2.IsZero(r) {
					continue
				}

				if i >= len(a) {
					a = append(a, r.Interface())
				} else {
					a[i] = r.Interface()
				}
			}
		}

		return a[0], nil
	}
}

// Tap runs the given function with the supplied values, then returns the first supplied value.
// If first supplied value is missing a nil value will be returned instead. If a error occurs it
// will be returned instead of the the first supplied value. This is for being compatible with
// the Pipe function.
func Tap(f interface{}, a ...interface{}) interface{} {
	tap := func(v ...interface{}) interface{} {
		_, err := callFunc(f, v...)
		if err != nil {
			return err
		}

		if len(v) == 0 {
			return nil
		}

		return v[0]
	}

	if len(a) > 0 {
		return tap(a...)
	}

	return tap
}

// Error will check if the given interface is a error and return it.
// If not a error it will return nil.
func Error(v interface{}) error {
	err, ok := v.(error)
	if ok && err != nil {
		return err
	}
	return nil
}