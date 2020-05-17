package strs

import (
	"fmt"
	"strings"
)

//noinspection GoUnusedExportedFunction
func ErrorOrDefault(err error, defaultMsg string) string {
	if err != nil {
		return err.Error()
	}
	return defaultMsg
}

// DO NOT USE WITH ERRORS!
//
// strs.Tertiary(err != nil, err.Error(), "N/A") leads to nil pointer dereference in case of nil error
// as soon as second parameter evaluated in any case
//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 string) string {
	if pick1 {
		return src1
	}
	return src2
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
func Update(curV, newV string, currentUpdated bool) (value string, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func UpdateFunc(curV, newV string, currentUpdated bool, setter func(string)) (updated bool) {
	if !different(curV, newV) {
		return currentUpdated
	}
	setter(newV)
	return true
}

//noinspection GoUnusedExportedFunction
func EqualNonEmpty(src1, src2 string) bool {
	return (src1 != "") && (src1 == src2)
}

func RawChange(what, oldValue, newValue string) string {
	return fmt.Sprintf("%s: %s -> %s", what, oldValue, newValue)
}

//noinspection GoUnusedExportedFunction
func Change(what, oldValue, newValue string) string {
	return RawChange(what, quote(oldValue), quote(newValue))
}

func AppendNonEmpty(slice []string, src string) []string {
	if src != "" {
		slice = append(slice, src)
	}
	return slice
}

func quote(value string) string {
	return "\"" + value + "\""
}

func different(curV, newV string) bool {
	return curV != newV
}
