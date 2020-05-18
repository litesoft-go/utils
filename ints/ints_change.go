package ints

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue int) string {
	return strs.RawChange(what, ToA(oldValue), ToA(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *int) string {
	return strs.RawChange(what, optionalToString(oldValue), optionalToString(newValue))
}

func optionalToString(value *int) string {
	return OptionalToA(value, "nil")
}
