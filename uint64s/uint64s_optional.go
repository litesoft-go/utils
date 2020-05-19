package uint64s

//noinspection GoUnusedExportedFunction
func FromOptional(src *uint64, defaultOnNil uint64) uint64 {
	if src == nil {
		return defaultOnNil
	}
	return *src
}

//noinspection GoUnusedExportedFunction
func OptionalEqual(src1, src2 *uint64) bool {
	if (src1 == nil) && (src2 == nil) {
		return true
	}
	if (src1 == nil) || (src2 == nil) {
		return false
	}
	return *src1 == *src2
}
