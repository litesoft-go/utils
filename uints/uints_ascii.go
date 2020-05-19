package uints

import (
	"math/bits"
	"strconv"
	// standard libs only above!
)

func OptionalFromA(src string) (value *uint, err error) {
	if src == "" {
		return
	}
	var pValue uint64
	pValue, err = strconv.ParseUint(src, 10, bits.UintSize)
	if err == nil {
		value = optional(uint(pValue))
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value uint, err error) {
	var pValue *uint
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src uint) string {
	return strconv.FormatUint(uint64(src), 10)
}

func OptionalToA(src *uint, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
