package bools

import (
	"github.com/litesoft-go/utils/strs"
	// standard libs only above!
)

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue bool) string {
	return strs.RawChange(what, ToA(oldValue), ToA(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *bool) string {
	return strs.RawChange(what, optionalToString(oldValue), optionalToString(newValue))
}

func optionalToString(value *bool) string {
	return OptionalToA(value, "nil")
}
