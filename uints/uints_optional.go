package uints

//noinspection GoUnusedExportedFunction
func FromOptional(src *uint, defaultOnNil uint) uint {
	if src == nil {
		return defaultOnNil
	}
	return *src
}

//noinspection GoUnusedExportedFunction
func OptionalEqual(src1, src2 *uint) bool {
	if (src1 == nil) && (src2 == nil) {
		return true
	}
	if (src1 == nil) || (src2 == nil) {
		return false
	}
	return *src1 == *src2
}
