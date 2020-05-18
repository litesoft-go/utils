package ints

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *int, err error) {
	if src == "" {
		return
	}
	var pValue int
	pValue, err = strconv.Atoi(src)
	if err == nil {
		value = optional(pValue)
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value int, err error) {
	var pValue *int
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src int) string {
	return strconv.Itoa(src)
}

func OptionalToA(src *int, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
