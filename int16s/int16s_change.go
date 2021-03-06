package int16s

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue int16) string {
	return ints.Change(what, int(oldValue), int(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *int16) string {
	return strs.RawChange(what, optionalToString(oldValue), optionalToString(newValue))
}

func optionalToString(value *int16) string {
	return OptionalToA(value, "nil")
}
