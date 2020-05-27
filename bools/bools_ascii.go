package bools

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *bool, err error) {
	if src == "" {
		return
	}
	var pValue bool
	pValue, err = strconv.ParseBool(src)
	if err == nil {
		value = optional(pValue)
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value bool, err error) {
	var pValue *bool
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src bool) string {
	return strconv.FormatBool(src)
}

func OptionalToA(src *bool, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
