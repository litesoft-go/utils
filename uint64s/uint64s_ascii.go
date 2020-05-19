package uint64s

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *uint64, err error) {
	if src == "" {
		return
	}
	var pValue uint64
	pValue, err = strconv.ParseUint(src, 10, 64)
	if err == nil {
		value = optional(pValue)
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value uint64, err error) {
	var pValue *uint64
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src uint64) string {
	return strconv.FormatUint(src, 10)
}

func OptionalToA(src *uint64, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
