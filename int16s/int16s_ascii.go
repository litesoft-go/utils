package int16s

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
)

func OptionalFromA(src string) (value *int16, err error) {
	if src == "" {
		return
	}
	var pValue *int
	pValue, err = ints.OptionalFromA(src)
	if (err == nil) && (pValue != nil) {
		value = optional(int16(*pValue))
	}
	return
}

//noinspection GoUnusedExportedFunction
func FromA(src string) (value int16, err error) {
	var pValue *int16
	pValue, err = OptionalFromA(src)
	if err == nil {
		value = *pValue
	}
	return
}

func ToA(src int16) string {
	return ints.ToA(int(src))
}

func OptionalToA(src *int16, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}
