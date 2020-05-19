package uint8s

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *uint8, err error) {
	if src == "" {
		return
	}
	var pValue uint64
	pValue, err = strconv.ParseUint(src, 10, 8)
	if err == nil {
		value = optional(uint8(pValue))
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value uint8, err error) {
	var pValue *uint8
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src uint8) string {
	return strconv.FormatUint(uint64(src), 10)
}

func OptionalToA(src *uint8, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
