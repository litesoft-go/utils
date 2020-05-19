package uint32s

import (
	// standard libs only above!

	"github.com/litesoft-go/utils/ints"
	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue uint32) string {
	return ints.Change(what, int(oldValue), int(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *uint32) string {
	return strs.RawChange(what, optionalToString(oldValue), optionalToString(newValue))
}

func optionalToString(value *uint32) string {
	return OptionalToA(value, "nil")
}
