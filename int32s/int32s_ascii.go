package int32s

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
)

func OptionalFromA(src string) (value *int32, err error) {
	if src == "" {
		return
	}
	var pValue *int
	pValue, err = ints.OptionalFromA(src)
	if (err == nil) && (pValue != nil) {
		value = optional(int32(*pValue))
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value int32, err error) {
	var pValue *int32
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src int32) string {
	return ints.ToA(int(src))
}

func OptionalToA(src *int32, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
