package cast

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// getArg returns a arg value or the default value.
func getArg(def interface{}, i int, args ...interface{}) interface{} {
	if len(args) >= i+1 && args[i] != nil {
		return args[i]
	}

	return def
}

// getArgInt returns a arg value as int.
func getArgInt(def int, i int, args ...interface{}) int {
	return getArg(def, i, args...).(int)
}

// Bool converts argument to bool or return a error.
func Bool(value interface{}) (bool, error) {
	switch value.(type) {
	case bool:
		return value.(bool), nil
	case float32:
		v, _ := value.(float32)

		return v > 0, nil
	case float64:
		v, _ := value.(float64)

		return v > 0, nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return v > 0, nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return v > 0, nil
	case string:
		v, _ := value.(string)

		return strconv.ParseBool(v)
	default:
		return false, errors.New("Unknown type")
	}
}

// MustBool converts argument to bool or panic if an error occurred.
func MustBool(value interface{}) bool {
	v, err := Bool(value)

	if err != nil {
		panic(err)
	}

	return v
}

// Float32 will converts argument to float32 or return a error.
func Float32(value interface{}) (float32, error) {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}

		return 0, nil
	case float32:
		v, _ := value.(float32)

		return float32(v), nil
	case float64:
		v, _ := value.(float64)

		return float32(v), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return float32(v), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return float32(v), nil
	case string:
		v, _ := value.(string)
		f, err := strconv.ParseFloat(v, 32)

		return float32(f), err
	default:
		return float32(0), errors.New("Unknown type")
	}
}

// MustFloat32 converts argument to float32 or panic if an error occurred.
func MustFloat32(value interface{}) float32 {
	v, err := Float32(value)

	if err != nil {
		panic(err)
	}

	return v
}

// Float64 will converts argument to float64 or return a error.
func Float64(value interface{}) (float64, error) {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}

		return 0, nil
	case float32:
		v, _ := value.(float32)

		return float64(v), nil
	case float64:
		v, _ := value.(float64)

		return float64(v), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return float64(v), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return float64(v), nil
	case string:
		v, _ := value.(string)
		f, err := strconv.ParseFloat(v, 64)

		return f, err
	default:
		return float64(0), errors.New("Unknown type")
	}
}

// MustFloat64 converts argument to float64 or panic if an error occurred.
func MustFloat64(value interface{}) float64 {
	v, err := Float64(value)

	if err != nil {
		panic(err)
	}

	return v
}

// Int will converts argument to int or return a error.
func Int(value interface{}) (int, error) {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}

		return 0, nil
	case float32:
		v, _ := value.(float32)

		return int(v), nil
	case float64:
		v, _ := value.(float64)

		return int(v), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return int(v), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return int(v), nil
	case string:
		v, _ := value.(string)
		f, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return 0, err
		}

		return int(f), nil
	default:
		return 0, errors.New("Unknown type")
	}
}

// MustInt converts argument to int or panic if an error occurred.
func MustInt(value interface{}) int {
	v, err := Int(value)

	if err != nil {
		panic(err)
	}

	return v
}

// String will converts argument to string or return a error.
func String(args ...interface{}) (string, error) {
	value := args[0]

	switch value.(type) {
	case bool:
		v, _ := value.(bool)

		return strconv.FormatBool(v), nil
	case []byte:
		v, _ := value.([]byte)

		return string(v), nil
	case float32:
		v, _ := value.(float32)

		return strconv.FormatFloat(float64(v), 'f', getArgInt(12, 1, args), 64), nil
	case float64:
		v, _ := value.(float64)

		return strconv.FormatFloat(v, 'f', getArgInt(12, 1, args), 64), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return strconv.FormatInt(int64(v), getArgInt(10, 1, args)), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return strconv.FormatUint(uint64(v), getArgInt(10, 1, args)), nil
	case string:
		v, _ := value.(string)

		return v, nil
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

// MustString converts argument to string or panic if an error occurred.
func MustString(value interface{}) string {
	v, err := String(value)

	if err != nil {
		panic(err)
	}

	return v
}
