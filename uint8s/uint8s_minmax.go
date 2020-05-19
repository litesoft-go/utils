package uint8s

//noinspection GoUnusedExportedFunction
func Min(src1, src2 uint8) uint8 {
	return Tertiary(src1 < src2, src1, src2)
}

//noinspection GoUnusedExportedFunction
func Max(src1, src2 uint8) uint8 {
	return Tertiary(src1 > src2, src1, src2)
}
