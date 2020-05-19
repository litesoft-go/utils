package strs

import (
	"strings"
	// standard libs only above!
)

//noinspection GoUnusedExportedFunction
func ErrorOrDefault(err error, defaultMsg string) string {
	if err != nil {
		return err.Error()
	}
	return defaultMsg
}

//noinspection GoUnusedExportedFunction
func FromOptional(src *string, defaultValue string) string {
	if src != nil {
		return *src
	}
	return defaultValue
}

//noinspection GoUnusedExportedFunction
func MustStartWithIfNotEmpty(src, prefix string) string {
	if strings.HasPrefix(src, prefix) {
		return src
	}
	if src != "" {
		return prefix + src
	}
	return src
}

//noinspection GoUnusedExportedFunction
func MustStartWith(src, prefix string) string {
	if strings.HasPrefix(src, prefix) {
		return src
	}
	return prefix + src
}

//noinspection GoUnusedExportedFunction
func EqualNonEmpty(src1, src2 string) bool {
	return (src1 != "") && (src1 == src2)
}

func AppendNonEmpty(slice []string, src string) []string {
	if src != "" {
		slice = append(slice, src)
	}
	return slice
}
