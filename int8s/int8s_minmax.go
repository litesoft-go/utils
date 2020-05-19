package int8s

//noinspection GoUnusedExportedFunction
func Min(src1, src2 int8) int8 {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 int8) int8 {
	return Tertiary(src1 > src2, src1, src2)
}
