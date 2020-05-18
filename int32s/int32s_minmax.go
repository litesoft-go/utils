package int32s

//noinspection GoUnusedExportedFunction
func Min(src1, src2 int32) int32 {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 int32) int32 {
	return Tertiary(src1 > src2, src1, src2)
}
