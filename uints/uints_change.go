package uints

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue uint) string {
	return strs.RawChange(what, ToA(oldValue), ToA(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *uint) string {
	return strs.RawChange(what, optionalToString(oldValue), optionalToString(newValue))
}

func optionalToString(value *uint) string {
	return OptionalToA(value, "nil")
}
