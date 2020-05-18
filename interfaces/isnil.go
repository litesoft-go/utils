package interfaces

import "reflect"

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	kind := reflect.TypeOf(i).Kind()
	if kind != reflect.Ptr { // Manually check the common case
		switch kind { // switch on the other options
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice: // continue
		default:
			return false
		}
	}
	return reflect.ValueOf(i).IsNil()
}
