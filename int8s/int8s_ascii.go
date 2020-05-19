package int8s

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
)

func OptionalFromA(src string) (value *int8, err error) {
	if src == "" {
		return
	}
	var pValue *int
	pValue, err = ints.OptionalFromA(src)
	if (err == nil) && (pValue != nil) {
		value = optional(int8(*pValue))
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value int8, err error) {
	var pValue *int8
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src int8) string {
	return ints.ToA(int(src))
}

func OptionalToA(src *int8, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
