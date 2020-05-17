package int32s

import (
	"github.com/litesoft-go/utils/ints"
	"github.com/litesoft-go/utils/strs"
)

func Tertiary(pick1 bool, src1, src2 int32) int32 {
	if pick1 {
		return src1
	}
	return src2
}

func FromA(src string) (value *int32, err error) {
	var atoi *int
	atoi, err = ints.FromA(src)
	if err == nil {
		size := FromInt(*atoi)
		value = &size
	}
	return
}

func ToA(src int32) string {
	return ints.ToA(int(src))
}

//noinspection GoUnusedExportedFunction
func Min(src1, src2 int32) int32 {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 int32) int32 {
	return Tertiary(src1 > src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func OptionalFromA(src string) (value *int32, err error) {
	if src != "" {
		return FromA(src)
	}
	return
}

func OptionalToA(src *int32, defaultOnNil string) string {
	if src == nil {
		return defaultOnNil
	}
	return ToA(*src)
}

//noinspection GoUnusedExportedFunction
func OptionalFromInt(src int) *int32 {
	i := FromInt(src)
	return &i
}

//noinspection GoUnusedExportedFunction
func FromOptional(src *int32, defaultOnNil int32) int32 {
	if src == nil {
		return defaultOnNil
	}
	return *src
}

func FromInt(src int) int32 {
	return int32(src)
}

//noinspection GoUnusedExportedFunction
func OptionalEqual(src1, src2 *int32) bool {
	if (src1 == nil) && (src2 == nil) {
		return true
	}
	if (src1 == nil) || (src2 == nil) {
		return false
	}
	return *src1 == *src2
}

//noinspection GoUnusedExportedFunction
func Update(curV, newV int32, currentUpdated bool) (value int32, updated bool) {
	return newV, currentUpdated || different(curV, newV)
}

//noinspection GoUnusedExportedFunction
func Change(what string, oldValue, newValue int32) string {
	return ints.Change(what, int(oldValue), int(newValue))
}

//noinspection GoUnusedExportedFunction
func OptionalChange(what string, oldValue, newValue *int32) string {
	return strs.RawChange(what, optionalInt32(oldValue), optionalInt32(newValue))
}

func optionalInt32(value *int32) string {
	return OptionalToA(value, "nil")
}

func different(curV, newV int32) bool {
	return curV != newV
}
