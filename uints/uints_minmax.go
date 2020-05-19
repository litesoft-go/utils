package uints

//noinspection GoUnusedExportedFunction
func Min(src1, src2 uint) uint {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 uint) uint {
	return Tertiary(src1 > src2, src1, src2)
}
