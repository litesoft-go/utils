package int16s

//noinspection GoUnusedExportedFunction
func Min(src1, src2 int16) int16 {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 int16) int16 {
	return Tertiary(src1 > src2, src1, src2)
}
