package int64s

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue int64) string {
	return ints.Change(what, int(oldValue), int(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *int64) string {
	return strs.RawChange(what, optionalToString(oldValue), optionalToString(newValue))
}

func optionalToString(value *int64) string {
	return OptionalToA(value, "nil")
}
