package structs

import (
	"reflect"
)

// Iterate ...
func Iterate(s interface{}, action func(k string, v interface{})) {
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		action(typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}
