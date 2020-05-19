package uint16s

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *uint16, err error) {
	if src == "" {
		return
	}
	var pValue uint64
	pValue, err = strconv.ParseUint(src, 10, 16)
	if err == nil {
		value = optional(uint16(pValue))
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value uint16, err error) {
	var pValue *uint16
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src uint16) string {
	return strconv.FormatUint(uint64(src), 10)
}

func OptionalToA(src *uint16, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
