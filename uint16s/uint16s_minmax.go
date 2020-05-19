package uint16s

//noinspection GoUnusedExportedFunction
func Min(src1, src2 uint16) uint16 {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 uint16) uint16 {
	return Tertiary(src1 > src2, src1, src2)
}
