package int32s

//noinspection GoUnusedExportedFunction
func FromOptional(src *int32, defaultOnNil int32) int32 {
	if src == nil {
		return defaultOnNil
	}
	return *src
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
