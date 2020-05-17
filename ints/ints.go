package ints

import (
	"strconv"

	"github.com/litesoft-go/utils/strs"
)

//noinspection GoUnusedExportedFunction
func Tertiary(pick1 bool, src1, src2 int) int {
	if pick1 {
		return src1
	}
	return src2
}

func FromA(src string) (value *int, err error) {
	var atoi int
	atoi, err = strconv.Atoi(src)
	if err == nil {
		value = &atoi
	}
	return
}

func ToA(src int) string {
	return strconv.Itoa(src)
}

//noinspection GoUnusedExportedFunction
func Min(src1, src2 int) int {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 int) int {
	return Tertiary(src1 > src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func IsOneOf(check int, acceptables ...int) bool {
	for _, acceptable := range acceptables {
		if check == acceptable {
			return true
		}
	}
	return false
}

//noinspection GoUnusedExportedFunction
func OptionalInt32toInt(src *int32, defaultOnNil int32) int {
	return int(OptionalInt32toInt32(src, defaultOnNil))
}

func OptionalInt32toInt32(src *int32, defaultOnNil int32) int32 {
	if src == nil {
		return defaultOnNil
	}
	return *src
}

func Change(what string, oldValue, newValue int) string {
	return strs.RawChange(what, ToA(oldValue), ToA(newValue))
}
