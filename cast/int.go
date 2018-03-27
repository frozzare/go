package cast

import (
	"errors"
	"strconv"
)

// Int will converts argument to int64 or return a error.
func Int(value interface{}) (int64, error) {
	switch v := value.(type) {
	case bool:
		if v {
			return int64(1), nil
		}

		return int64(0), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return int64(0), nil
		}
		return int64(f), nil
	case []byte:
		return strconv.ParseInt(string(v), 10, 64)
	case nil:
		return int64(0), nil
	default:
		return int64(0), errors.New("Unknown type")
	}
}

// MustInt converts argument to int or panic if an error occurred.
func MustInt(value interface{}) int64 {
	v, err := Int(value)

	if err != nil {
		panic(err)
	}

	return v
}
