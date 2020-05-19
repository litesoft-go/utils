package int16s

//noinspection GoUnusedExportedFunction
func FromOptional(src *int16, defaultOnNil int16) int16 {
	if src == nil {
		return defaultOnNil
	}
	return *src
}

//noinspection GoUnusedExportedFunction
func OptionalEqual(src1, src2 *int16) bool {
	if (src1 == nil) && (src2 == nil) {
		return true
	}
	if (src1 == nil) || (src2 == nil) {
		return false
	}
	return *src1 == *src2
}
