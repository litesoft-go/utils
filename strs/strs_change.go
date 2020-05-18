package strs

import (
	"fmt"
	// standard libs only above!
)

//noinspection GoUnusedExportedFunction
func Change(what, oldValue, newValue string) string {
	return RawChange(what, quote(oldValue), quote(newValue))
}

func RawChange(what, oldValue, newValue string) string {
	return fmt.Sprintf("%s: %s -> %s", what, oldValue, newValue)
}

func quote(value string) string {
	return "\"" + value + "\""
}
