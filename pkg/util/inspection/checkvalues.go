package inspection

import (
	"reflect"
)

func IsZeros(v interface{}) bool {
	if v == nil {
		return true
	}
	value := reflect.ValueOf(v)
	return value.IsZero()
}
