package int64s

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *int64, err error) {
	if src == "" {
		return
	}
	var pValue int64
	pValue, err = strconv.ParseInt(src, 10, 64)
	if err == nil {
		value = optional(pValue)
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value int64, err error) {
	var pValue *int64
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src int64) string {
	return strconv.FormatInt(src, 10)
}

func OptionalToA(src *int64, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
