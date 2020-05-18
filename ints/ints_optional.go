package ints

//noinspection GoUnusedExportedFunction
func FromOptional(src *int, defaultOnNil int) int {
	if src == nil {
		return defaultOnNil
	}
	return *src
}

//noinspection GoUnusedExportedFunction
func OptionalEqual(src1, src2 *int) bool {
	if (src1 == nil) && (src2 == nil) {
		return true
	}
	if (src1 == nil) || (src2 == nil) {
		return false
	}
	return *src1 == *src2
}
