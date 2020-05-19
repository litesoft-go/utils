package uint32s

import (
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *uint32, err error) {
	if src == "" {
		return
	}
	var pValue uint64
	pValue, err = strconv.ParseUint(src, 10, 32)
	if err == nil {
		value = optional(uint32(pValue))
	}

	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value uint32, err error) {
	var pValue *uint32
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src uint32) string {
	return strconv.FormatUint(uint64(src), 10)
}

func OptionalToA(src *uint32, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
